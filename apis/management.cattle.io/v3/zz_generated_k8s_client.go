package v3

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

	UserAttributesGetter
}

type Client struct {
	sync.Mutex
	restClient rest.Interface
	starters   []controller.Starter

	userAttributeControllers map[string]UserAttributeController
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

		userAttributeControllers: map[string]UserAttributeController{},
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

type UserAttributesGetter interface {
	UserAttributes(namespace string) UserAttributeInterface
}

func (c *Client) UserAttributes(namespace string) UserAttributeInterface {
	objectClient := objectclient.NewObjectClient(namespace, c.restClient, &UserAttributeResource, UserAttributeGroupVersionKind, userAttributeFactory{})
	return &userAttributeClient{
		ns:           namespace,
		client:       c,
		objectClient: objectClient,
	}
}
