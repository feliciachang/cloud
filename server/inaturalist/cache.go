package inaturalist

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx/types"

	"github.com/conservify/gonaturalist"
	"github.com/conservify/sqlxcache"

	"github.com/fieldkit/cloud/server/data"
	"github.com/fieldkit/cloud/server/jobs"
)

type CachedObservation struct {
	ID        int64          `db:"id,omitempty"`
	SiteID    int64          `db:"site_id,omitempty"`
	UpdatedAt time.Time      `db:"updated_at"`
	Timestamp time.Time      `db:"timestamp"`
	Location  *data.Location `db:"location"`
}

type FullCachedObservation struct {
	CachedObservation
	Data types.JSONText `db:"data"`
}

func NewCachedObservation(o *gonaturalist.SimpleObservation) (co *FullCachedObservation, err error) {
	jsonData, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}

	timestamp := o.TimeObservedAtUtc
	if timestamp.IsZero() {
		timestamp, err = o.TryParseObservedOn()
		if err != nil {
			return nil, nil
		}
	}

	co = &FullCachedObservation{
		CachedObservation: CachedObservation{
			ID:        o.Id,
			SiteID:    o.SiteId,
			UpdatedAt: time.Now(),
			Timestamp: timestamp,
			Location:  data.NewLocation([]float64{o.Longitude, o.Latitude}),
		},
		Data: jsonData,
	}

	return
}

func (co *CachedObservation) Valid() bool {
	return !co.Timestamp.IsZero() && !co.Location.IsZero()
}

type INaturalistCache struct {
	Config            *INaturalistConfig
	Database          *sqlxcache.DB
	INaturalistClient *gonaturalist.Client
	Queue             *jobs.PgJobQueue
}

func NewINaturalistCache(config *INaturalistConfig, client *gonaturalist.Client, db *sqlxcache.DB, queue *jobs.PgJobQueue) (in *INaturalistCache, err error) {
	in = &INaturalistCache{
		Config:            config,
		INaturalistClient: client,
		Database:          db,
		Queue:             queue,
	}

	return
}

func (in *INaturalistCache) DeleteOld(ctx context.Context) error {
	log := Logger(ctx).Sugar()

	log.Infow("Deleting old Observations")

	empty := struct{}{}
	_, err := in.Database.NamedExecContext(ctx, `DELETE FROM fieldkit.inaturalist_observations WHERE timestamp < NOW() - interval '7 days'`, empty)
	if err != nil {
		return err
	}

	return nil
}

func (in *INaturalistCache) AddOrUpdateObservation(ctx context.Context, o *gonaturalist.SimpleObservation) error {
	co, err := NewCachedObservation(o)
	if err != nil {
		return err
	}

	if co == nil || !co.Valid() {
		return nil
	}

	_, err = in.Database.NamedExecContext(ctx, `
		INSERT INTO fieldkit.inaturalist_observations (id, site_id, updated_at, timestamp, location, data)
		VALUES (:id, :site_id, :updated_at, :timestamp, ST_SetSRID(ST_GeomFromText(:location), 4326), :data)
		ON CONFLICT (id, site_id)
		DO UPDATE SET updated_at = excluded.updated_at, timestamp = excluded.timestamp, location = excluded.location, data = excluded.data
		`, co)
	if err != nil {
		return err
	}

	if err := in.Queue.Publish(ctx, co.CachedObservation); err != nil {
		return err
	}

	return nil
}

func (in *INaturalistCache) RefreshRecentlyUpdated(ctx context.Context, updatedSince time.Time) error {
	orderBy := "created_at"
	hasGeo := true

	err := in.DeleteOld(ctx)
	if err != nil {
		return err
	}

	options := &gonaturalist.GetObservationsOpt{
		HasGeo:       &hasGeo,
		OrderBy:      &orderBy,
		UpdatedSince: &updatedSince,
	}

	return in.refreshUntilEmptyPage(ctx, options)
}

func (in *INaturalistCache) RefreshObservedOn(ctx context.Context, on time.Time) error {
	hasGeo := true

	options := &gonaturalist.GetObservationsOpt{
		HasGeo: &hasGeo,
		On:     &on,
	}

	return in.refreshUntilEmptyPage(ctx, options)
}

func (in *INaturalistCache) RefreshObservation(ctx context.Context, id int) error {
	observation, err := in.INaturalistClient.GetSimpleObservation(int64(id))
	if err != nil {
		return err
	}
	return in.AddOrUpdateObservation(ctx, observation)
}

func (in *INaturalistCache) getLastUpdated() (time.Time, error) {
	return time.Time{}, nil
}

func (in *INaturalistCache) refreshUntilEmptyPage(ctx context.Context, options *gonaturalist.GetObservationsOpt) error {
	perPage := 100
	page := 1
	numberOfObservations := 0

	log := Logger(ctx).Sugar()

	for {
		options.Page = &page
		options.PerPage = &perPage

		log.Infow(fmt.Sprintf("Querying %d", page), "api_url", in.Config.RootUrl)

		observations, err := in.INaturalistClient.GetObservations(options)
		if err != nil {
			return fmt.Errorf("Refresh error: %v", err)
		}

		if len(observations.Observations) == 0 {
			break
		}

		for _, o := range observations.Observations {
			numberOfObservations += 1
			if err := in.AddOrUpdateObservation(ctx, o); err != nil {
				return err
			}
		}

		page += 1
	}

	log.Infow(fmt.Sprintf("Refreshed"), "api_url", in.Config.RootUrl, "number_of_observations", numberOfObservations)

	return nil
}
