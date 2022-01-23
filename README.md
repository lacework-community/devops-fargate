Lacework Agent + AWS ECS Fargate Demo
=========================

- App: nginx
- CI/CD: CircleCI (with Orbs)
- IaC: Terraform AWS Provider (with CloudFormation)
- Registry: AWS ECR
- Env: AWS ECS using AWS Fargate

### REQUIRED: Configure [environment variables](https://circleci.com/docs/2.0/env-vars/#setting-an-environment-variable-in-a-project) on CircleCI

- For the AWS ECR Orb
- For the Terraform Orb
  - AWS Provider
- Add the LW_AGENT_ACCESS_TOKEN
