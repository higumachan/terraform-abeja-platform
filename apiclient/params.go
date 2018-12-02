package apiclient

type CreateModelParam struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

type CreateModelVersionParam struct {
	Version string `json:"version"`
	Handler string `json:"handler"`
	Image string `json:"image"`
	ContentType string `json:"content_type"`
}

type CreateDeploymentParam struct {
	Name string `json:"name"`
	DefaultEnvironment string `json:"default_environment"`
}