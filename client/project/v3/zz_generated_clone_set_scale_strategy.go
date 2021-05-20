package client

const (
	CloneSetScaleStrategyType              = "cloneSetScaleStrategy"
	CloneSetScaleStrategyFieldPodsToDelete = "podsToDelete"
)

type CloneSetScaleStrategy struct {
	PodsToDelete []string `json:"podsToDelete,omitempty" yaml:"podsToDelete,omitempty"`
}
