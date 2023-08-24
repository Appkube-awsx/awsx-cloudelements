/*
Copyright © 2023 Manoj Sharma manoj.sharma@synectiks.com
*/
package cmd

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-cloudelements/client"
	"github.com/Appkube-awsx/awsx-cloudelements/util"
	"github.com/Appkube-awsx/awsx-cloudelements/vault"
	"github.com/aws/aws-sdk-go/service/configservice"
	"github.com/spf13/cobra"
	"log"
)

// AwsxCloudElementsCmd represents the base command when called without any subcommands
var AwsxCloudElementsCmd = &cobra.Command{
	Use:   "getElementDetails",
	Short: "getElementDetails command gets resource counts",
	Long:  `getElementDetails command gets resource counts details of an AWS account`,

	Run: func(cmd *cobra.Command, args []string) {
		vaultUrl, _ := cmd.Flags().GetString("vaultUrl")
		vaultToken, _ := cmd.Flags().GetString("vaultToken")
		accountNo, _ := cmd.Flags().GetString("accountId")
		region, _ := cmd.Flags().GetString("zone")
		acKey, _ := cmd.Flags().GetString("accessKey")
		secKey, _ := cmd.Flags().GetString("secretKey")
		crossAccountRoleArn, _ := cmd.Flags().GetString("crossAccountRoleArn")
		externalId, _ := cmd.Flags().GetString("externalId")
		sessionName := util.RamdomString(5)
		if vaultUrl != "" && accountNo != "" {
			if region == "" {
				log.Println("zone missing")
				cmd.Help()
				return
			}
			if vaultToken == "" {
				log.Println("vault token missing")
				cmd.Help()
				return
			}
			GetCloudConfigSummary(region, vaultUrl, vaultToken, accountNo)
		} else if region != "" && acKey != "" && secKey != "" && crossAccountRoleArn != "" && externalId != "" {
			getConfigResources(region, crossAccountRoleArn, acKey, secKey, sessionName, externalId)
		} else {
			cmd.Help()
			return
		}

	},
}

func getConfigResources(region string, crossAccountRoleArn string, accessKey string, secretKey string, sessionName string, externalId string) (*configservice.GetDiscoveredResourceCountsOutput, error) {
	log.Println("Getting aws config resource count summary")
	configServiceClient, err := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, sessionName, externalId)
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

func GetConfig(region string, crossAccountRoleArn string, accessKey string, secretKey string, externalId string) (*configservice.GetDiscoveredResourceCountsOutput, error) {
	sessionName := util.RamdomString(5)
	response, err := getConfigResources(region, crossAccountRoleArn, accessKey, secretKey, sessionName, externalId)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("failed to load config summary. %v", err)
	}
	return response, nil
}

func GetCloudConfigSummary(region string, vaultUrl string, vaultToken string, accountNo string) (*configservice.GetDiscoveredResourceCountsOutput, error) {
	if region == "" {
		log.Println("region is missing")
		return nil, fmt.Errorf("region is missing")
	} else if vaultUrl == "" {
		log.Println("vault url is missing")
		return nil, fmt.Errorf("vault url is missing")
	} else if accountNo == "" {
		log.Println("account no is missing ")
		return nil, fmt.Errorf("account no is missing")
	} else if vaultToken == "" {
		log.Println("vault token is missing ")
		return nil, fmt.Errorf("vault token is missing")
	}
	log.Println("Getting account details")
	vaultResp, err := vault.GetAccountDetails(vaultUrl, vaultToken, accountNo)
	if vaultResp.Data.AccessKey == "" || vaultResp.Data.SecretKey == "" || vaultResp.Data.CrossAccountRoleArn == "" || vaultResp.Data.ExternalId == "" {
		log.Println("Account details not found")
		return nil, err
	}
	sessionName := util.RamdomString(5)
	response, err := getConfigResources(region, vaultResp.Data.CrossAccountRoleArn, vaultResp.Data.AccessKey, vaultResp.Data.SecretKey, sessionName, vaultResp.Data.ExternalId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return response, nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := AwsxCloudElementsCmd.Execute()
	if err != nil {
		log.Fatal("There was some error while executing the CLI: ", err)
		return
	}
}

func init() {
	AwsxCloudElementsCmd.Flags().String("vaultUrl", "", "vault end point")
	AwsxCloudElementsCmd.Flags().String("vaultToken", "", "vault token")
	AwsxCloudElementsCmd.Flags().String("accountId", "", "aws account number")
	AwsxCloudElementsCmd.Flags().String("zone", "", "aws region")
	AwsxCloudElementsCmd.Flags().String("accessKey", "", "aws access key")
	AwsxCloudElementsCmd.Flags().String("secretKey", "", "aws secret key")
	AwsxCloudElementsCmd.Flags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxCloudElementsCmd.Flags().String("externalId", "", "aws external id")
}
