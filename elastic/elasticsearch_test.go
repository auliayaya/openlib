package elastickuy

import (
	"context"
	"github.com/auliayaya/openlib/env"
	"github.com/auliayaya/openlib/logger"
	"testing"
)

func TestNewElasticMaster(t *testing.T) {
	abc := NewElasticMaster(env.NewEnv(), logger.NewLogger(env.NewEnv()))
	ctx := context.Background()
	abc.CheckAssetByDocumentID(ctx, RequestElastic{})
}