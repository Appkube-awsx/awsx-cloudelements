- [What is awsx-cloudelements](#awsx-cloudelements)
- [How to write plugin subcommand](#how-to-write-plugin-subcommand)
- [How to build / Test](#how-to-build--test)
- [what it does ](#what-it-does)
- [command input](#command-input)
- [command output](#command-output)
- [How to run ](#how-to-run)

# awsx-cloudelements
This is a plugin subcommand for awsx cli ( https://github.com/Appkube-awsx/awsx#awsx ) cli.

For details about awsx commands and how its used in Appkube platform , please refer to the diagram below:

![alt text](https://raw.githubusercontent.com/AppkubeCloud/appkube-architectures/main/LayeredArchitecture-phase2.svg)


# How to write plugin subcommand 
Please refer to the instaruction -
https://github.com/Appkube-awsx/awsx#how-to-write-a-plugin-subcommand

It has detailed instruction on how to write a subcommand plugin , build / test / debug  / publish and integrate into the main commmand.

# How to build / Test
            - With Vault URL    
            go run .\main.go --zone=us-east-1 --vaultUrl=http://<server-address>:<port>/v1/kv-v1 --accountId=SYNECTIKS/HR/AWS/1111111111 --vaultToken=xxxxxxxxxx
            or
            go run .\main.go getElementDetails --zone=us-east-1 --vaultUrl=http://<server-address>:<port>/v1/kv-v1 --accountId=SYNECTIKS/HR/AWS/111111111 --vaultToken=xxxxxxxxxx
            
            - With Access/Secret key
            go run .\main.go --zone=us-east-1 --accessKey=######### --secretKey=######### --externalId=######### --crossAccountRoleArn=#########
            or
            go run .\main.go  getElementDetails --zone=us-east-1 --accessKey=######### --secretKey=######### --externalId=######### --crossAccountRoleArn=#########

            Another way of testing is by running go install command
            go install
            - go install command creates an exe with the name of the module (e.g. awsx-cloudelements) and save it in the GOPATH
            - Now we can execute this command on command prompt as below
            - With Vault URL    
            awsx-cloudelements --zone=us-east-1 --vaultUrl=http://<server-address>:<port>/v1/kv-v1 --accountId=SYNECTIKS/HR/AWS/111111111 --vaultToken=xxxxxxxxxx
            - With Access/Secret key
            awsx-cloudelements --zone=us-east-1 --accessKey=######### --secretKey=######### --externalId=######### --crossAccountRoleArn=#########

# what it does 
This subcommand implement the following functionalities
  - getElementDetails - It  will get the resource count summary for a given AWS account id

# command input
  - valutURL = URL location of vault - that stores credentials to call API
  - acountId = The AWS account id.
  - zone = AWS region
  - vaultToken = Vault token to access vault

#  command output
{
        ResourceCounts: [
            {
                Count: 124,
                ResourceType: "AWS::S3::Bucket"
            },
            {
                Count: 121,
                ResourceType: "AWS::Lambda::Function"
            },
            {
                Count: 72,
                ResourceType: "AWS::CloudFormation::Stack"
            },
            {
                Count: 50,
                ResourceType: "AWS::CloudWatch::Alarm"
            }
        ],
        TotalDiscoveredResources: 367
}

# How to run 
From main awsx command , it is called as follows:
  - awsx getElementDetails  --zone=us-east-1 --vaultUrl=http://<server-address>:<port>/v1/kv-v1 --accountId=SYNECTIKS/HR/AWS/1111111111 --vaultToken=xxxxxxxxxx
  
    OR
  - awsx getElementDetails  --zone=us-east-1 --accessKey=######### --secretKey=######### --externalId=######### --crossAccountRoleArn=#########

If you build it locally , you can simply run it as standalone command as 
  - awsx-cloudelements --zone=us-east-1 --vaultUrl=http://<server-address>:<port>/v1/kv-v1 --accountId=SYNECTIKS/HR/AWS/1111111111 --vaultToken=xxxxxxxxxx

    OR
  - awsx-cloudelements  --zone=us-east-1 --accessKey=######### --secretKey=######### --externalId=######### --crossAccountRoleArn=#########

