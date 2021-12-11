terraform {
  required_version = ">= 1.0.10"

  required_providers {
    nws = {
      source  = "nws/nws"
      version = "0.1.3"
    }
  }
}

module "subnet" {
  source = "../../src"

  name    = var.name
  cidr    = var.cidr
  vpc_id  = var.vpc_id
  domain  = var.domain
  public  = var.public
  has_vpc = var.has_vpc

  acl_allowed_port_list = var.acl_allowed_port_list
}
