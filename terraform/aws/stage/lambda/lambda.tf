
locals {
  app_id = "${lower(var.app_name)}-${lower(var.app_env)}-${random_id.unique_suffix.hex}"
}

data "archive_file" "lambda_zip" {
  type        = "zip"
  source_file = "../../bin/app"
  output_path = "../../bin/app.zip"
}

resource "random_id" "unique_suffix" {
  byte_length = 2
}

resource "aws_lambda_function" "twitter_bot_func" {
  filename         = data.archive_file.lambda_zip.output_path
  function_name    = local.app_id
  handler          = "app"
  source_code_hash = base64sha256(data.archive_file.lambda_zip.output_path)
  runtime          = "go1.x"
  role             = aws_iam_role.lambda_exec.arn

  environment {
    variables = {
      BUCKET_NAME = "english-proverbs-cirta",
      FILE_NAME = "proverbs.txt",
      TWITTER_CONSUMER_KEY = "o37I7QrJe6Jf7DqRuT7hHeuCo",
      TWITTER_CONSUMER_SECRET = "1HglkuCH30gceeSPGiFeSS4j5m3WGkUEeE92P09eN2tnUAJxKU",
      TWITTER_ACCESS_TOKEN = "1508039357898317828-CI5Mj4vMp9i21qy5FsXQCxrU5fmBjP",
      TWITTER_ACCESS_SECRET = "mgoyZI5MLoijqGNyBEgOpNkj9zekD9EU2SO1jJZs0rh6R"
    }
  }
}


# Assume role setup
resource "aws_iam_role" "lambda_exec" {
  name_prefix = local.app_id

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF

}

# Attach role to Managed Policy
variable "iam_policy_arn" {
  description = "IAM Policy to be attached to role"
  type        = list(string)

  default = [
    "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
  ]
}

resource "aws_iam_policy_attachment" "role_attach" {

  name       = "policy-${local.app_id}"
  roles      = [aws_iam_role.lambda_exec.id]
  count      = length(var.iam_policy_arn)
  policy_arn = element(var.iam_policy_arn, count.index)
}

resource "aws_iam_policy" "lambda_policy" {
  name        = "${var.app_env}_lambda_policy"
  description = "${var.app_env}_lambda_policy"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "s3:ListBucket",
        "s3:GetObject",
        "s3:CopyObject",
        "s3:HeadObject"
      ],
      "Effect": "Allow",
      "Resource": [
        "${var.s3_arn}/*"
      ]
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "terraform_lambda_iam_policy_basic_execution" {
 role = "${aws_iam_role.lambda_exec.id}"
 policy_arn = "${aws_iam_policy.lambda_policy.arn}"
}

resource "aws_lambda_permission" "allow_terraform_bucket" {
   statement_id = "AllowExecutionFromS3Bucket"
   action = "lambda:InvokeFunction"
   function_name = "${aws_lambda_function.twitter_bot_func.arn}"
   principal = "s3.amazonaws.com"
   source_arn = "${var.s3_arn}"
}