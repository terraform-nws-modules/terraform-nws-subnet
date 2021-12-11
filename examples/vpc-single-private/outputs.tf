output "subnet_id" {
  description = "UUID of the new subnet"
  value       = module.subnet.subnet_id
}

output "subnet_public" {
  description = "Is this subnet public ?"
  value       = var.public
}

output "subnet_has_vpc" {
  description = "Does this subnect belong to a VPC ?"
  value       = var.has_vpc
}

output "subnet_acl_id" {
  description = "ACL id for this subnet"
  value       = module.subnet.subnet_acl_id
}

output "subnet_acl_rule_id" {
  description = "ACL rule id for this subnet"
  value       = module.subnet.subnet_acl_rule_id
}
