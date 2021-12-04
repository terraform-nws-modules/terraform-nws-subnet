output "subnet_id" {
  description = "UUID for this subnet"
  value       = nws_network.net[*].id
}

output "subnet_public" {
  description = "Network domain for this subnet"
  value       = var.public
}

output "subnet_acl_id" {
  description = "ACL id for this subnet"
  value       = nws_network_acl.acl[*].id
}

output "subnet_acl_rule_id" {
  description = "ACL rule id for this subnet"
  value       = nws_network_acl_rule.rule[*].id
}
