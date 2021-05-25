package client

const (
	UpdatePriorityOrderTermType            = "updatePriorityOrderTerm"
	UpdatePriorityOrderTermFieldOrderedKey = "orderedKey"
)

type UpdatePriorityOrderTerm struct {
	OrderedKey string `json:"orderedKey,omitempty" yaml:"orderedKey,omitempty"`
}
