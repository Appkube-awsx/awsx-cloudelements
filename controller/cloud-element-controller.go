package controller

import (
	"github.com/Appkube-awsx/awsx-cloudelements/command/appconfigcmd"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/aws/aws-sdk-go/service/configservice"
	"log"
)

func GetDiscoveredResourceByAccountNo(vaultUrl string, vaultToken string, accountNo string, region string) (*configservice.GetDiscoveredResourceCountsOutput, error) {
	authFlag, clientAuth, err := authenticate.AuthenticateData(vaultUrl, vaultToken, accountNo, region, "", "", "", "")
	return GetDiscoveredResource(authFlag, clientAuth, err)
}

func GetDiscoveredResourceByUserCreds(region string, accesskey string, secretKey string, crossAccountRoleArn string, externalId string) (*configservice.GetDiscoveredResourceCountsOutput, error) {
	authFlag, clientAuth, err := authenticate.AuthenticateData("", "", "", region, accesskey, secretKey, crossAccountRoleArn, externalId)
	return GetDiscoveredResource(authFlag, clientAuth, err)
}

func GetDiscoveredResource(authFlag bool, clientAuth *client.Auth, err error) (*configservice.GetDiscoveredResourceCountsOutput, error) {
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if !authFlag {
		log.Println(err.Error())
		return nil, err
	}
	response, err := appconfigcmd.GetDiscoveredResourceCounts(*clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return response, nil
}
