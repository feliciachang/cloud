data "aws_iam_policy_document" "fieldkit-server-staging" {
  statement {
    actions = [
      "sqs:SendMessage",
      "sqs:DeleteMessage",
      "sqs:ReceiveMessage",
    ]

    resources = [
      "${aws_sqs_queue.fieldkit-staging.arn}",
    ]
  }

  statement {
    actions = [
      "ses:SendEmail",
    ]

    resources = [
      "*",
    ]

    condition {
      test     = "StringEquals"
      variable = "ses:FromAddress"

      values = [
        "admin@fieldkit.team",
      ]
    }
  }
}

resource "aws_iam_role" "fieldkit-server-staging" {
  name = "fieldkit-server-staging"

  assume_role_policy = <<EOF
{
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "ec2.amazonaws.com"
      },
      "Effect": "Allow"
    }
  ]
}
EOF
}

resource "aws_iam_instance_profile" "fieldkit-server-staging" {
  name  = "fieldkit-server-staging"
  roles = ["${aws_iam_role.fieldkit-server-staging.name}"]
}

resource "aws_iam_role_policy" "fieldkit-server-staging" {
  name   = "fieldkit-server-staging"
  role   = "${aws_iam_role.fieldkit-server-staging.id}"
  policy = "${data.aws_iam_policy_document.fieldkit-server-staging.json}"
}