package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
)

// Client for AWS services
type Client struct {
	iam iamiface.IAMAPI
}

// getUserName returns user name based on the AWS access key ID
func (c *Client) getUserName() (string, error) {
	resp, err := c.iam.GetUser(&iam.GetUserInput{})
	if err != nil {
		return "", err
	}
	return *resp.User.UserName, nil
}

func main() {
	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}
	svc := iam.New(sess)
	aws := &Client{iam: svc}
	fmt.Println(aws.getUserName())
}
