resource "null_resource" "name" {}

terraform {
  required_providers {
    null = {
      source = "terraform/abf/null"
    }
  }
}
