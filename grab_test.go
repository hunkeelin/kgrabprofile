package kgrabprofile

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"testing"
)

func TestGet(t *testing.T) {
	fmt.Println("testing")
	cred, err := GrabProfile("varodev")
	if err != nil {
		panic(err)
	}
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: cred,
	})
	if err != nil {
		panic(err)
	}
	svc := ec2.New(sess)
	results, err := svc.DescribeInstances(nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(results)
}
