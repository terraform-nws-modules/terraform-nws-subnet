terraform {
  required_version = ">= 1.0.10"

  required_providers {
    nws = {
      source  = "nws/nws"
      version = "0.1.2"
    }
  }
}

module "subnet" {
  source = "../../src"

  name   = var.name
  cidr   = var.cidr
  domain = var.domain
  public = var.public
}
