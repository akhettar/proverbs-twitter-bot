provider "aws" {
  profile             = var.profile
  allowed_account_ids = [var.allowed_account_id]
}


resource "aws_s3_bucket" "terraform_state" {
  bucket = "terraform-state-cirta-1"
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

module "s3_bucket" {
  source      = "./s3"
  bucket_name = "eng-proverbs-cirta"
}

module "twitter-bot" {
  source = "./lambda"
  s3_arn = module.s3_bucket.bucket_arn
}