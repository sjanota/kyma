package main

import (
	"github.com/avast/retry-go"
	"github.com/kyma-project/kyma/components/console-backend-service/integration/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vrischmann/envconfig"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"net"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sync"
	"testing"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
)

func TestIntegration(t *testing.T) {
	// setup and start kube-apiserver
	environment := &envtest.Environment{}
	restConfig, err := environment.Start()
	require.NoError(t, err)

	_, err = envtest.InstallCRDs(restConfig, envtest.CRDInstallOptions{
		Paths:              []string{"../../resources/cluster-essentials/templates/crds"},
		ErrorIfPathMissing: true,
	})
	require.NoError(t, err)

	appConfig := config{}
	err = envconfig.InitWithOptions(&appConfig, envconfig.Options{Prefix: "CBS_TEST", AllOptional: true})
	require.NoError(t, err)
	appConfig.OIDC.IssuerURL = "https://dex.kyma.local"
	appConfig.AuthEnabled = false

	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)
	port := listener.Addr().(*net.TCPAddr).Port

	wg := &sync.WaitGroup{}
	stopCh := make(chan struct{})

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = run(listener, stopCh, appConfig, restConfig)
		require.NoError(t, err)
	}()

	test(t, restConfig, port)

	close(stopCh)
	wg.Wait()
}

func test(t *testing.T, restConfig *rest.Config, port int) {
	suite := givenNewTestNamespaceSuite(t, restConfig, port)

	t.Log("Creating namespace...")
	createRsp, err := suite.whenNamespaceIsCreated()
	suite.thenThereIsNoError(t, err)
	suite.thenThereIsNoGqlError(t, createRsp.GqlErrors)
	suite.thenCreateNamespaceResponseIsAsExpected(t, createRsp)
	suite.thenNamespaceExistsInK8s(t)

	t.Log("Quering for namespace...")
	queryRsp, err := suite.whenNamespaceIsQueried()
	suite.thenThereIsNoError(t, err)
	suite.thenThereIsNoGqlError(t, queryRsp.GqlErrors)
	suite.thenNamespaceResponseIsAsExpected(t, queryRsp)

	t.Log("Updating namespace...")
	updateResp, err := suite.whenNamespaceIsUpdated()
	suite.thenThereIsNoError(t, err)
	suite.thenThereIsNoGqlError(t, updateResp.GqlErrors)
	suite.thenUpdateNamespaceResponseIsAsExpected(t, updateResp)
	suite.thenNamespaceIsUpdatedInK8s(t)

	t.Log("Deleting namespace...")
	deleteRsp, err := suite.whenNamespaceIsDeleted()
	suite.thenThereIsNoError(t, err)
	suite.thenThereIsNoGqlError(t, deleteRsp.GqlErrors)
	suite.thenDeleteNamespaceResponseIsAsExpected(t, deleteRsp)
	suite.thenNamespaceIsTerminating(t)
}

type testNamespaceSuite struct {
	gqlClient     *graphql.Client
	k8sClient     *corev1.CoreV1Client
	namespaceName string
	labels        map[string]string
	updatedLabels map[string]string
}

func givenNewTestNamespaceSuite(t *testing.T, restConfig *rest.Config, port int) testNamespaceSuite {
	c, err := graphql.New(port)
	require.NoError(t, err)

	k8s, err := corev1.NewForConfig(restConfig)
	require.NoError(t, err)

	suite := testNamespaceSuite{
		gqlClient:     c,
		k8sClient:     k8s,
		namespaceName: "test-namespace",
		labels: map[string]string{
			"aaa": "bbb",
		},
		updatedLabels: map[string]string{
			"ccc": "ddd",
		},
	}
	return suite
}

func (s testNamespaceSuite) whenNamespaceIsCreated() (createNamespaceResponse, error) {
	var rsp createNamespaceResponse
	err := s.gqlClient.Do(s.fixNamespaceCreate(), &rsp)
	return rsp, err
}

func (s testNamespaceSuite) thenThereIsNoError(t *testing.T, err error) {
	require.NoError(t, err)
}

func (s testNamespaceSuite) thenThereIsNoGqlError(t *testing.T, gqlErr GqlErrors) {
	require.Empty(t, gqlErr.Errors)
}

func (s testNamespaceSuite) thenCreateNamespaceResponseIsAsExpected(t *testing.T, rsp createNamespaceResponse) {
	assert.Equal(t, s.fixCreateNamespaceResponse(), rsp)
}

func (s testNamespaceSuite) thenNamespaceExistsInK8s(t *testing.T) {
	ns, err := s.k8sClient.Namespaces().Get(s.namespaceName, metav1.GetOptions{})
	require.NoError(t, err)
	assert.Equal(t, ns.Name, s.namespaceName)
	assert.Equal(t, ns.Labels, s.labels)
}

func (s testNamespaceSuite) thenNamespaceIsUpdatedInK8s(t *testing.T) {
	ns, err := s.k8sClient.Namespaces().Get(s.namespaceName, metav1.GetOptions{})
	require.NoError(t, err)
	assert.Equal(t, ns.Name, s.namespaceName)
	assert.Equal(t, ns.Labels, s.updatedLabels)
}

