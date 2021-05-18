package v1alpha1

import (
	"context"
	"time"

	"github.com/openkruise/kruise-api/apps/v1alpha1"
	"github.com/rancher/norman/controller"
	"github.com/rancher/norman/objectclient"
	"github.com/rancher/norman/resource"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

var (
	CloneSetGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "CloneSet",
	}
	CloneSetResource = metav1.APIResource{
		Name:         "clonesets",
		SingularName: "cloneset",
		Namespaced:   true,

		Kind: CloneSetGroupVersionKind.Kind,
	}

	CloneSetGroupVersionResource = schema.GroupVersionResource{
		Group:    GroupName,
		Version:  Version,
		Resource: "clonesets",
	}
)

func init() {
	resource.Put(CloneSetGroupVersionResource)
}

func NewCloneSet(namespace, name string, obj v1alpha1.CloneSet) *v1alpha1.CloneSet {
	obj.APIVersion, obj.Kind = CloneSetGroupVersionKind.ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type CloneSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []v1alpha1.CloneSet `json:"items"`
}

type CloneSetHandlerFunc func(key string, obj *v1alpha1.CloneSet) (runtime.Object, error)

type CloneSetChangeHandlerFunc func(obj *v1alpha1.CloneSet) (runtime.Object, error)

type CloneSetLister interface {
	List(namespace string, selector labels.Selector) (ret []*v1alpha1.CloneSet, err error)
	Get(namespace, name string) (*v1alpha1.CloneSet, error)
}

type CloneSetController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() CloneSetLister
	AddHandler(ctx context.Context, name string, handler CloneSetHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync CloneSetHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler CloneSetHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, handler CloneSetHandlerFunc)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, after time.Duration)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type CloneSetInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*v1alpha1.CloneSet) (*v1alpha1.CloneSet, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v1alpha1.CloneSet, error)
	Get(name string, opts metav1.GetOptions) (*v1alpha1.CloneSet, error)
	Update(*v1alpha1.CloneSet) (*v1alpha1.CloneSet, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*CloneSetList, error)
	ListNamespaced(namespace string, opts metav1.ListOptions) (*CloneSetList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() CloneSetController
	AddHandler(ctx context.Context, name string, sync CloneSetHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync CloneSetHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle CloneSetLifecycle)
	AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle CloneSetLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync CloneSetHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync CloneSetHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle CloneSetLifecycle)
	AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle CloneSetLifecycle)
}

type cloneSetLister struct {
	controller *cloneSetController
}

func (l *cloneSetLister) List(namespace string, selector labels.Selector) (ret []*v1alpha1.CloneSet, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*v1alpha1.CloneSet))
	})
	return
}

func (l *cloneSetLister) Get(namespace, name string) (*v1alpha1.CloneSet, error) {
	var key string
	if namespace != "" {
		key = namespace + "/" + name
	} else {
		key = name
	}
	obj, exists, err := l.controller.Informer().GetIndexer().GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(schema.GroupResource{
			Group:    CloneSetGroupVersionKind.Group,
			Resource: "cloneSet",
		}, key)
	}
	return obj.(*v1alpha1.CloneSet), nil
}

type cloneSetController struct {
	controller.GenericController
}

func (c *cloneSetController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *cloneSetController) Lister() CloneSetLister {
	return &cloneSetLister{
		controller: c,
	}
}

func (c *cloneSetController) AddHandler(ctx context.Context, name string, handler CloneSetHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1alpha1.CloneSet); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *cloneSetController) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, handler CloneSetHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1alpha1.CloneSet); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *cloneSetController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler CloneSetHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1alpha1.CloneSet); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *cloneSetController) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, cluster string, handler CloneSetHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1alpha1.CloneSet); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type cloneSetFactory struct {
}

func (c cloneSetFactory) Object() runtime.Object {
	return &v1alpha1.CloneSet{}
}

func (c cloneSetFactory) List() runtime.Object {
	return &CloneSetList{}
}

func (s *cloneSetClient) Controller() CloneSetController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.cloneSetControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(CloneSetGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &cloneSetController{
		GenericController: genericController,
	}

	s.client.cloneSetControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type cloneSetClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   CloneSetController
}

func (s *cloneSetClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *cloneSetClient) Create(o *v1alpha1.CloneSet) (*v1alpha1.CloneSet, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*v1alpha1.CloneSet), err
}

func (s *cloneSetClient) Get(name string, opts metav1.GetOptions) (*v1alpha1.CloneSet, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*v1alpha1.CloneSet), err
}

func (s *cloneSetClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v1alpha1.CloneSet, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*v1alpha1.CloneSet), err
}

func (s *cloneSetClient) Update(o *v1alpha1.CloneSet) (*v1alpha1.CloneSet, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*v1alpha1.CloneSet), err
}

func (s *cloneSetClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *cloneSetClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *cloneSetClient) List(opts metav1.ListOptions) (*CloneSetList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*CloneSetList), err
}

func (s *cloneSetClient) ListNamespaced(namespace string, opts metav1.ListOptions) (*CloneSetList, error) {
	obj, err := s.objectClient.ListNamespaced(namespace, opts)
	return obj.(*CloneSetList), err
}

func (s *cloneSetClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *cloneSetClient) Patch(o *v1alpha1.CloneSet, patchType types.PatchType, data []byte, subresources ...string) (*v1alpha1.CloneSet, error) {
	obj, err := s.objectClient.Patch(o.Name, o, patchType, data, subresources...)
	return obj.(*v1alpha1.CloneSet), err
}

func (s *cloneSetClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *cloneSetClient) AddHandler(ctx context.Context, name string, sync CloneSetHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *cloneSetClient) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync CloneSetHandlerFunc) {
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *cloneSetClient) AddLifecycle(ctx context.Context, name string, lifecycle CloneSetLifecycle) {
	sync := NewCloneSetLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *cloneSetClient) AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle CloneSetLifecycle) {
	sync := NewCloneSetLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *cloneSetClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync CloneSetHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *cloneSetClient) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync CloneSetHandlerFunc) {
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}

func (s *cloneSetClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle CloneSetLifecycle) {
	sync := NewCloneSetLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *cloneSetClient) AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle CloneSetLifecycle) {
	sync := NewCloneSetLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}
