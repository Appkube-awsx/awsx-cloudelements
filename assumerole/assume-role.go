package assumerole

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-cloudelements/awssession"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sts"
)

func GetAssumeRole(region string, accessKey string, secretKey string, roleArn string, sessionName string, externalId string) (*sts.AssumeRoleOutput, error) {
	sess, err := awssession.GetSessionByCreds(region, accessKey, secretKey)
	if err != nil {
		fmt.Printf("Failed to create aws session, %v\n", err)
		return nil, err
	}

	securityTokenServiceObj := sts.New(sess)
	assumeRoleInput := sts.AssumeRoleInput{
		RoleArn:         aws.String(roleArn),
		RoleSessionName: aws.String(sessionName),
		DurationSeconds: aws.Int64(60 * 60 * 1),
		ExternalId:      aws.String(externalId),
	}

	assumeRoleOutput, err := securityTokenServiceObj.AssumeRole(&assumeRoleInput)
	if err != nil {
		fmt.Printf("Failed to create assume role, %v\n", err)
		return nil, err
	}
	return assumeRoleOutput, nil
}
