variable "name" {
  description = "Your subnet name"
  type        = list(string)
}

variable "cidr" {
  description = "Your network full CIDR"
  type        = list(string)
}

variable "domain" {
  description = "Your network domain"
  type        = string
}

variable "public" {
  description = "Should this subnet be publicly available ?"
  type        = bool
}
