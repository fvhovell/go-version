// model.go
package version

import (
	"fmt"
	"time"
)

type MetaDataModel struct {
	AppName        string    `json:"app_name"`
	AppVersion     string    `json:"app_version"`
	BuildType      string    `json:"build_type"`
	BuildTime      time.Time `json:"build_time"`
	BuildHostname  string    `json:"build_hostname"`
	BuildUsername  string    `json:"build_username"`
	BuildGoVersion string    `json:"build_go_version"`
	GitBranch      string    `json:"git_branch"`
	GitCommit      string    `json:"git_commit"`
}

func (m *MetaDataModel) String() string {
	return fmt.Sprintf("Appication:\n"+
		"\tName    : %s\n"+
		"\tVersion : %s\n"+
		"Build data:\n"+
		"\tType       : %s\n"+
		"\tTimestamp  : %v\n"+
		"\tHostname   : %s\n"+
		"\tUsername   : %s\n"+
		"\tGo version : %s\n"+
		"Git:\n"+
		"\tBranch : %s\n"+
		"\tCommit : %s\n",
		m.AppName,
		m.AppVersion,
		m.BuildType,
		m.BuildTime,
		m.BuildHostname,
		m.BuildUsername,
		m.BuildGoVersion,
		m.GitBranch,
		m.GitCommit)
}
