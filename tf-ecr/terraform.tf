provider "aws" {
  access_key = var.aws_access_key
  secret_key = var.aws_secret_key
  region     = var.aws_region
}

terraform {
  backend "s3" {
    bucket = "bb8-devops-fargate"
    key    = "cicd/job1/terraform.tfstate"
    region = "us-east-2"
  }
}

locals {
  aws_ecr_repository_name = var.aws_resource_prefix
}

resource "aws_ecr_repository" "repo" {
  name                 = local.aws_ecr_repository_name
  image_tag_mutability = "MUTABLE"
  image_scanning_configuration {
    scan_on_push = true
  }
}