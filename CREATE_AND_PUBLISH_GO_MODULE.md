# Create independent go module and Publish it to GIT
    This documentation provides help to create go cli and publish it to git as independent module. 
    Other modules if needs this functionality can download it from GIT

# Create a git repository
    https://github.com/Appkube-awsx/awsx-cloudelements

# Clone this repository 
    git clone https://github.com/Appkube-awsx/awsx-cloudelements.git

# Go in the awsx-cloudelements directory and execute the following commands
    1. initialize the project
        go mod init github.com/Appkube-awsx/awsx-cloudelements
    2. download the latest version of cobra cli
        go get github.com/spf13/cobra@latest
    3. install the Cobra cli
	    go install github.com/spf13/cobra-cli@latest
    4. execute cobra-cli init command. This command will generate the application with the correct file structure and imports:
        cobra-cli init
    
# The above command will create directory structure as below and generate the basic cli code in root.go and main.go
	
	awsx-cloudelements
	|
	|__cmd
        |__root.go
	|__main.go

    In the root.go you will find the code as below
	
        var rootCmd = &cobra.Command{
            Use:   "aws-cloudelements",
            Short: "A brief description of your application",
            Long: `A longer description that spans multiple lines and likely contains
                    examples and usage of using your application. For example:
    
                    Cobra is a CLI library for Go that empowers applications.
                    This application is a tool to generate the needed files
                    to quickly create a Cobra application.`,
            // Uncomment the following line if your bare application
            // has an action associated with it:
            Run: func(cmd *cobra.Command, args []string) {
                fmt.Println("Calling aws-cloudelements")
            },
        }

        func Execute() {
            err := rootCmd.Execute()
            if err != nil {
                log.Fatal("There was some error while executing the CLI: ", err)
                os.Exit(1)
            }
        }

        func init() {
            
        }

        - In the Run inline function we should write our cli code. 
        - In our example we have written a fmt.Println("Calling aws-cloudelements")
        - When we execute this command, this message will be printed on console
    
    In main.go we should call the command. So the main.go should be as below:
    
        package main

        import "github.com/Appkube-awsx/awsx-cloudelements/cmd"

        func main() {
            cmd.Execute()
        }

# After writing the code we can test the code as below
    go run main.go
    - Program will print Calling aws-cloudelements on console 
    
    Another way of testing is by running go install command
    go install
    - go install command creates an exe with the name of the module (e.g. awsx-cloudelements) and save it in the GOPATH
    - Now we can execute this command on command prompt as below
    awsx-cloudelements

# Publish the code in git so that other modules can download this code as dependency from git
    1. Commit and push the code
    2. Tag the code. Use the following git commands to tag it
        git tag "v1.0.0"
	    git push --tags
        
    3. Developers interested in this module, import it by running the go get command as below
        go get github.com/Appkube-awsx/awsx-cloudelements@v1.0.0
        
        - In the above go get command (go get github.com/Appkube-awsx/awsx-cloudelements@v1.0.0) we have specified the version (v1.0.0).
          This version is the git tag, what we specified in the git tag command earlier. 