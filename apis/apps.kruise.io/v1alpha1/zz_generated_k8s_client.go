package v1alpha1

import (
	"context"
	"sync"

	"github.com/rancher/norman/controller"
	"github.com/rancher/norman/objectclient"
	"github.com/rancher/norman/objectclient/dynamic"
	"github.com/rancher/norman/restwatch"
	"k8s.io/client-go/rest"
)

type (
	contextKeyType        struct{}
	contextClientsKeyType struct{}
)

type Interface interface {
	RESTClient() rest.Interface
	controller.Starter

	CloneSetsGetter
}

type Client struct {
	sync.Mutex
	restClient rest.Interface
	starters   []controller.Starter

	cloneSetControllers map[string]CloneSetController
}

func NewForConfig(config rest.Config) (Interface, error) {
	if config.NegotiatedSerializer == nil {
		config.NegotiatedSerializer = dynamic.NegotiatedSerializer
	}

	restClient, err := restwatch.UnversionedRESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	return &Client{
		restClient: restClient,

		cloneSetControllers: map[string]CloneSetController{},
	}, nil
}

func (c *Client) RESTClient() rest.Interface {
	return c.restClient
}

func (c *Client) Sync(ctx context.Context) error {
	return controller.Sync(ctx, c.starters...)
}

func (c *Client) Start(ctx context.Context, threadiness int) error {
	return controller.Start(ctx, threadiness, c.starters...)
}

type CloneSetsGetter interface {
	CloneSets(namespace string) CloneSetInterface
}

func (c *Client) CloneSets(namespace string) CloneSetInterface {
	objectClient := objectclient.NewObjectClient(namespace, c.restClient, &CloneSetResource, CloneSetGroupVersionKind, cloneSetFactory{})
	return &cloneSetClient{
		ns:           namespace,
		client:       c,
		objectClient: objectClient,
	}
}
