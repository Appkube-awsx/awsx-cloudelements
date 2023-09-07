package client

import (
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/aws/aws-sdk-go/service/configservice"
)

func GetClient(region string, crossAccountRoleArn string, accessKey string, secretKey string, externalId string) (*configservice.ConfigService, error) {
	auth := client.Auth{
		Region:              region,
		CrossAccountRoleArn: crossAccountRoleArn,
		AccessKey:           accessKey,
		SecretKey:           secretKey,
		ExternalId:          externalId,
	}
	awsConfigSession := client.GetSessionWithAssumeRole(auth)
	configClient := configservice.New(awsConfigSession)
	return configClient, nil
}
