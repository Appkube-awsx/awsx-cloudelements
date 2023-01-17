package client

import (
	"github.com/Appkube-awsx/awsx-cloudelements/awssession"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/service/configservice"
	"time"
)

func GetClient(region string, crossAccountRoleArn string, accessKey string, secretKey string) *configservice.ConfigService {
	awsSession := awssession.GetSessionByCreds(region, accessKey, secretKey)
	creds := stscreds.NewCredentials(awsSession, crossAccountRoleArn, func(arp *stscreds.AssumeRoleProvider) {
		//arp.RoleSessionName = "my session role name"
		arp.Duration = 60 * time.Minute
		arp.ExpiryWindow = 30 * time.Second
	})

	configClient := configservice.New(awsSession, aws.NewConfig().WithCredentials(creds))
	return configClient
}
