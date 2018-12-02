package apiclient

type ModelResult struct {
	ModelId	string	`json:"model_id"`
}

type ModelVersionResult struct {
	VersionId string `json:"version_id"`
	UploadUrl string `json:"upload_url"`
}

type DeploymentResult struct {
	ModelId string `json:"model_id"`
	DeploymentId string `json:"deployment_id"`
}

type DeploymentServiceResult struct {
	ServiceId string `json:"service_id"`
}

type DeploymentEndpointResult struct {
	EndpointId string `json:"endpoint_id"`
	ServiceId string `json:"ServiceId"`
}

