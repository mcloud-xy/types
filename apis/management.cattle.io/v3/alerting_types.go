package v3

import (
	"strings"

	"github.com/rancher/norman/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ClusterAlert struct {
	types.Namespaced

	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ClusterAlertSpec `json:"spec"`
	// Most recent observed status of the alert. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status AlertStatus `json:"status"`
}

func (c *ClusterAlert) ObjClusterName() string {
	return c.Spec.ObjClusterName()
}

type ProjectAlert struct {
	types.Namespaced

	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ProjectAlertSpec `json:"spec"`
	// Most recent observed status of the alert. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status AlertStatus `json:"status"`
}

func (p *ProjectAlert) ObjClusterName() string {
	return p.Spec.ObjClusterName()
}

type AlertCommonSpec struct {
	DisplayName           string      `json:"displayName,omitempty" norman:"required"`
	Description           string      `json:"description,omitempty"`
	Severity              string      `json:"severity,omitempty" norman:"required,options=info|critical|warning,default=critical"`
	Recipients            []Recipient `json:"recipients,omitempty" norman:"required"`
	InitialWaitSeconds    int         `json:"initialWaitSeconds,omitempty" norman:"required,default=180,min=0"`
	RepeatIntervalSeconds int         `json:"repeatIntervalSeconds,omitempty"  norman:"required,default=3600,min=0"`
}

type ClusterAlertSpec struct {
	AlertCommonSpec

	ClusterName         string               `json:"clusterName" norman:"type=reference[cluster]"`
	TargetNode          *TargetNode          `json:"targetNode,omitempty"`
	TargetSystemService *TargetSystemService `json:"targetSystemService,omitempty"`
	TargetEvent         *TargetEvent         `json:"targetEvent,omitempty"`
}

func (c *ClusterAlertSpec) ObjClusterName() string {
	return c.ClusterName
}

type ProjectAlertSpec struct {
	AlertCommonSpec

	ProjectName    string          `json:"projectName" norman:"type=reference[project]"`
	TargetWorkload *TargetWorkload `json:"targetWorkload,omitempty"`
	TargetPod      *TargetPod      `json:"targetPod,omitempty"`
}

func (p *ProjectAlertSpec) ObjClusterName() string {
	if parts := strings.SplitN(p.ProjectName, ":", 2); len(parts) == 2 {
		return parts[0]
	}
	return ""
}

type Recipient struct {
	Recipient    string `json:"recipient,omitempty"`
	NotifierName string `json:"notifierName,omitempty" norman:"required,type=reference[notifier]"`
	NotifierType string `json:"notifierType,omitempty" norman:"required,options=slack|email|pagerduty|webhook|wechat"`
}

type TargetNode struct {
	NodeName     string            `json:"nodeName,omitempty" norman:"type=reference[node]"`
	Selector     map[string]string `json:"selector,omitempty"`
	Condition    string            `json:"condition,omitempty" norman:"required,options=notready|mem|cpu,default=notready"`
	MemThreshold int               `json:"memThreshold,omitempty" norman:"min=1,max=100,default=70"`
	CPUThreshold int               `json:"cpuThreshold,omitempty" norman:"min=1,default=70"`
}

type TargetPod struct {
	PodName                string `json:"podName,omitempty" norman:"required,type=reference[/v3/projects/schemas/pod]"`
	Condition              string `json:"condition,omitempty" norman:"required,options=notrunning|notscheduled|restarts,default=notrunning"`
	RestartTimes           int    `json:"restartTimes,omitempty" norman:"min=1,default=3"`
	RestartIntervalSeconds int    `json:"restartIntervalSeconds,omitempty"  norman:"min=1,default=300"`
}

type TargetEvent struct {
	EventType    string `json:"eventType,omitempty" norman:"required,options=Normal|Warning,default=Warning"`
	ResourceKind string `json:"resourceKind,omitempty" norman:"required,options=Pod|Node|Deployment|StatefulSet|DaemonSet"`
}

type TargetWorkload struct {
	WorkloadID          string            `json:"workloadId,omitempty"`
	Selector            map[string]string `json:"selector,omitempty"`
	AvailablePercentage int               `json:"availablePercentage,omitempty" norman:"required,min=1,max=100,default=70"`
}

type TargetSystemService struct {
	Condition string `json:"condition,omitempty" norman:"required,options=etcd|controller-manager|scheduler,default=scheduler"`
}

