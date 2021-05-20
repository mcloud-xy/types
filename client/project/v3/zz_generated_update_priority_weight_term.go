package client

const (
	UpdatePriorityWeightTermType               = "updatePriorityWeightTerm"
	UpdatePriorityWeightTermFieldMatchSelector = "matchSelector"
	UpdatePriorityWeightTermFieldWeight        = "weight"
)

type UpdatePriorityWeightTerm struct {
	MatchSelector *LabelSelector `json:"matchSelector,omitempty" yaml:"matchSelector,omitempty"`
	Weight        int64          `json:"weight,omitempty" yaml:"weight,omitempty"`
}
