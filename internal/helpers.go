package internal

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

// SetEvn sets role credentials to the OS environment for subsequently executed commands/subshells
func SetEvn(c *sts.Credentials) {

	fmt.Println(c)
	fmt.Println(*c.AccessKeyId)

	os.Setenv("AWS_ACCESS_KEY_ID", *c.AccessKeyId)
	os.Setenv("AWS_SECRET_ACCESS_KEY", *c.SecretAccessKey)
	os.Setenv("AWS_SESSION_TOKEN", *c.SessionToken)

	fmt.Println(os.Environ())
}

// SourceableBashEnv prints commands to set credentials
// can be used in combination with $eval() to affect parent shell
func SourceableBashEnv(c *sts.Credentials) {
	fmt.Printf("export AWS_ACCESS_KEY_ID=\"%s\"\n", *c.AccessKeyId)
	fmt.Printf("export AWS_SECRET_ACCESS_KEY=\"%s\"\n", *c.SecretAccessKey)
	fmt.Printf("export AWS_SESSION_TOKEN=\"%s\"\n", *c.SessionToken)
}

// SourceableUnsetBashEnv prints commands to unset credentials
// can be used in combination with $eval() to affect parent shell
func SourceableUnsetBashEnv() {
	fmt.Println("unset AWS_ACCESS_KEY_ID")
	fmt.Println("unset AWS_SECRET_ACCESS_KEY")
	fmt.Println("unset AWS_SESSION_TOKEN")
}

// AssumeRole attempts to acquire temporary role credentials using AWS config settings paired with explicit config parameters
// respects default chain of credential providers - i.e. env, shared credentials file (~/.aws/credentials) or EC2 instance role
func AssumeRole(region string, roleArn string, sessionName string) (*sts.Credentials, error) {

	// create session
	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})

	if err != nil {
		return nil, err
	}

	// create service
	svc := sts.New(sess)

	// set input values
	params := &sts.AssumeRoleInput{
		RoleArn:         aws.String(roleArn),
		RoleSessionName: aws.String(sessionName),
	}

	// attempt
	result, err := svc.AssumeRole(params)

	if err != nil {
		return nil, err
	}

	return result.Credentials, nil
}

// AssumeRoleViaProfile attempts to acquire temporary role credentials using AWS config settings paired with a profile name
// respects default chain of credential providers - i.e. env, shared credentials file (~/.aws/credentials) or EC2 instance role
func AssumeRoleViaProfile(profile string) (*sts.Credentials, error) {

	// create session
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile:           profile,
		SharedConfigState: session.SharedConfigEnable,
	})

	if err != nil {
		return nil, err
	}

	result, err := sess.Config.Credentials.Get()

	// force result into (currently) from credentials.Value into standardized upon sts.Credentials
	var c sts.Credentials
	c.AccessKeyId = aws.String(result.AccessKeyID)
	c.SecretAccessKey = aws.String(result.SecretAccessKey)
	c.SessionToken = aws.String(result.SessionToken)

	return &c, nil
}
