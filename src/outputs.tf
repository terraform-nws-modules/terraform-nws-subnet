output "id" {
  description = "UUID of the new subnet"
  value       = nws_network.net[*].id
}

output "public" {
  description = "Network domain of the new subnet"
  value       = var.public
}
