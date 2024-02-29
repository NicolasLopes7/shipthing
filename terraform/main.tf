resource "aws_s3_bucket" "shipthing" {
  bucket = "shipthing"
}

resource "aws_s3_bucket_logging" "shipthing" {
    bucket = aws_s3_bucket.shipthing.id

    target_bucket = "shipthing-access-logs"
    target_prefix = aws_s3_bucket.shipthing.id
}