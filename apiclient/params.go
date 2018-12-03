package apiclient

type CreateModelParam struct {
	Name string `json:"name" tf-schema:"name"`
	Description string `json:"description" tf-schema:"description"`
}

type CreateModelVersionParam struct {
	Version string `json:"version" tf-schema:"version"`
	Handler string `json:"handler" tf-schema:"handler"`
	Image string `json:"image" tf-schema:"image"`
	ContentType string `json:"content_type" tf-schema:"content_type"`
}

type CreateDeploymentParam struct {
	Name string `json:"name" tf-schema:"name"`
	DefaultEnvironment string `json:"default_environment" tf-schema:"default_environment"`
}

type CreateDeploymentServiceParam struct {
	InstanceNumber int `json:"instance_number" tf-schema:"instance_number"`
	InstanceType string `json:"instance_type" tf-schema:"instance_type"`
	Environment string `json:"environment" tf-schema:"environment"`
	VersionId string `json:"version_id" tf-schema:"version_id"`
}

type CreateDeploymentEndpointParam struct {
	ServiceId string `json:"service_id" tf-schema:"service_id"`
	CustomAlias string `json:"custom_alias" tf-schema:"custom_alias"`
}