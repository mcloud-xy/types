package client

const (
	DeploymentRuleType                     = "deploymentRule"
	DeploymentRuleFieldCPUThreshold        = "cpuThreshold"
	DeploymentRuleFieldCondition           = "condition"
	DeploymentRuleFieldDeploymentName      = "deploymentName"
	DeploymentRuleFieldDeploymentNamespace = "deploymentNamespace"
	DeploymentRuleFieldMemThreshold        = "memThreshold"
	DeploymentRuleFieldSelector            = "selector"
)

type DeploymentRule struct {
	CPUThreshold        int64             `json:"cpuThreshold,omitempty" yaml:"cpuThreshold,omitempty"`
	Condition           string            `json:"condition,omitempty" yaml:"condition,omitempty"`
	DeploymentName      string            `json:"deploymentName,omitempty" yaml:"deploymentName,omitempty"`
	DeploymentNamespace string            `json:"deploymentNamespace,omitempty" yaml:"deploymentNamespace,omitempty"`
	MemThreshold        int64             `json:"memThreshold,omitempty" yaml:"memThreshold,omitempty"`
	Selector            map[string]string `json:"selector,omitempty" yaml:"selector,omitempty"`
}
