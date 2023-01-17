/*
Copyright © 2023 Manoj Sharma manoj.sharma@synectiks.com
*/
package cmd

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-cloudelements/client"
	"github.com/Appkube-awsx/awsx-cloudelements/vault"
	"github.com/aws/aws-sdk-go/service/configservice"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// AwsxCloudElementsCmd represents the base command when called without any subcommands
var AwsxCloudElementsCmd = &cobra.Command{
	Use:   "aws-cloudelements",
	Short: "aws aws-cloudelements details",
	Long:  `aws aws-cloudelements details`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Calling aws-cloudelements summary api")
		vaultUrl, _ := cmd.Flags().GetString("vaultUrl")
		accountNo, _ := cmd.Flags().GetString("ac")
		region, _ := cmd.Flags().GetString("region")
		acKey, _ := cmd.Flags().GetString("accessKey")
		secKey, _ := cmd.Flags().GetString("secretKey")
		crossAccountRoleArn, _ := cmd.Flags().GetString("crossAccountRoleArn")

		if accountNo == "" && region == "" && acKey == "" && secKey == "" && crossAccountRoleArn == "" {
			fmt.Println("AWS credentials like account number or accesskey/secretkey/region/crossAccountRoleArn not provided")
			return
		}

		if vaultUrl != "" {
			if accountNo == "" {
				fmt.Println("AWS account number not provided")
				return
			}
			fmt.Println("Account number provided. Calling API to get account details")
			// call rest api to get accesGetAccountDetailssKey/secretKey/crossAccountRoleArn/region
			data, err := vault.GetAccountDetails(vaultUrl, accountNo)
			if err != nil {
				fmt.Println("Error while calling the account details api. Error ", err)
				return
			}
			if data.AccessKey == "" {
				log.Println("Account credentials not found.")
				return
			}
			getConfigResources(data.Region, data.CrossAccountRoleArn, data.AccessKey, data.SecretKey)
		} else {
			if region == "" || acKey == "" || secKey == "" || crossAccountRoleArn == "" {
				fmt.Println("AWS credentials like accesskey/secretkey/region/crossAccountRoleArn not provided")
				return
			}
			fmt.Println("Getting aws appconfig summary")
			getConfigResources(region, crossAccountRoleArn, acKey, secKey)
		}

	},
}

func getConfigResources(region string, crossAccountRoleArn string, accessKey string, secretKey string) *configservice.GetDiscoveredResourceCountsOutput {
	configServiceClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey)
	configResourceRequest := &configservice.GetDiscoveredResourceCountsInput{}
	configResourceResponse, err := configServiceClient.GetDiscoveredResourceCounts(configResourceRequest)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	log.Println(configResourceResponse)
	return configResourceResponse
}

func GetConfig(region string, crossAccountRoleArn string, accessKey string, secretKey string) *configservice.GetDiscoveredResourceCountsOutput {
	return getConfigResources(region, crossAccountRoleArn, accessKey, secretKey)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := AwsxCloudElementsCmd.Execute()
	if err != nil {
		log.Fatal("There was some error while executing the CLI: ", err)
		os.Exit(1)
	}
}

func init() {
	AwsxCloudElementsCmd.Flags().String("vaultUrl", "", "vault end point")
	AwsxCloudElementsCmd.Flags().String("ac", "", "aws account number")
	AwsxCloudElementsCmd.Flags().String("region", "", "aws region")
	AwsxCloudElementsCmd.Flags().String("accessKey", "", "aws access key")
	AwsxCloudElementsCmd.Flags().String("secretKey", "", "aws secret key")
	AwsxCloudElementsCmd.Flags().String("crossAccountRoleArn", "", "aws cross account role arn")
}
