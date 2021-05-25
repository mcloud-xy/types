package client

const (
	CloneSetStatusType                      = "cloneSetStatus"
	CloneSetStatusFieldAvailableReplicas    = "availableReplicas"
	CloneSetStatusFieldCollisionCount       = "collisionCount"
	CloneSetStatusFieldConditions           = "conditions"
	CloneSetStatusFieldCurrentRevision      = "currentRevision"
	CloneSetStatusFieldLabelSelector        = "labelSelector"
	CloneSetStatusFieldObservedGeneration   = "observedGeneration"
	CloneSetStatusFieldReadyReplicas        = "readyReplicas"
	CloneSetStatusFieldReplicas             = "replicas"
	CloneSetStatusFieldUpdateRevision       = "updateRevision"
	CloneSetStatusFieldUpdatedReadyReplicas = "updatedReadyReplicas"
	CloneSetStatusFieldUpdatedReplicas      = "updatedReplicas"
)

type CloneSetStatus struct {
	AvailableReplicas    int64               `json:"availableReplicas,omitempty" yaml:"availableReplicas,omitempty"`
	CollisionCount       *int64              `json:"collisionCount,omitempty" yaml:"collisionCount,omitempty"`
	Conditions           []CloneSetCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	CurrentRevision      string              `json:"currentRevision,omitempty" yaml:"currentRevision,omitempty"`
	LabelSelector        string              `json:"labelSelector,omitempty" yaml:"labelSelector,omitempty"`
	ObservedGeneration   int64               `json:"observedGeneration,omitempty" yaml:"observedGeneration,omitempty"`
	ReadyReplicas        int64               `json:"readyReplicas,omitempty" yaml:"readyReplicas,omitempty"`
	Replicas             int64               `json:"replicas,omitempty" yaml:"replicas,omitempty"`
	UpdateRevision       string              `json:"updateRevision,omitempty" yaml:"updateRevision,omitempty"`
	UpdatedReadyReplicas int64               `json:"updatedReadyReplicas,omitempty" yaml:"updatedReadyReplicas,omitempty"`
	UpdatedReplicas      int64               `json:"updatedReplicas,omitempty" yaml:"updatedReplicas,omitempty"`
}
