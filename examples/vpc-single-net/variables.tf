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
