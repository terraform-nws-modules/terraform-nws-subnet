terraform {
  required_version = ">= 1.0.10"

  required_providers {
    nws = {
      source  = "nws/nws"
      version = "0.1.2"
    }
  }
}

resource "nws_network" "net" {
  count            = length(var.name)
  zone             = var.zone
  name             = var.name[count.index]
  cidr             = var.cidr[count.index]
  network_domain   = var.domain
  vpc_id           = var.vpc_id
  network_offering = var.public ? "DefaultSharedNetworkOffering" : "DefaultIsolatedNetworkOfferingForVpcNetworks"
}