type AlertStatus struct {
	AlertState string `json:"alertState,omitempty" norman:"options=active|inactive|alerting|muted,default=active"`

	// 设置的状态。不受实时告警状态影响。
	DesiredState string `json:"desiredState,omitempty" norman:"options=active|inactive|muted"`
}

type ClusterAlertGroup struct {
	types.Namespaced

	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ClusterGroupSpec `json:"spec"`
	// Most recent observed status of the alert. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status AlertStatus `json:"status"`
}

func (c *ClusterAlertGroup) ObjClusterName() string {
	return c.Spec.ObjClusterName()
}

type ProjectAlertGroup struct {
	types.Namespaced

	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ProjectGroupSpec `json:"spec"`
	// Most recent observed status of the alert. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status AlertStatus `json:"status"`
}

func (p *ProjectAlertGroup) ObjClusterName() string {
	return p.Spec.ObjClusterName()
}

type ClusterGroupSpec struct {
	ClusterName string      `json:"clusterName" norman:"type=reference[cluster]"`
	Recipients  []Recipient `json:"recipients,omitempty"`
	CommonGroupField
}

func (c *ClusterGroupSpec) ObjClusterName() string {
	return c.ClusterName
}

type ProjectGroupSpec struct {
	ProjectName string      `json:"projectName" norman:"type=reference[project]"`
	Recipients  []Recipient `json:"recipients,omitempty"`
	CommonGroupField
}

func (p *ProjectGroupSpec) ObjClusterName() string {
	if parts := strings.SplitN(p.ProjectName, ":", 2); len(parts) == 2 {
		return parts[0]
	}
	return ""
}

type ClusterAlertRule struct {
	types.Namespaced

	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ClusterAlertRuleSpec `json:"spec"`
	// Most recent observed status of the alert. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status AlertStatus `json:"status"`
}

func (c *ClusterAlertRule) ObjClusterName() string {
	return c.Spec.ObjClusterName()
}

type ClusterAlertRuleSpec struct {
	CommonRuleField
	ClusterName       string             `json:"clusterName" norman:"type=reference[cluster]"`
	GroupName         string             `json:"groupName" norman:"type=reference[clusterAlertGroup]"`
	NodeRule          *NodeRule          `json:"nodeRule,omitempty"`
	EventRule         *EventRule         `json:"eventRule,omitempty"`
	SystemServiceRule *SystemServiceRule `json:"systemServiceRule,omitempty"`
	MetricRule        *MetricRule        `json:"metricRule,omitempty"`
	ClusterScanRule   *ClusterScanRule   `json:"clusterScanRule,omitempty"`
	DeploymentRule    *DeploymentRule    `json:"deploymentRule,omitempty"`
	TargetType        string             `json:"targetType" norman:"options=deployment|node|metric"`
	TargetObject      map[string]string  `json:"targetObject"`
	TargetObjects     map[string]string  `json:"targetObjects"`
	Rule              *MetricRuleV2      `json:"rule,omitempty"`
	Scope             string             `json:"scope"`
	IDC               string             `json:"idc"`
}

func (c *ClusterAlertRuleSpec) ObjClusterName() string {
	return c.ClusterName
}

type ProjectAlertRule struct {
	types.Namespaced

	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ProjectAlertRuleSpec `json:"spec"`
	// Most recent observed status of the alert. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status AlertStatus `json:"status"`
}

func (p *ProjectAlertRule) ObjClusterName() string {
	return p.Spec.ObjClusterName()
}

type ProjectAlertRuleSpec struct {
	CommonRuleField
	ProjectName  string        `json:"projectName" norman:"type=reference[project]"`
	GroupName    string        `json:"groupName" norman:"type=reference[projectAlertGroup]"`
	PodRule      *PodRule      `json:"podRule,omitempty"`
	WorkloadRule *WorkloadRule `json:"workloadRule,omitempty"`
	MetricRule   *MetricRule   `json:"metricRule,omitempty"`
}

func (p *ProjectAlertRuleSpec) ObjClusterName() string {
	if parts := strings.SplitN(p.ProjectName, ":", 2); len(parts) == 2 {
		return parts[0]
	}
	return ""
}

type CommonGroupField struct {
	DisplayName string `json:"displayName,omitempty" norman:"required"`
	Description string `json:"description,omitempty"`
	TimingField
}

type CommonRuleField struct {
	DisplayName string `json:"displayName,omitempty"`
	Severity    string `json:"severity,omitempty" norman:"required,options=info|critical|warning,default=critical"`
	Inherited   *bool  `json:"inherited,omitempty" norman:"default=true"`
	TimingField
}

