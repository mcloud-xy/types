package client

const (
	UpdateScatterTermType       = "updateScatterTerm"
	UpdateScatterTermFieldKey   = "key"
	UpdateScatterTermFieldValue = "value"
)

type UpdateScatterTerm struct {
	Key   string `json:"key,omitempty" yaml:"key,omitempty"`
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
