variable "user_id" {}
variable "personal_access_token" {}
variable "organization_id" {}

provider "abeja" {
  user_id = "${var.user_id}"
  personal_access_token = "${var.personal_access_token}"
  organization_id = "${var.organization_id}"
}

resource "abeja_model" "my-model" {
  name = "nadeko"
  description = "nadeko is cute"
}

resource "abeja_model_version" "my-modelversion" {
  model_id = "${abeja_model.my-model.id}"
  version = "0.0.1"
  image = "abeja-inc/all-cpu:18.10"
  handler = "main:handler"
  //upload_file_path = "run.sh"
  upload_file_path = "upload/model.tar.gz"
}

resource "abeja_deployment" "my-deployment" {
  model_id = "${abeja_model.my-model.id}"
  name = "rikka"
  default_environment = "{}"
}

resource "abeja_deployment_service" "my-service" {
  deployment_id = "${abeja_deployment.my-deployment.id}"
  instance_number = 1
  instance_type = "cpu-0.25"
  environment = "{}"
  version_id = "${abeja_model_version.my-modelversion.id}"
}
