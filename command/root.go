/*
Copyright © 2023 Manoj Sharma manoj.sharma@synectiks.com
*/
package command

import (
	"github.com/Appkube-awsx/awsx-cloudelements/command/appconfigcmd"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/spf13/cobra"
	"log"
)

// AwsxCloudElementsCmd represents the base command when called without any subcommands
var AwsxCloudElementsCmd = &cobra.Command{
	Use:   "getElementDetails",
	Short: "getElementDetails command gets resource counts",
	Long:  `getElementDetails command gets resource counts`,

	Run: func(cmd *cobra.Command, args []string) {
		authFlag, clientAuth, err := authenticate.CommandAuth(cmd)
		if err != nil {
			cmd.Help()
			return
		}
		if authFlag {
			appconfigcmd.GetDiscoveredResourceCounts(*clientAuth)
		} else {
			cmd.Help()
			return
		}
	},
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
	AwsxCloudElementsCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxCloudElementsCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxCloudElementsCmd.PersistentFlags().String("accountId", "", "aws account number")
	AwsxCloudElementsCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxCloudElementsCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxCloudElementsCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxCloudElementsCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxCloudElementsCmd.PersistentFlags().String("externalId", "", "aws external id")
}
