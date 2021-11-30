variable "name" {
  description = "Your subnet name"
  type        = string
}

variable "vpc_id" {
  description = "Your VPC UUID, for which you create a subnet"
  type        = string
}

variable "cidr" {
  description = "Your network full CIDR"
  type        = string
}

variable "domain" {
  description = "Your network domain"
  type        = string
}
