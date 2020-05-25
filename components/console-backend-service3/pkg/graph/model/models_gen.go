// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type BackendModule struct {
	Name string `json:"name"`
}

type ClusterMicroFrontend struct {
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	Category        string            `json:"category"`
	ViewBaseURL     string            `json:"viewBaseUrl"`
	Placement       string            `json:"placement"`
	PreloadURL      string            `json:"preloadUrl"`
	NavigationNodes []*NavigationNode `json:"navigationNodes"`
}

type MicroFrontend struct {
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	Category        string            `json:"category"`
	ViewBaseURL     string            `json:"viewBaseUrl"`
	NavigationNodes []*NavigationNode `json:"navigationNodes"`
}

type NavigationNode struct {
	Label               string                 `json:"label"`
	NavigationPath      string                 `json:"navigationPath"`
	ViewURL             string                 `json:"viewUrl"`
	ShowInNavigation    bool                   `json:"showInNavigation"`
	Order               int                    `json:"order"`
	Settings            map[string]interface{} `json:"settings"`
	ExternalLink        string                 `json:"externalLink"`
	RequiredPermissions []*RequiredPermission  `json:"requiredPermissions"`
}

type RequiredPermission struct {
	Verbs    []string `json:"verbs"`
	APIGroup string   `json:"apiGroup"`
	Resource string   `json:"resource"`
}

type ResourceAttributes struct {
	Verb            string  `json:"verb"`
	APIGroup        *string `json:"apiGroup"`
	APIVersion      *string `json:"apiVersion"`
	Resource        *string `json:"resource"`
	ResourceArg     *string `json:"resourceArg"`
	Subresource     string  `json:"subresource"`
	NameArg         *string `json:"nameArg"`
	NamespaceArg    *string `json:"namespaceArg"`
	IsChildResolver bool    `json:"isChildResolver"`
}