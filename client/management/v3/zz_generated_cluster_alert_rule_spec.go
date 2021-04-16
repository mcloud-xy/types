package client

const (
	ClusterAlertRuleSpecType                       = "clusterAlertRuleSpec"
	ClusterAlertRuleSpecFieldClusterID             = "clusterId"
	ClusterAlertRuleSpecFieldClusterScanRule       = "clusterScanRule"
	ClusterAlertRuleSpecFieldDeploymentRule        = "deploymentRule"
	ClusterAlertRuleSpecFieldDisplayName           = "displayName"
	ClusterAlertRuleSpecFieldEventRule             = "eventRule"
	ClusterAlertRuleSpecFieldGroupID               = "groupId"
	ClusterAlertRuleSpecFieldGroupIntervalSeconds  = "groupIntervalSeconds"
	ClusterAlertRuleSpecFieldGroupWaitSeconds      = "groupWaitSeconds"
	ClusterAlertRuleSpecFieldIDC                   = "idc"
	ClusterAlertRuleSpecFieldInherited             = "inherited"
	ClusterAlertRuleSpecFieldMetricRule            = "metricRule"
	ClusterAlertRuleSpecFieldNodeRule              = "nodeRule"
	ClusterAlertRuleSpecFieldRepeatIntervalSeconds = "repeatIntervalSeconds"
	ClusterAlertRuleSpecFieldRule                  = "rule"
	ClusterAlertRuleSpecFieldScope                 = "scope"
	ClusterAlertRuleSpecFieldSeverity              = "severity"
	ClusterAlertRuleSpecFieldSystemServiceRule     = "systemServiceRule"
	ClusterAlertRuleSpecFieldTargetObject          = "targetObject"
	ClusterAlertRuleSpecFieldTargetObjects         = "targetObjects"
	ClusterAlertRuleSpecFieldTargetType            = "targetType"
)

type ClusterAlertRuleSpec struct {
	ClusterID             string             `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`
	ClusterScanRule       *ClusterScanRule   `json:"clusterScanRule,omitempty" yaml:"clusterScanRule,omitempty"`
	DeploymentRule        *DeploymentRule    `json:"deploymentRule,omitempty" yaml:"deploymentRule,omitempty"`
	DisplayName           string             `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	EventRule             *EventRule         `json:"eventRule,omitempty" yaml:"eventRule,omitempty"`
	GroupID               string             `json:"groupId,omitempty" yaml:"groupId,omitempty"`
	GroupIntervalSeconds  int64              `json:"groupIntervalSeconds,omitempty" yaml:"groupIntervalSeconds,omitempty"`
	GroupWaitSeconds      int64              `json:"groupWaitSeconds,omitempty" yaml:"groupWaitSeconds,omitempty"`
	IDC                   string             `json:"idc,omitempty" yaml:"idc,omitempty"`
	Inherited             *bool              `json:"inherited,omitempty" yaml:"inherited,omitempty"`
	MetricRule            *MetricRule        `json:"metricRule,omitempty" yaml:"metricRule,omitempty"`
	NodeRule              *NodeRule          `json:"nodeRule,omitempty" yaml:"nodeRule,omitempty"`
	RepeatIntervalSeconds int64              `json:"repeatIntervalSeconds,omitempty" yaml:"repeatIntervalSeconds,omitempty"`
	Rule                  *MetricRuleV2      `json:"rule,omitempty" yaml:"rule,omitempty"`
	Scope                 string             `json:"scope,omitempty" yaml:"scope,omitempty"`
	Severity              string             `json:"severity,omitempty" yaml:"severity,omitempty"`
	SystemServiceRule     *SystemServiceRule `json:"systemServiceRule,omitempty" yaml:"systemServiceRule,omitempty"`
	TargetObject          map[string]string  `json:"targetObject,omitempty" yaml:"targetObject,omitempty"`
	TargetObjects         map[string]string  `json:"targetObjects,omitempty" yaml:"targetObjects,omitempty"`
	TargetType            string             `json:"targetType,omitempty" yaml:"targetType,omitempty"`
}
