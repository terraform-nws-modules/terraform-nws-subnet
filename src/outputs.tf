output "id" {
  description = "UUID of the new subnet"
  value       = nws_network.net.id
}

output "domain" {
  description = "Network domain of the new subnet"
  value       = nws_network.net.network_domain
}
