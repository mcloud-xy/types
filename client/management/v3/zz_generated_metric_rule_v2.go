package client

const (
	MetricRuleV2Type                = "metricRuleV2"
	MetricRuleV2FieldComparison     = "comparison"
	MetricRuleV2FieldDescription    = "description"
	MetricRuleV2FieldDuration       = "duration"
	MetricRuleV2FieldInterval       = "interval"
	MetricRuleV2FieldMetric         = "metric"
	MetricRuleV2FieldMetricName     = "metricName"
	MetricRuleV2FieldOperator       = "operator"
	MetricRuleV2FieldThresholdValue = "thresholdValue"
)

type MetricRuleV2 struct {
	Comparison     string  `json:"comparison,omitempty" yaml:"comparison,omitempty"`
	Description    string  `json:"description,omitempty" yaml:"description,omitempty"`
	Duration       string  `json:"duration,omitempty" yaml:"duration,omitempty"`
	Interval       string  `json:"interval,omitempty" yaml:"interval,omitempty"`
	Metric         string  `json:"metric,omitempty" yaml:"metric,omitempty"`
	MetricName     string  `json:"metricName,omitempty" yaml:"metricName,omitempty"`
	Operator       string  `json:"operator,omitempty" yaml:"operator,omitempty"`
	ThresholdValue float64 `json:"thresholdValue,omitempty" yaml:"thresholdValue,omitempty"`
}
