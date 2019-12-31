package kgrabprofile

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"io/ioutil"
	"os/user"
)

func GrabProfile(profile string) (*credentials.Credentials, error) {
	var emptyReturn *credentials.Credentials
	user, err := user.Current()
	if err != nil {
		return emptyReturn, err
	}
	awsDir := user.HomeDir + "/.aws/credentials"
	credentialFile, err := ioutil.ReadFile(awsDir)
	if err != nil {
		return emptyReturn, err
	}
	splitByLines := bytes.Split(credentialFile, []byte("\n"))
	credMap := make(map[string]awsCredentials)
	var cp string
	for _, line := range splitByLines {
		switch {
		case bytes.HasPrefix(line, []byte("[")):
			cp = string(bytes.Trim(bytes.Trim(line, "["), "]"))
		case bytes.HasPrefix(line, []byte("aws_access_key_id")):
			tmp := credMap[cp]
			tmp.awsAccessKeyId = string(bytes.TrimLeft(line, "aws_access_key_id = "))
			credMap[cp] = tmp
		case bytes.HasPrefix(line, []byte("aws_secret_access_key")):
			tmp := credMap[cp]
			tmp.awsSecretAccessKey = string(bytes.TrimLeft(line, "aws_secret_access_key = "))
			credMap[cp] = tmp
		case bytes.HasPrefix(line, []byte("aws_session_token")):
			tmp := credMap[cp]
			tmp.awsSessionToken = string(bytes.TrimLeft(line, "aws_session_token = "))
			credMap[cp] = tmp
		default:
			continue
		}
	}
	awsProfile, ok := credMap[profile]
	if !ok {
		return emptyReturn, fmt.Errorf("selected profile doesn't exist")
	}
	return credentials.NewStaticCredentials(awsProfile.awsAccessKeyId, awsProfile.awsSecretAccessKey, awsProfile.awsSessionToken), nil
}

type awsCredentials struct {
	awsAccessKeyId     string
	awsSecretAccessKey string
	awsSessionToken    string
}
