package apiclient

type ModelResult struct {
	ModelId	string	`json:"model_id"`
}

type ModelVersionResult struct {
	VersionId string `json:"version_id"`
	UploadUrl string `json:"upload_url"`
}

type DeploymentResult struct {
	DeploymentId string `json:"deployment_id"`
}