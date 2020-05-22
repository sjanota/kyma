package resource

import (
	"fmt"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"time"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
)

var NotFound = fmt.Errorf("resource not found")

type ServiceFactory struct {
	Client          dynamic.Interface
	InformerFactory dynamicinformer.DynamicSharedInformerFactory
}

func NewServiceFactoryForConfig(config *rest.Config, informerResyncPeriod time.Duration) (*ServiceFactory, error) {
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	informerFactory := dynamicinformer.NewDynamicSharedInformerFactory(dynamicClient, informerResyncPeriod)
	return NewServiceFactory(dynamicClient, informerFactory), nil
}

func NewServiceFactory(client dynamic.Interface, informerFactory dynamicinformer.DynamicSharedInformerFactory) *ServiceFactory {
	return &ServiceFactory{
		Client:          client,
		InformerFactory: informerFactory,
	}
}

type Service struct {
	client   dynamic.NamespaceableResourceInterface
	informer cache.SharedIndexInformer
}

func (f *ServiceFactory) ForResource(gvr schema.GroupVersionResource) *Service {
	return &Service{
		client:   f.Client.Resource(gvr),
		informer: f.InformerFactory.ForResource(gvr).Informer(),
	}
}

func (s *Service) ListByIndex(index, key string, result Appendable) error {
	items, err := s.informer.GetIndexer().ByIndex(index, key)
	if err != nil {
		return err
	}

	for _, item := range items {
		converted := result.Append()
		err := FromUnstructured(item.(*unstructured.Unstructured), converted)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) ListInNamespace(namespace string, result Appendable) error {
	return s.ListByIndex(cache.NamespaceIndex, namespace, result)
}

func (s *Service) List(result Appendable) error {
	return s.ListInNamespace("", result)
}

func (s *Service) GetByKey(key string, result interface{}) error {
	item, exists, err := s.informer.GetStore().GetByKey(key)
	if err != nil {
		return err
	}
	if !exists {
		return NotFound
	}

	err = FromUnstructured(item.(*unstructured.Unstructured), result)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetInNamespace(name, namespace string, result interface{}) error {
	key := fmt.Sprintf("%s/%s", namespace, name)
	return s.GetByKey(key, result)
}

func (s *Service) Get(name string, result interface{}) error {
	return s.GetByKey(name, result)
}