type ClusterScanRule struct {
	ScanRunType  ClusterScanRunType `json:"scanRunType,omitempty" norman:"required,options=manual|scheduled,default=scheduled"`
	FailuresOnly bool               `json:"failuresOnly,omitempty"`
}

type MetricRule struct {
	Expression     string  `json:"expression,omitempty" norman:"required"`
	Description    string  `json:"description,omitempty"`
	Duration       string  `json:"duration,omitempty" norman:"required"`
	Comparison     string  `json:"comparison,omitempty" norman:"type=enum,options=equal|not-equal|greater-than|less-than|greater-or-equal|less-or-equal|has-value,default=equal"`
	ThresholdValue float64 `json:"thresholdValue,omitempty" norman:"type=float"`
}

type TimingField struct {
	GroupWaitSeconds      int `json:"groupWaitSeconds,omitempty" norman:"required,default=30,min=1"`
	GroupIntervalSeconds  int `json:"groupIntervalSeconds,omitempty" norman:"required,default=180,min=1"`
	RepeatIntervalSeconds int `json:"repeatIntervalSeconds,omitempty"  norman:"required,default=3600,min=1"`
}

type NodeRule struct {
	NodeName     string            `json:"nodeName,omitempty" norman:"type=reference[node]"`
	Selector     map[string]string `json:"selector,omitempty"`
	Condition    string            `json:"condition,omitempty" norman:"required,options=notready|mem|cpu,default=notready"`
	MemThreshold int               `json:"memThreshold,omitempty" norman:"min=1,max=100,default=70"`
	CPUThreshold int               `json:"cpuThreshold,omitempty" norman:"min=1,default=70"`
}

type DeploymentRule struct {
	Condition           string            `json:"condition,omitempty" norman:"options=hasRestart"`
	MemThreshold        int               `json:"memThreshold,omitempty" norman:"min=0,max=100"`
	CPUThreshold        int               `json:"cpuThreshold,omitempty" norman:"min=0,max=100"`
	DeploymentName      string            `json:"deploymentName,omitempty"`
	DeploymentNamespace string            `json:"deploymentNamespace,omitempty"`
	Selector            map[string]string `json:"selector,omitempty"`
}

type MetricRuleV2 struct {
	Description    string  `json:"description,omitempty"`
	Duration       string  `json:"duration,omitempty" norman:"required"`
	Comparison     string  `json:"comparison,omitempty" norman:"type=enum,options=equal|not-equal|greater-than|less-than|greater-or-equal|less-or-equal|has-value,default=equal"`
	ThresholdValue float64 `json:"thresholdValue,omitempty" norman:"type=float"`
	Interval       string  `json:"interval" norman:"type=string"`
	Metric         string  `json:"metric" norman:"required"`
	MetricName     string  `json:"metricName" norman:"required"`
	Operator       string  `json:"operator"`
}

type PodRule struct {
	PodName                string `json:"podName,omitempty" norman:"required,type=reference[/v3/projects/schemas/pod]"`
	Condition              string `json:"condition,omitempty" norman:"required,options=notrunning|notscheduled|restarts,default=notrunning"`
	RestartTimes           int    `json:"restartTimes,omitempty" norman:"min=1,default=3"`
	RestartIntervalSeconds int    `json:"restartIntervalSeconds,omitempty"  norman:"min=1,default=300"`
}

type EventRule struct {
	EventType    string `json:"eventType,omitempty" norman:"required,options=Normal|Warning,default=Warning"`
	ResourceKind string `json:"resourceKind,omitempty" norman:"required,options=Pod|Node|Deployment|StatefulSet|DaemonSet"`
}

type WorkloadRule struct {
	WorkloadID          string            `json:"workloadId,omitempty"`
	Selector            map[string]string `json:"selector,omitempty"`
	AvailablePercentage int               `json:"availablePercentage,omitempty" norman:"required,min=1,max=100,default=70"`
}

type SystemServiceRule struct {
	Condition string `json:"condition,omitempty" norman:"required,options=etcd|controller-manager|scheduler,default=scheduler"`
}

type Notifier struct {
	types.Namespaced

	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec NotifierSpec `json:"spec"`
	// Most recent observed status of the notifier. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status NotifierStatus `json:"status"`
}

func (n *Notifier) ObjClusterName() string {
	return n.Spec.ObjClusterName()
}

