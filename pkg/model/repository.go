package model

import (
	"context"
	"sort"

	"emperror.dev/errors"
	"github.com/spaghettifunk/vector-operator/api/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	corev1 "k8s.io/api/core/v1"
)

func NewVectorResourceRepository(client client.Reader) *VectorResourceRepository {
	return &VectorResourceRepository{
		Client: client,
	}
}

type VectorResourceRepository struct {
	Client client.Reader
}

func (v VectorResourceRepository) VectorResourcesFor(ctx context.Context, vector v1alpha1.Vector) (res VectorResources, errs error) {
	res.Vector = vector

	watchNamespaces := vector.Spec.WatchNamespaces
	if len(watchNamespaces) == 0 {
		var nsList corev1.NamespaceList
		if err := v.Client.List(ctx, &nsList); err != nil {
			errs = errors.Append(errs, errors.WrapIf(err, "listing namespaces"))
			return
		}

		for _, i := range nsList.Items {
			watchNamespaces = append(watchNamespaces, i.Name)
		}
	}
	sort.Strings(watchNamespaces)

	return
}
