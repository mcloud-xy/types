package client

const (
	AlertStatusType              = "alertStatus"
	AlertStatusFieldAlertState   = "alertState"
	AlertStatusFieldDesiredState = "desiredState"
)

type AlertStatus struct {
	AlertState   string `json:"alertState,omitempty" yaml:"alertState,omitempty"`
	DesiredState string `json:"desiredState,omitempty" yaml:"desiredState,omitempty"`
}
