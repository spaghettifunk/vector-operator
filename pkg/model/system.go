package model

import (
	"strconv"

	"github.com/go-logr/logr"
)

func CreateSystem(resources VectorResources, logger logr.Logger) (*types.System, error) {
	vector := resources.Vector

	router := types.NewRouter("main", types.Params{
		"metrics": strconv.FormatBool(vector.Spec.AggregatorSpec.Metrics != nil),
	})

	builder := types.NewSystemBuilder(rootInput, globalFilters, router)

	for _, flowCr := range resources.Flows {
		flow, err := FlowForFlow(flowCr, resources.ClusterOutputs, resources.Outputs, secrets)
		if err != nil {
			if vector.Spec.SkipInvalidResources {
				logger.Error(err, "Flow contains errors.")
			} else {
				return nil, err
			}
		}
		err = builder.RegisterFlow(flow)
		if err != nil {
			return nil, err
		}
	}	

	return system, err
}
