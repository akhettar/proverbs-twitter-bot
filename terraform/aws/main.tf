provider "aws" {
    profile =                 "${var.profile}"
    shared_credentials_file = "~/.aws/credentials"
}


resource "aws_s3_bucket" "terraform_state" {
  bucket = "terraform-state-cirta"
  # Enable versioning so we can see the full revision history of our
  # state files
  versioning {
    enabled = true
  }
  # Enable server-side encryption by default
  server_side_encryption_configuration {
    rule {
      apply_server_side_encryption_by_default {
        sse_algorithm = "AES256"
      }
    }
  }
}

module "s3" {
    source = "./s3"
    #bucket name should be unique
    bucket_name = "english-proverbs-cirta"       
}

module "twitter-bot" {
  source = "./lambda"
}