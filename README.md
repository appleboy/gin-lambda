# gin-lambda

running golang using gin framework in AWS Lambda &amp; API Gateway

## AWS Policy

Add the following AWS policy if you want to integrate with CI/CD tools like Jenkins, GitLab Ci or Drone.

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "s3:PutObject",
        "iam:ListRoles",
        "lambda:UpdateFunctionCode",
        "lambda:CreateFunction"
      ],
      "Resource": "arn:aws:logs:*:*:*"
    }
  ]
}
```
