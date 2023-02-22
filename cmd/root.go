/*
Copyright © 2023 Manoj Sharma manoj.sharma@synectiks.com
*/
package cmd

import (
	"github.com/Appkube-awsx/awsx-cloudelements/client"
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
		accountNo, _ := cmd.Flags().GetString("accountId")
		region, _ := cmd.Flags().GetString("zone")
		acKey, _ := cmd.Flags().GetString("accessKey")
		secKey, _ := cmd.Flags().GetString("secretKey")
		crossAccountRoleArn, _ := cmd.Flags().GetString("crossAccountRoleArn")

		if vaultUrl != "" && accountNo != "" {
			if region == "" {
				cmd.Help()
				return
			}
			log.Println("Getting account details")
			data, err := vault.GetAccountDetails(vaultUrl, accountNo)
			if err != nil {
				log.Println("Error in calling the account details api. \n", err)
				return
			}
			if data.AccessKey == "" || data.SecretKey == "" || data.CrossAccountRoleArn == "" {
				log.Println("Account details not found.")
				return
			}
			getConfigResources(region, data.CrossAccountRoleArn, data.AccessKey, data.SecretKey)
		} else if region != "" && acKey != "" && secKey != "" && crossAccountRoleArn != "" {
			getConfigResources(region, crossAccountRoleArn, acKey, secKey)
		} else {
			cmd.Help()
			return
		}

	},
}

func getConfigResources(region string, crossAccountRoleArn string, accessKey string, secretKey string) *configservice.GetDiscoveredResourceCountsOutput {
	log.Println("Getting aws config resource count summary")
	configServiceClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey)
	configResourceRequest := &configservice.GetDiscoveredResourceCountsInput{}
	configResourceResponse, err := configServiceClient.GetDiscoveredResourceCounts(configResourceRequest)
	if err != nil {
		log.Fatalln("Error: ", err)
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
		return
	}
}

func init() {
	AwsxCloudElementsCmd.Flags().String("vaultUrl", "", "vault end point")
	AwsxCloudElementsCmd.Flags().String("accountId", "", "aws account number")
	AwsxCloudElementsCmd.Flags().String("zone", "", "aws region")
	AwsxCloudElementsCmd.Flags().String("accessKey", "", "aws access key")
	AwsxCloudElementsCmd.Flags().String("secretKey", "", "aws secret key")
	AwsxCloudElementsCmd.Flags().String("crossAccountRoleArn", "", "aws cross account role arn")
}
