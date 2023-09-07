package appconfigcmd

import (
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/aws/aws-sdk-go/service/configservice"
	"log"
)

func GetDiscoveredResourceCounts(auth client.Auth) (*configservice.GetDiscoveredResourceCountsOutput, error) {
	log.Println("Getting aws config resource count summary")
	configServiceClient := client.GetClient(auth, client.CONFIG_SERVICE_CLIENT).(*configservice.ConfigService)
	configResourceRequest := &configservice.GetDiscoveredResourceCountsInput{}
	configResourceResponse, err := configServiceClient.GetDiscoveredResourceCounts(configResourceRequest)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	log.Println(configResourceResponse)
	return configResourceResponse, nil
}
