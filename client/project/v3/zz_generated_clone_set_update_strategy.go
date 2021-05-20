package client

import "k8s.io/apimachinery/pkg/util/intstr"

const (
	CloneSetUpdateStrategyType                       = "cloneSetUpdateStrategy"
	CloneSetUpdateStrategyFieldInPlaceUpdateStrategy = "inPlaceUpdateStrategy"
	CloneSetUpdateStrategyFieldMaxSurge              = "maxSurge"
	CloneSetUpdateStrategyFieldMaxUnavailable        = "maxUnavailable"
	CloneSetUpdateStrategyFieldPartition             = "partition"
	CloneSetUpdateStrategyFieldPaused                = "paused"
	CloneSetUpdateStrategyFieldPriorityStrategy      = "priorityStrategy"
	CloneSetUpdateStrategyFieldScatterStrategy       = "scatterStrategy"
	CloneSetUpdateStrategyFieldType                  = "type"
)

type CloneSetUpdateStrategy struct {
	InPlaceUpdateStrategy *InPlaceUpdateStrategy  `json:"inPlaceUpdateStrategy,omitempty" yaml:"inPlaceUpdateStrategy,omitempty"`
	MaxSurge              intstr.IntOrString      `json:"maxSurge,omitempty" yaml:"maxSurge,omitempty"`
	MaxUnavailable        intstr.IntOrString      `json:"maxUnavailable,omitempty" yaml:"maxUnavailable,omitempty"`
	Partition             intstr.IntOrString      `json:"partition,omitempty" yaml:"partition,omitempty"`
	Paused                bool                    `json:"paused,omitempty" yaml:"paused,omitempty"`
	PriorityStrategy      *UpdatePriorityStrategy `json:"priorityStrategy,omitempty" yaml:"priorityStrategy,omitempty"`
	ScatterStrategy       []UpdateScatterTerm     `json:"scatterStrategy,omitempty" yaml:"scatterStrategy,omitempty"`
	Type                  string                  `json:"type,omitempty" yaml:"type,omitempty"`
}
