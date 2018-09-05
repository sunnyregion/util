package version

import (
	"flag"
	"fmt"
	"os"
)

var (
	BuildVersion string
	BuildTime    string
	BuildName    string
	CommitID     string
)

func init() {
	var showVer bool
	flag.BoolVar(&showVer, "v", false, "show build version")
	flag.Parse()
	if showVer {
		// Printf( "build name:\t%s\nbuild ver:\t%s\nbuild time:\t%s\nCommitID:%s\n", BuildName, BuildVersion, BuildTime, CommitID )
		fmt.Printf("build name:\t%s\n", BuildName)
		fmt.Printf("build ver:\t%s\n", BuildVersion)
		fmt.Printf("build time:\t%s\n", BuildTime)
		fmt.Printf("Commit ID:\t%s\n", CommitID)
		os.Exit(0)
	}
}
