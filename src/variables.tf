# ---------------------------------------------------------------------------------------------------------------------
# REQUIRED PARAMETERS
# You must provide a value for each of these parameters.
# ---------------------------------------------------------------------------------------------------------------------
variable "name" {
  description = "Your network name"
  type        = list(string)
}
variable "cidr" {
  description = "Your network CIDR"
  type        = list(string)
}
# ---------------------------------------------------------------------------------------------------------------------
# OPTIONAL PARAMETERS
# These parameters have reasonable defaults.
# ---------------------------------------------------------------------------------------------------------------------
variable "zone" {
  description = "Your zone name"
  type        = string
  default     = "ru-msk-0"
}

variable "public" {
  description = "Should this subnet be publicly available ?"
  type        = bool
  default     = false
}

variable "vpc_id" {
  description = "Your VPC UUID. Null for standalone networks without a VPC"
  type        = string
  default     = null
}
variable "domain" {
  type    = string
  default = "my.local"
}
