package inaturalist

import (
	"context"

	"go.uber.org/zap"

	"github.com/fieldkit/cloud/server/logging"
)

func Logger(ctx context.Context) *zap.Logger {
	return logging.Logger(ctx).Named("inaturalist")
}
