package client

const (
	InPlaceUpdateStrategyType                    = "inPlaceUpdateStrategy"
	InPlaceUpdateStrategyFieldGracePeriodSeconds = "gracePeriodSeconds"
)

type InPlaceUpdateStrategy struct {
	GracePeriodSeconds int64 `json:"gracePeriodSeconds,omitempty" yaml:"gracePeriodSeconds,omitempty"`
}
