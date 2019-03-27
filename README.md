
`assume-role` is a convenience CLI for assuming temporary credentials for [Amazon Web Services](https://aws.amazon.com/)

# Installation 


# Configuration

`assume-role` expects the AWS config files to exist, for more information please see [AWS CLI Configuration Variables](https://docs.aws.amazon.com/cli/latest/topic/config-vars.html)

## Examples

`~/.aws/credentials`

```
[default]
aws_access_key_id=foo
aws_secret_access_key=bar
```

`~/.aws/config`

```
[default]
region = us-west-2

[profile staging-read]
role_arn = arn:aws:iam::111111111111:role/readOnly
source_profile = default

[profile staging-admin]
role_arn = arn:aws:iam::111111111111:role/fullAdmin
source_profile = default

[profile production-read]
role_arn = arn:aws:iam::999999999999:role/readOnly
source_profile = default
```

# Usage

### Basics

`assume-role` is built using [cobra](https://github.com/spf13/cobra) and has a full built-in help menu accessible via: `assume-role help` or just `assume-role`

### Examples

If you want to export the resulting credentials you can use `eval` like this:

```
$ eval $(assume-role become staging-read)
```

# Contributing 

In order to contribute to `assume-role` you will need the following installed:

[Go](https://github.com/golang) - primary language

*  `brew install go` - `go version` should result in `go1.11.5` or higher