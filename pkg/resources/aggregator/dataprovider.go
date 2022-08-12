package aggregator

import (
	"context"

	"emperror.dev/errors"
	"github.com/spaghettifunk/vector-operator/api/v1alpha1"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type DataProvider struct {
	client client.Client
}

func NewDataProvider(client client.Client) *DataProvider {
	return &DataProvider{
		client: client,
	}
}

func (p *DataProvider) GetReplicaCount(ctx context.Context, vector *v1alpha1.Vector) (*int32, error) {
	sts := &v1.StatefulSet{}
	om := vector.AggregatorObjectMeta(StatefulSetName, ComponentAggregator)
	err := p.client.Get(ctx, types.NamespacedName{Namespace: om.Namespace, Name: om.Name}, sts)
	if err != nil {
		return nil, errors.WrapIf(client.IgnoreNotFound(err), "getting fluentd statefulset")
	}
	return sts.Spec.Replicas, nil
}
