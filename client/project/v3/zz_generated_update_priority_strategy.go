package client

const (
	UpdatePriorityStrategyType                = "updatePriorityStrategy"
	UpdatePriorityStrategyFieldOrderPriority  = "orderPriority"
	UpdatePriorityStrategyFieldWeightPriority = "weightPriority"
)

type UpdatePriorityStrategy struct {
	OrderPriority  []UpdatePriorityOrderTerm  `json:"orderPriority,omitempty" yaml:"orderPriority,omitempty"`
	WeightPriority []UpdatePriorityWeightTerm `json:"weightPriority,omitempty" yaml:"weightPriority,omitempty"`
}
