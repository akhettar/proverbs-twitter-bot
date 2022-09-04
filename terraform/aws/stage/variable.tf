
variable "profile" {
  default = "default"
}

variable "region" {
  default = "eu-central-1"
}

variable "allowed_account_id" {
  description = "The id of the one AWS account this code is permitted to run against"
  type        = string
  default     = "762587740553"
}