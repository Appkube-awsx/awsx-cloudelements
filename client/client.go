package client

import (
	"github.com/Appkube-awsx/awsx-cloudelements/assumerole"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/configservice"
	"log"
)

func GetClient(region string, crossAccountRoleArn string, accessKey string, secretKey string, sessionName string, externalId string) (*configservice.ConfigService, error) {
	assumeRoleOutput, err := assumerole.GetAssumeRole(region, accessKey, secretKey, crossAccountRoleArn, sessionName, externalId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	awsConfigSession, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(*assumeRoleOutput.Credentials.AccessKeyId, *assumeRoleOutput.Credentials.SecretAccessKey, *assumeRoleOutput.Credentials.SessionToken),
		Region:      aws.String(region),
	})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	configClient := configservice.New(awsConfigSession)
	return configClient, nil
}
