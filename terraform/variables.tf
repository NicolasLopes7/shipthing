variable "region" {
  default = "us-east-1"
}

variable "tfstate_bucket_name" {
 
}

variable "tfstate_bucket_key_name" {
 
}

provider "aws" {
  region     = var.region
}


