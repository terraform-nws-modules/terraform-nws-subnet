output "id" {
  description = "UUID of the new subnet"
  value       = module.subnet.id
}

output "domain" {
  description = "Network domain of the new subnet"
  value       = module.subnet.domain
}
