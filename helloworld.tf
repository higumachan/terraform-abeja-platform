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
