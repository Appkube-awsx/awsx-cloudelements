package appconfigcmd

import (
	"github.com/Appkube-awsx/awsx-cloudelements/client"
	"github.com/aws/aws-sdk-go/service/configservice"
	"log"
)

func GetDiscoveredResourceCounts(region string, crossAccountRoleArn string, accessKey string, secretKey string, externalId string) (*configservice.GetDiscoveredResourceCountsOutput, error) {
	log.Println("Getting aws config resource count summary")
	configServiceClient, err := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	configResourceRequest := &configservice.GetDiscoveredResourceCountsInput{}
	configResourceResponse, err := configServiceClient.GetDiscoveredResourceCounts(configResourceRequest)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	log.Println(configResourceResponse)
	return configResourceResponse, nil
}
