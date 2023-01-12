/*
Copyright © 2023 Manoj Sharma manoj.sharma@synectiks.com
*/
package cmd

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-cloudelements/appconfig/client"
	"github.com/Appkube-awsx/awsx-cloudelements/appconfig/vault"
	"github.com/aws/aws-sdk-go/service/configservice"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// AppconfigCmd represents the base command when called without any subcommands
var AppconfigCmd = &cobra.Command{
	Use:   "appconfig",
	Short: "aws appconfig details",
	Long:  `aws appconfig details`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Calling appconfig summary api")
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

func getConfigResources(region string, crossAccountRoleArn string, accessKey string, secretKey string) {
	configServiceClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey)
	configResourceRequest := &configservice.GetDiscoveredResourceCountsInput{}
	configResourceResponse, err := configServiceClient.GetDiscoveredResourceCounts(configResourceRequest)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	log.Println(configResourceResponse)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := AppconfigCmd.Execute()
	if err != nil {
		log.Fatal("There was some error while executing the CLI: ", err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application. Global means available in child/sub commands

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.appconfig.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//AppconfigCmd.PersistentFlags().String("ac", "", "aws account number")
	//AppconfigCmd.PersistentFlags().String("region", "", "aws region")
	//AppconfigCmd.PersistentFlags().String("accessKey", "", "aws access key")
	//AppconfigCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	//AppconfigCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	//AppconfigCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
