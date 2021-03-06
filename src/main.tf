terraform {
  required_version = ">= 1.0.10"

  required_providers {
    nws = {
      source  = "nws/nws"
      version = "0.1.3"
    }
  }
}

resource "nws_network" "net" {
  count            = var.name != null ? length(var.name) : 0
  zone             = var.zone
  network_domain   = var.domain
  name             = var.name[count.index]
  cidr             = var.cidr[count.index]
  vpc_id           = var.has_vpc && var.vpc_id != null ? var.vpc_id : null
  network_offering = var.has_vpc ? "DefaultIsolatedNetworkOfferingForVpcNetworks" : "DefaultIsolatedNetworkOfferingWithSourceNatService"
}

// Create ACL resources for Public networks only
resource "nws_network_acl" "acl" {
  count  = var.has_vpc && var.public && var.name != null ? 1 : 0
  name   = var.acl_name
  vpc_id = var.vpc_id
}

resource "nws_network_acl_rule" "rule" {
  count  = var.has_vpc && var.public && var.name != null ? 1 : 0
  acl_id = nws_network_acl.acl[0].id

  rule {
    action       = "allow"
    protocol     = "tcp"
    cidr_list    = var.acl_allowed_cidr_list
    ports        = var.acl_allowed_port_list
    traffic_type = "ingress"
  }
}
