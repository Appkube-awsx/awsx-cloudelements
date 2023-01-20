# GO CLI **getElementDetails**
    This documentation provides help to create go cli getElementDetails and publish it to git as independent module. 

    getElementDetails --vaultURL=vault.dummy.net --accountId=xxxxxxxxxx --zone=us-west-2
    

# Command structure
    getElementDetails  --vaultURL=vault.dummy.net --accountId=xxxxxxxxxx --zone=us-west-2
    
> Command \
    **getElementDetails**: getElementDetails command will get the resource count summary for a given AWS account id and regions 
    
> Command Parameters \
    **--vaultURL**: It is a server url and used to store user cloud credentials. It stores information in key=value pair in encrypted format. \
    **--accountId**: It is a AWS account id. We pass this account id as key to the vault serve to get user credentials \
    **--zone**: This parameter is the AWS region

# Response
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
        TotalDiscoveredResources: 809
    }

# [How to create and publish a go module](CREATE_AND_PUBLISH_GO_MODULE.md)
# [Command Params - Flags](https://dev.to/divrhino/adding-flags-to-a-command-line-tool-built-with-go-and-cobra-34f1)