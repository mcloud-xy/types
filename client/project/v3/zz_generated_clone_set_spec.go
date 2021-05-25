package client

const (
	CloneSetSpecType                      = "cloneSetSpec"
	CloneSetSpecFieldLifecycle            = "lifecycle"
	CloneSetSpecFieldMinReadySeconds      = "minReadySeconds"
	CloneSetSpecFieldReplicas             = "replicas"
	CloneSetSpecFieldRevisionHistoryLimit = "revisionHistoryLimit"
	CloneSetSpecFieldScaleStrategy        = "scaleStrategy"
	CloneSetSpecFieldSelector             = "selector"
	CloneSetSpecFieldTemplate             = "template"
	CloneSetSpecFieldUpdateStrategy       = "updateStrategy"
	CloneSetSpecFieldVolumeClaimTemplates = "volumeClaimTemplates"
)

type CloneSetSpec struct {
	Lifecycle            *Lifecycle              `json:"lifecycle,omitempty" yaml:"lifecycle,omitempty"`
	MinReadySeconds      int64                   `json:"minReadySeconds,omitempty" yaml:"minReadySeconds,omitempty"`
	Replicas             *int64                  `json:"replicas,omitempty" yaml:"replicas,omitempty"`
	RevisionHistoryLimit *int64                  `json:"revisionHistoryLimit,omitempty" yaml:"revisionHistoryLimit,omitempty"`
	ScaleStrategy        *CloneSetScaleStrategy  `json:"scaleStrategy,omitempty" yaml:"scaleStrategy,omitempty"`
	Selector             *LabelSelector          `json:"selector,omitempty" yaml:"selector,omitempty"`
	Template             *PodTemplateSpec        `json:"template,omitempty" yaml:"template,omitempty"`
	UpdateStrategy       *CloneSetUpdateStrategy `json:"updateStrategy,omitempty" yaml:"updateStrategy,omitempty"`
	VolumeClaimTemplates []PersistentVolumeClaim `json:"volumeClaimTemplates,omitempty" yaml:"volumeClaimTemplates,omitempty"`
}
