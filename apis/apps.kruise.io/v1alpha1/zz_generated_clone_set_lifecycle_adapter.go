package v1alpha1

import (
	"github.com/openkruise/kruise-api/apps/v1alpha1"
	"github.com/rancher/norman/lifecycle"
	"github.com/rancher/norman/resource"
	"k8s.io/apimachinery/pkg/runtime"
)

type CloneSetLifecycle interface {
	Create(obj *v1alpha1.CloneSet) (runtime.Object, error)
	Remove(obj *v1alpha1.CloneSet) (runtime.Object, error)
	Updated(obj *v1alpha1.CloneSet) (runtime.Object, error)
}

type cloneSetLifecycleAdapter struct {
	lifecycle CloneSetLifecycle
}

func (w *cloneSetLifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *cloneSetLifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *cloneSetLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*v1alpha1.CloneSet))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *cloneSetLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*v1alpha1.CloneSet))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *cloneSetLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*v1alpha1.CloneSet))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewCloneSetLifecycleAdapter(name string, clusterScoped bool, client CloneSetInterface, l CloneSetLifecycle) CloneSetHandlerFunc {
	if clusterScoped {
		resource.PutClusterScoped(CloneSetGroupVersionResource)
	}
	adapter := &cloneSetLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *v1alpha1.CloneSet) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}
