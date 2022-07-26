package aggregatordataprovider

import (
	"context"

	"github.com/spaghettifunk/vector-operator/api/v1alpha1"
)

type AggregatorDataProvider interface {
	GetReplicaCount(ctx context.Context, vector *v1alpha1.Vector) (*int32, error)
}
