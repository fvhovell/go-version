# go-version
Small package to provide version-like metadata to any application burnt into the binary

## Usage

To use this package, clone it into a subfolder of your go workspace, e.g.:

    cd $GOPATH
    go get github.com/fvhovell/go-version
    src/github.com/fvhovell/go-version/make.sh

Demo application will be available under $GOPATH/bin/version-test
Sample output:

    2016/01/03 18:49:19 Hello World:
    Appication:
    	Name    : version-test
    	Version : local-test-build
    Build data:
    	Type       : integration-build
    	Timestamp  : 2016-01-03 18:49:17 +0100 CET
    	Hostname   : localhost.local
    	Username   : builduser
    	Go version : go1.5.1-darwin/amd64
    Git:
    	Branch : master
    	Commit : 9818ee7f8ac3d57da9fe25d5d5c8848a46475f86
    {"app_name":"version-test","app_version":"local-test-build","build_type":"integration-build","build_time":"2016-01-03T18:49:17+01:00","build_hostname":"localhost.local","build_username":"builduser","build_go_version":"go1.5.1-darwin/amd64","git_branch":"master","git_commit":"9818ee7f8ac3d57da9fe25d5d5c8848a46475f86"}
