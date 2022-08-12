package model

import "github.com/spaghettifunk/vector-operator/api/v1alpha1"

type VectorResources struct {
	Vector    v1alpha1.Vector
	Pipelines []v1alpha1.Pipeline
}
