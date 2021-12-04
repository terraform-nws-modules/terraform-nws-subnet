output "subnet_id" {
  description = "UUID of the new subnet"
  value       = module.subnet.subnet_id
}

output "subnet_public" {
  description = "Network domain of the new subnet"
  value       = var.public
}

output "subnet_acl_id" {
  description = "ACL id for this subnet"
  value       = module.subnet.subnet_acl_id
}

output "subnet_acl_rule_id" {
  description = "ACL rule id for this subnet"
  value       = module.subnet.subnet_acl_rule_id
}
