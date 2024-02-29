terraform {
    backend "s3" {
        bucket = "shipthing-tfstates"
        key    = "terraform.tfstate"
        region = "us-east-1"
    }
}