resource "aws_s3_bucket" "english-proverbs" {
    bucket = "${var.bucket_name}" 
    acl = "${var.acl_value}"    
}