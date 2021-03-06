package jobs

import (
	"context"

	"github.com/conservify/sqlxcache"
)

type QueueSystem struct {
	Queues map[string]*PgJobQueue
	Defs   map[string]*QueueDef
}

func OpenQueueSystem(ctx context.Context, url string, defs []*QueueDef) (qs *QueueSystem, err error) {
	queues := make(map[string]*PgJobQueue)
	defsMap := make(map[string]*QueueDef)

	db, err := sqlxcache.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	for _, def := range defs {
		jq, err := NewPqJobQueue(ctx, db, url, def.Name)
		if err != nil {
			return nil, err
		}
		queues[def.Name] = jq
		defsMap[def.Name] = def
	}

	qs = &QueueSystem{
		Queues: queues,
		Defs:   defsMap,
	}

	return
}