type NotifierSpec struct {
	ClusterName string `json:"clusterName" norman:"type=reference[cluster]"`

	DisplayName     string           `json:"displayName,omitempty" norman:"required"`
	Description     string           `json:"description,omitempty"`
	SendResolved    bool             `json:"sendResolved,omitempty"`
	SMTPConfig      *SMTPConfig      `json:"smtpConfig,omitempty"`
	SlackConfig     *SlackConfig     `json:"slackConfig,omitempty"`
	PagerdutyConfig *PagerdutyConfig `json:"pagerdutyConfig,omitempty"`
	WebhookConfig   *WebhookConfig   `json:"webhookConfig,omitempty"`
	WechatConfig    *WechatConfig    `json:"wechatConfig,omitempty"`
}

func (n *NotifierSpec) ObjClusterName() string {
	return n.ClusterName
}

type Notification struct {
	Message         string           `json:"message,omitempty"`
	SMTPConfig      *SMTPConfig      `json:"smtpConfig,omitempty"`
	SlackConfig     *SlackConfig     `json:"slackConfig,omitempty"`
	PagerdutyConfig *PagerdutyConfig `json:"pagerdutyConfig,omitempty"`
	WebhookConfig   *WebhookConfig   `json:"webhookConfig,omitempty"`
	WechatConfig    *WechatConfig    `json:"wechatConfig,omitempty"`
}

type SMTPConfig struct {
	Host             string `json:"host,omitempty" norman:"required,type=hostname"`
	Port             int    `json:"port,omitempty" norman:"required,min=1,max=65535,default=587"`
	Username         string `json:"username,omitempty"`
	Password         string `json:"password,omitempty" norman:"type=password"`
	Sender           string `json:"sender,omitempty" norman:"required"`
	DefaultRecipient string `json:"defaultRecipient,omitempty" norman:"required"`
	TLS              *bool  `json:"tls,omitempty" norman:"required,default=true"`
}

type SlackConfig struct {
	DefaultRecipient string `json:"defaultRecipient,omitempty"`
	URL              string `json:"url,omitempty" norman:"required"`
	*HTTPClientConfig
}

type PagerdutyConfig struct {
	ServiceKey string `json:"serviceKey,omitempty" norman:"required"`
	*HTTPClientConfig
}

type WebhookConfig struct {
	URL string `json:"url,omitempty" norman:"required"`
	*HTTPClientConfig
}

type WechatConfig struct {
	DefaultRecipient string `json:"defaultRecipient,omitempty" norman:"required"`
	Secret           string `json:"secret,omitempty" norman:"type=password,required"`
	Agent            string `json:"agent,omitempty" norman:"required"`
	Corp             string `json:"corp,omitempty" norman:"required"`
	RecipientType    string `json:"recipientType,omitempty" norman:"required,options=tag|party|user,default=party"`
	APIURL           string `json:"apiUrl,omitempty"`
	*HTTPClientConfig
}

type NotifierStatus struct {
}

// HTTPClientConfig configures an HTTP client.
type HTTPClientConfig struct {
	// HTTP proxy server to use to connect to the targets.
	ProxyURL string `json:"proxyUrl,omitempty"`
}

// 告警记录
type AlertRecord struct {

	// 告警规则相关
	AlertName string `json:"alertName"`
	// 状态
	AlertState string `json:"alertState"`

	// 集群相关
	ClusterName string `json:"clusterName"`
	ClusterID   string `json:"clusterId"`
	IDC         string `json:"idc"`

	RuleID   string `json:"ruleId"`
	Severity string `json:"severity"`

	// 告警条件
	MetricName     string  `json:"metricName"`
	Metric         string  `json:"metric"`
	Expression     string  `json:"expression"`
	Comparison     string  `json:"comparison"`
	Duration       string  `json:"duration"`
	ThresholdValue float64 `json:"thresholdValue"`
	CurrentValue   float64 `json:"currentValue"`

	// 告警对象
	TargetType   string `json:"targetType"`
	TargetObject string `json:"targetObject"`

	// 告警持续时长。
	// Duration int64  计算字段
	StartTs int64 `json:"startTs"`
	EndTs   int64 `json:"endTs"`
	MuteTs  int64 `json:"muteTs"`

	//接收人
	GroupID    string      `json:"groupId"`
	Recipients []Recipient `json:"recipients"`
	GroupName  string      `json:"groupName"` // 告警组
}
