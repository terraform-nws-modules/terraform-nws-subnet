variable "name" {
  description = "Your subnet name"
  type        = list(string)
}

variable "cidr" {
  description = "Your network full CIDR"
  type        = list(string)
}
variable "vpc_id" {
  description = "Your VPC UUID, for which you create a subnet"
  type        = string
}

variable "domain" {
  description = "Your network domain"
  type        = string
}

variable "public" {
  description = "Should this subnet be publicly available ?"
  type        = bool
}

variable "has_vpc" {
  description = "Should this subnet belong to a VPC?"
  type        = bool
}

variable "acl_allowed_port_list" {
  description = "Port list for allowed ingress traffic"
  type        = list(string)
}