func (s testNamespaceSuite) whenNamespaceIsQueried() (namespaceResponse, error) {
	var rsp namespaceResponse
	err := s.gqlClient.Do(s.fixNamespaceQuery(), &rsp)
	return rsp, err
}

func (s testNamespaceSuite) thenNamespaceResponseIsAsExpected(t *testing.T, rsp namespaceResponse) {
	assert.Equal(t, s.fixNamespaceResponse(), rsp)
}

func (s testNamespaceSuite) whenNamespaceIsUpdated() (updateNamespaceResponse, error) {
	var rsp updateNamespaceResponse
	err := s.gqlClient.Do(s.fixNamespaceUpdate(), &rsp)
	return rsp, err
}

func (s testNamespaceSuite) thenUpdateNamespaceResponseIsAsExpected(t *testing.T, rsp updateNamespaceResponse) {
	assert.Equal(t, s.fixUpdateNamespaceResponse(), rsp)
}

func (s testNamespaceSuite) whenNamespaceIsDeleted() (deleteNamespaceResponse, error) {
	var rsp deleteNamespaceResponse
	err := s.gqlClient.Do(s.fixNamespaceDelete(), &rsp)
	return rsp, err
}

func (s testNamespaceSuite) thenDeleteNamespaceResponseIsAsExpected(t *testing.T, rsp deleteNamespaceResponse) {
	assert.Equal(t, s.fixDeleteNamespaceResponse(), rsp)
}

func (s testNamespaceSuite) thenNamespaceIsTerminating(t *testing.T) {
	err := retry.Do(func() error {
		ns, err := s.k8sClient.Namespaces().Get(s.namespaceName, metav1.GetOptions{})
		if apierrors.IsNotFound(err) || ns.Status.Phase == v1.NamespaceTerminating {
			return nil
		}

		return err
	})
	require.NoError(t, err)
}

func (s testNamespaceSuite) fixNamespaceObj() namespaceObj {
	return namespaceObj{
		Name:   s.namespaceName,
		Labels: s.labels,
	}
}

func (s testNamespaceSuite) fixNamespaceObjAfterUpdate() namespaceObj {
	return namespaceObj{
		Name:   s.namespaceName,
		Labels: s.updatedLabels,
	}
}

func (s testNamespaceSuite) fixCreateNamespaceResponse() createNamespaceResponse {
	return createNamespaceResponse{CreateNamespace: s.fixNamespaceObj()}
}

func (s testNamespaceSuite) fixNamespaceResponse() namespaceResponse {
	return namespaceResponse{Namespace: s.fixNamespaceObj()}
}

func (s testNamespaceSuite) fixUpdateNamespaceResponse() updateNamespaceResponse {
	return updateNamespaceResponse{UpdateNamespace: s.fixNamespaceObjAfterUpdate()}
}

func (s testNamespaceSuite) fixDeleteNamespaceResponse() deleteNamespaceResponse {
	return deleteNamespaceResponse{DeleteNamespace: s.fixNamespaceObjAfterUpdate()}
}

func (s testNamespaceSuite) fixNamespaceCreate() *graphql.Request {
	query := `mutation ($name: String!, $labels: Labels!) {
				  createNamespace(name: $name, labels: $labels) {
					name
					labels
				  }
				}`
	req := graphql.NewRequest(query)
	req.SetVar("name", s.namespaceName)
	req.SetVar("labels", s.labels)
	return req
}

func (s testNamespaceSuite) fixNamespaceQuery() *graphql.Request {
	query := `query ($name: String!) {
				  namespace(name: $name) {
					name
					labels
				  }
				}`
	req := graphql.NewRequest(query)
	req.SetVar("name", s.namespaceName)
	return req
}

func (s testNamespaceSuite) fixNamespaceUpdate() *graphql.Request {
	query := `mutation ($name: String!, $labels: Labels!) {
				  updateNamespace(name: $name, labels: $labels) {
					name
					labels
				  }
				}`
	req := graphql.NewRequest(query)
	req.SetVar("name", s.namespaceName)
	req.SetVar("labels", s.updatedLabels)
	return req
}

func (s testNamespaceSuite) fixNamespaceDelete() *graphql.Request {
	query := `mutation ($name: String!) {
				  deleteNamespace(name: $name) {
					name
					labels
				  }
				}`
	req := graphql.NewRequest(query)
	req.SetVar("name", s.namespaceName)
	return req
}

type namespaceObj struct {
	Name   string `json:"name"`
	Labels labels `json:"labels"`
}

type GqlErrors struct {
	Errors []interface{} `json:"errors"`
}

type createNamespaceResponse struct {
	GqlErrors
	CreateNamespace namespaceObj `json:"createNamespace"`
}

type namespaceResponse struct {
	GqlErrors
	Namespace namespaceObj `json:"namespace"`
}

type updateNamespaceResponse struct {
	GqlErrors
	UpdateNamespace namespaceObj `json:"updateNamespace"`
}

type deleteNamespaceResponse struct {
	GqlErrors
	DeleteNamespace namespaceObj `json:"deleteNamespace"`
}

type labels map[string]string
