# Download a go module from GIT and use it in other modules 
A go module can be downloaded from git by using go get command
       
# Example

# Create a git repository
    https://github.com/Appkube-awsx/awsx

# Clone this repository
    git clone https://github.com/Appkube-awsx/awsx.git

# Go in the awsx directory and initialize the project
    go mod init awsx

# Download the dependency 
    go get github.com/Appkube-awsx/awsx-cloudelements@v1.0.0
    - In the above go get command (go get github.com/Appkube-awsx/awsx-cloudelements@v1.0.0) we have specified the version (v1.0.0).
      This version is the git tag.
    - If we call the go get command without tag/version, it will download the latest version from GIT

