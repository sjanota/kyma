package apigateway

import (
	"context"
	"github.com/kyma-incubator/api-gateway/api/v1alpha1"

	"github.com/kyma-project/kyma/components/console-backend-service/internal/resource"

	"github.com/kyma-project/kyma/components/console-backend-service/internal/domain/apigateway/disabled"
	"github.com/kyma-project/kyma/components/console-backend-service/internal/gqlschema"
	"github.com/kyma-project/kyma/components/console-backend-service/internal/module"
)

type PluggableResolver struct {
	*module.Pluggable
	Resolver
	serviceFactory *resource.ServiceFactory
}

func New(serviceFactory *resource.ServiceFactory) (*PluggableResolver, error) {
	resolver := &PluggableResolver{
		Pluggable:      module.NewPluggable("apigateway"),
		serviceFactory: serviceFactory,
	}

	err := resolver.Disable()
	if err != nil {
		return nil, err
	}

	return resolver, nil
}

func (r *PluggableResolver) Enable() error {
	apiRuleService := NewService(r.serviceFactory)
	apiRuleResolver, err := newApiRuleResolver(apiRuleService)
	if err != nil {
		return err
	}

	r.Pluggable.EnableAndSyncDynamicInformerFactory(r.serviceFactory.InformerFactory, func() {
		r.Resolver = &domainResolver{
			apiRuleResolver: apiRuleResolver,
		}
	})
	return nil
}
func (r *PluggableResolver) Disable() error {
	r.Pluggable.Disable(func(disabledErr error) {
		r.Resolver = disabled.NewResolver(disabledErr)
	})
	return nil
}

//go:generate failery -name=Resolver -case=underscore -output disabled -outpkg disabled
type Resolver interface {
	APIRulesQuery(ctx context.Context, namespace string, serviceName *string, hostname *string) ([]*v1alpha1.APIRule, error)
	APIRuleQuery(ctx context.Context, name string, namespace string) (*v1alpha1.APIRule, error)
	CreateAPIRule(ctx context.Context, name string, namespace string, params gqlschema.APIRuleInput) (*v1alpha1.APIRule, error)
	UpdateAPIRule(ctx context.Context, name string, namespace string, params gqlschema.APIRuleInput) (*v1alpha1.APIRule, error)
	DeleteAPIRule(ctx context.Context, name string, namespace string) (*v1alpha1.APIRule, error)
	APIRuleEventSubscription(ctx context.Context, namespace string, serviceName *string) (<-chan *gqlschema.APIRuleEvent, error)
}

type domainResolver struct {
	*apiRuleResolver
}
