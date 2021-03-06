package api

import (
	"github.com/goadesign/goa"

	"github.com/conservify/sqlxcache"

	"github.com/fieldkit/cloud/server/api/app"
	"github.com/fieldkit/cloud/server/backend"
	"github.com/fieldkit/cloud/server/data"
)

type ExpeditionControllerOptions struct {
	Database *sqlxcache.DB
	Backend  *backend.Backend
}

func ExpeditionType(expedition *data.Expedition) *app.Expedition {
	return &app.Expedition{
		ID:          int(expedition.ID),
		Name:        expedition.Name,
		Slug:        expedition.Slug,
		Description: expedition.Description,
	}
}

func ExpeditionsType(expeditions []*data.Expedition) *app.Expeditions {
	expeditionsCollection := make([]*app.Expedition, len(expeditions))
	for i, expedition := range expeditions {
		expeditionsCollection[i] = ExpeditionType(expedition)
	}

	return &app.Expeditions{
		Expeditions: expeditionsCollection,
	}
}

// ExpeditionController implements the user resource.
type ExpeditionController struct {
	*goa.Controller
	options ExpeditionControllerOptions
}

func NewExpeditionController(service *goa.Service, options ExpeditionControllerOptions) *ExpeditionController {
	return &ExpeditionController{
		Controller: service.NewController("ExpeditionController"),
		options:    options,
	}
}

func (c *ExpeditionController) Add(ctx *app.AddExpeditionContext) error {
	expedition := &data.Expedition{
		ProjectID:   int32(ctx.ProjectID),
		Name:        ctx.Payload.Name,
		Slug:        ctx.Payload.Slug,
		Description: ctx.Payload.Description,
	}

	if err := c.options.Database.NamedGetContext(ctx, expedition, "INSERT INTO fieldkit.expedition (project_id, name, slug, description) VALUES (:project_id, :name, :slug, :description) RETURNING *", expedition); err != nil {
		return err
	}

	return ctx.OK(ExpeditionType(expedition))
}

func (c *ExpeditionController) Update(ctx *app.UpdateExpeditionContext) error {
	expedition := &data.Expedition{
		ID:          int32(ctx.ExpeditionID),
		Name:        ctx.Payload.Name,
		Slug:        ctx.Payload.Slug,
		Description: ctx.Payload.Description,
	}

	if err := c.options.Database.NamedGetContext(ctx, expedition, "UPDATE fieldkit.expedition SET name = name, slug = slug, description = description WHERE id = :id RETURNING *", expedition); err != nil {
		return err
	}

	return ctx.OK(ExpeditionType(expedition))
}

func (c *ExpeditionController) Get(ctx *app.GetExpeditionContext) error {
	expedition := &data.Expedition{}
	if err := c.options.Database.GetContext(ctx, expedition, "SELECT e.* FROM fieldkit.expedition AS e JOIN fieldkit.project AS p ON p.id = e.project_id WHERE p.slug = $1 AND e.slug = $2", ctx.Project, ctx.Expedition); err != nil {
		return err
	}

	return ctx.OK(ExpeditionType(expedition))
}

func (c *ExpeditionController) GetID(ctx *app.GetIDExpeditionContext) error {
	expedition := &data.Expedition{}
	if err := c.options.Database.GetContext(ctx, expedition, "SELECT * FROM fieldkit.expedition WHERE id = $1", ctx.ExpeditionID); err != nil {
		return err
	}

	return ctx.OK(ExpeditionType(expedition))
}

func (c *ExpeditionController) List(ctx *app.ListExpeditionContext) error {
	expeditions := []*data.Expedition{}
	if err := c.options.Database.SelectContext(ctx, &expeditions, "SELECT e.* FROM fieldkit.expedition AS e JOIN fieldkit.project AS p ON p.id = e.project_id WHERE p.slug = $1", ctx.Project); err != nil {
		return err
	}

	return ctx.OK(ExpeditionsType(expeditions))
}

func (c *ExpeditionController) ListID(ctx *app.ListIDExpeditionContext) error {
	expeditions := []*data.Expedition{}
	if err := c.options.Database.SelectContext(ctx, &expeditions, "SELECT * FROM fieldkit.expedition WHERE project_id = $1", ctx.ProjectID); err != nil {
		return err
	}

	return ctx.OK(ExpeditionsType(expeditions))
}
