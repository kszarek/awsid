package main

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
)

type mockIAMClient struct {
	iamiface.IAMAPI
}

var dummyUser = &iam.User{
	Arn:              aws.String("asdj910jd823hf298hf923ghf923h"),
	CreateDate:       aws.Time(time.Now()),
	PasswordLastUsed: aws.Time(time.Now()),
	Path:             aws.String("/"),
	UserId:           aws.String("AKIAIWGSMDPM4875FVZA"),
	UserName:         aws.String("mockUser"),
}
var dummyGetUserOutput = &iam.GetUserOutput{
	User: dummyUser,
}

func (m *mockIAMClient) GetUser(input *iam.GetUserInput) (*iam.GetUserOutput, error) {
	return dummyGetUserOutput, nil
}

func TestGetUserName(t *testing.T) {
	var dummyClient = &Client{iam: &mockIAMClient{}}
	user, err := dummyClient.getUserName()
	if err != nil {
		t.Fatal(err)
	}
	if user != "mockUser" {
		t.Fatalf("Wrong user")
	}
}
