package client

import (
	"github.com/rancher/norman/types"
)

const (
	CloneSetType                      = "cloneSet"
	CloneSetFieldAnnotations          = "annotations"
	CloneSetFieldCreated              = "created"
	CloneSetFieldCreatorID            = "creatorId"
	CloneSetFieldLabels               = "labels"
	CloneSetFieldLifecycle            = "lifecycle"
	CloneSetFieldMinReadySeconds      = "minReadySeconds"
	CloneSetFieldName                 = "name"
	CloneSetFieldNamespaceId          = "namespaceId"
	CloneSetFieldOwnerReferences      = "ownerReferences"
	CloneSetFieldProjectID            = "projectId"
	CloneSetFieldRemoved              = "removed"
	CloneSetFieldReplicas             = "replicas"
	CloneSetFieldRevisionHistoryLimit = "revisionHistoryLimit"
	CloneSetFieldScaleStrategy        = "scaleStrategy"
	CloneSetFieldSelector             = "selector"
	CloneSetFieldState                = "state"
	CloneSetFieldStatus               = "status"
	CloneSetFieldTemplate             = "template"
	CloneSetFieldTransitioning        = "transitioning"
	CloneSetFieldTransitioningMessage = "transitioningMessage"
	CloneSetFieldUUID                 = "uuid"
	CloneSetFieldUpdateStrategy       = "updateStrategy"
	CloneSetFieldVolumeClaimTemplates = "volumeClaimTemplates"
)

type CloneSet struct {
	types.Resource
	Annotations          map[string]string       `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Created              string                  `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID            string                  `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Labels               map[string]string       `json:"labels,omitempty" yaml:"labels,omitempty"`
	Lifecycle            *Lifecycle              `json:"lifecycle,omitempty" yaml:"lifecycle,omitempty"`
	MinReadySeconds      int64                   `json:"minReadySeconds,omitempty" yaml:"minReadySeconds,omitempty"`
	Name                 string                  `json:"name,omitempty" yaml:"name,omitempty"`
	NamespaceId          string                  `json:"namespaceId,omitempty" yaml:"namespaceId,omitempty"`
	OwnerReferences      []OwnerReference        `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	ProjectID            string                  `json:"projectId,omitempty" yaml:"projectId,omitempty"`
	Removed              string                  `json:"removed,omitempty" yaml:"removed,omitempty"`
	Replicas             *int64                  `json:"replicas,omitempty" yaml:"replicas,omitempty"`
	RevisionHistoryLimit *int64                  `json:"revisionHistoryLimit,omitempty" yaml:"revisionHistoryLimit,omitempty"`
	ScaleStrategy        *CloneSetScaleStrategy  `json:"scaleStrategy,omitempty" yaml:"scaleStrategy,omitempty"`
	Selector             *LabelSelector          `json:"selector,omitempty" yaml:"selector,omitempty"`
	State                string                  `json:"state,omitempty" yaml:"state,omitempty"`
	Status               interface{}             `json:"status,omitempty" yaml:"status,omitempty"`
	Template             *PodTemplateSpec        `json:"template,omitempty" yaml:"template,omitempty"`
	Transitioning        string                  `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`
	TransitioningMessage string                  `json:"transitioningMessage,omitempty" yaml:"transitioningMessage,omitempty"`
	UUID                 string                  `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	UpdateStrategy       *CloneSetUpdateStrategy `json:"updateStrategy,omitempty" yaml:"updateStrategy,omitempty"`
	VolumeClaimTemplates []PersistentVolumeClaim `json:"volumeClaimTemplates,omitempty" yaml:"volumeClaimTemplates,omitempty"`
}

type CloneSetCollection struct {
	types.Collection
	Data   []CloneSet `json:"data,omitempty"`
	client *CloneSetClient
}

type CloneSetClient struct {
	apiClient *Client
}

type CloneSetOperations interface {
	List(opts *types.ListOpts) (*CloneSetCollection, error)
	ListAll(opts *types.ListOpts) (*CloneSetCollection, error)
	Create(opts *CloneSet) (*CloneSet, error)
	Update(existing *CloneSet, updates interface{}) (*CloneSet, error)
	Replace(existing *CloneSet) (*CloneSet, error)
	ByID(id string) (*CloneSet, error)
	Delete(container *CloneSet) error
}

func newCloneSetClient(apiClient *Client) *CloneSetClient {
	return &CloneSetClient{
		apiClient: apiClient,
	}
}

func (c *CloneSetClient) Create(container *CloneSet) (*CloneSet, error) {
	resp := &CloneSet{}
	err := c.apiClient.Ops.DoCreate(CloneSetType, container, resp)
	return resp, err
}

func (c *CloneSetClient) Update(existing *CloneSet, updates interface{}) (*CloneSet, error) {
	resp := &CloneSet{}
	err := c.apiClient.Ops.DoUpdate(CloneSetType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *CloneSetClient) Replace(obj *CloneSet) (*CloneSet, error) {
	resp := &CloneSet{}
	err := c.apiClient.Ops.DoReplace(CloneSetType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *CloneSetClient) List(opts *types.ListOpts) (*CloneSetCollection, error) {
	resp := &CloneSetCollection{}
	err := c.apiClient.Ops.DoList(CloneSetType, opts, resp)
	resp.client = c
	return resp, err
}

func (c *CloneSetClient) ListAll(opts *types.ListOpts) (*CloneSetCollection, error) {
	resp := &CloneSetCollection{}
	resp, err := c.List(opts)
	if err != nil {
		return resp, err
	}
	data := resp.Data
	for next, err := resp.Next(); next != nil && err == nil; next, err = next.Next() {
		data = append(data, next.Data...)
		resp = next
		resp.Data = data
	}
	if err != nil {
		return resp, err
	}
	return resp, err
}

func (cc *CloneSetCollection) Next() (*CloneSetCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &CloneSetCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *CloneSetClient) ByID(id string) (*CloneSet, error) {
	resp := &CloneSet{}
	err := c.apiClient.Ops.DoByID(CloneSetType, id, resp)
	return resp, err
}

func (c *CloneSetClient) Delete(container *CloneSet) error {
	return c.apiClient.Ops.DoResourceDelete(CloneSetType, &container.Resource)
}
