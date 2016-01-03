// init.go
package version

import (
	"time"
)

const BUILD_TS_LAYOUT = "2006-01-02T15:04:05-0700"

var (
	// Below definitions are default / fallback values for the version data
	// These can (and should) be overwritten on the commandline during building, e.g. using:
	/*
	   VERSION_PKG="github.com/fvhovell/go-version" \
	   go build -ldflags "\
	   -X ${VERSION_PKG}/version.AppName=my-application \
	   -X ${VERSION_PKG}/version.AppVersion=local-test-build \
	   -X ${VERSION_PKG}/version.BuildType=integration-build \
	   -X ${VERSION_PKG}/version.BuildTs=$(date +%FT%T%z) \
	   -X ${VERSION_PKG}/version.BuildHostname=$HOSTNAME \
	   -X ${VERSION_PKG}/version.BuildUsername=$USER \
	   -X ${VERSION_PKG}/version.BuildGoVersion=$(go version | sed -e 's/go version //; s/ /_/g;') \
	   -X ${VERSION_PKG}/version.GitHash=$(cd src/${VERSION_PKG} && git rev-parse HEAD) \
	   -X ${VERSION_PKG}/version.GitBranch=$(cd src/${VERSION_PKG} && git symbolic-ref --short HEAD) \
	   " ${VERSION_PKG}/...
	*/
	AppName        = "unnamed-application"
	AppVersion     = "undefined-version"
	BuildType      = "development"
	BuildTs        = "1970-01-01T00:00:00+0000"
	BuildHostname  = "localhost"
	BuildUsername  = "me"
	BuildGoVersion = "undefined-go-version"
	GitBranch      = "undefined-branch"
	GitCommit      = "undefined-commit"

	MetaData *MetaDataModel
)

func init() {
	buildTime, err := time.Parse(BUILD_TS_LAYOUT, BuildTs)
	if err != nil {
		panic(err)
	}
	MetaData = &MetaDataModel{
		AppName:        AppName,
		AppVersion:     AppVersion,
		BuildType:      BuildType,
		BuildTime:      buildTime,
		BuildHostname:  BuildHostname,
		BuildUsername:  BuildUsername,
		BuildGoVersion: BuildGoVersion,
		GitBranch:      GitBranch,
		GitCommit:      GitCommit,
	}
}
