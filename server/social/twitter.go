package social

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/fieldkit/cloud/server/backend"
	"github.com/fieldkit/cloud/server/data"
	"github.com/fieldkit/cloud/server/social/twitter"
)

const (
	TwitterJSONSchema = `
{
	"id": "https://api.fieldkit.org/schemas/tweet.json",
	"type": "object",
	"properties": {
		"id_str": {
			"type": "string"
		},
		"text": {
			"type": "string"
		}
	}
}
`
)

type TwitterOptions struct {
	StreamOptions twitter.StreamOptions
	Backend       *backend.Backend
	Done          <-chan struct{}
}

func Twitter(options TwitterOptions) {
	jsonSchema := &data.JSONSchema{}
	if err := jsonSchema.UnmarshalJSON([]byte(TwitterJSONSchema)); err != nil {
		panic(err)
	}

	schema := &data.Schema{
		JSONSchema: jsonSchema,
	}

	schemaID, err := options.Backend.SetSchemaID(context.Background(), schema)
	if err != nil {
		panic(err)
	}

	stream := twitter.NewStream(options.StreamOptions)
	for {
		select {
		case <-options.Done:
			stream.Stop()
			return

		case tweet := <-stream.Tweets:
			valid, errs, err := jsonSchema.Validate(tweet)
			if err != nil {
				log.Println("!", "JSON schema error:", err)
			}

			if !valid {
				log.Println("!", "invalid tweet:", errs)
			}

			timestamp := time.Now()
			tweetData := new(bytes.Buffer)
			if err := json.NewEncoder(tweetData).Encode(tweet); err != nil {
				log.Println("!", "tweet error:", err)
				continue
			}

			coordinates := twitter.Coordinates(tweet)
			location := data.NewLocation([]float64{coordinates[0], coordinates[1]})

			ctx := context.Background()
			twitterAccountSources, err := options.Backend.ListTwitterAccountSourcesByAccountID(ctx, tweet.User.ID)
			if err != nil {
				log.Println("!", "twitter account sources error:", err)
				continue
			}

			for _, twitterAccountSource := range twitterAccountSources {
				record := &data.Record{
					SchemaID:  schemaID,
					SourceID:  twitterAccountSource.ID,
					TeamID:    twitterAccountSource.TeamID,
					UserID:    twitterAccountSource.UserID,
					Timestamp: timestamp,
					Location:  location,
				}

				if err := record.Data.UnmarshalJSON(tweetData.Bytes()); err != nil {
					log.Println("!", "record data error:", err)
				}

				if err := options.Backend.AddRecord(ctx, record); err != nil {
					log.Println("!", "record error:", err)
				}
			}
		}
	}
}
