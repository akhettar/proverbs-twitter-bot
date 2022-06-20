
variable "app_name" {
  description = "Application name"
  default     = "twitter-bot-lambda"
}

variable "app_env" {
  description = "Application environment tag"
  default     = "dev"
}

variable "aws_profile" {
  description = "Given name in the credential file"
  type        = string
  default     = "rahul-admin"
}

variable "shared_credentials_file" {
  description = "Profile file with credentials to the AWS account"
  type        = string
  default     = "~/.aws/credentials"
}

variable "tags" {
  description = "A map of tags to add to all resources."
  type        = map(string)
  default = {
    application = "Learning-Tutor"
    env         = "Test"
  }
}