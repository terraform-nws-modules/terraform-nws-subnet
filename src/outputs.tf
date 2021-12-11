output "subnet_id" {
  description = "UUID list for these subnets"
  value       = nws_network.net[*].id
}

output "subnet_public" {
  description = "Should this subnet be public ?"
  value       = var.public
}

output "subnet_has_vpc" {
  description = "Does this subnect belong to a VPC ?"
  value       = var.has_vpc
}

output "subnet_acl_id" {
  description = "ACL id for this subnet"
  value       = nws_network_acl.acl[*].id
}

output "subnet_acl_rule_id" {
  description = "ACL rule id for this subnet"
  value       = nws_network_acl_rule.rule[*].id
}
