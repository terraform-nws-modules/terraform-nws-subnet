output "id" {
  description = "UUID of the new subnet"
  value       = module.subnet.id
}

output "public" {
  description = "Network domain of the new subnet"
  value       = var.public
}
