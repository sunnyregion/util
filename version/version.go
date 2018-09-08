package version

import (
	"flag"
	"fmt"
	"os"
)

var (
	VERSION    string
	BUILD_TIME string
	BUILD_NAME string
	GO_VERSION string
)

func init() {
	var showVer bool
	flag.BoolVar(&showVer, "v", false, "show build version")
	flag.Parse()
	if showVer {
		fmt.Printf("Build Name:\t%s\n", BUILD_NAME)
		fmt.Printf("Build Version:\t%s\n", VERSION)
		fmt.Printf("Build Time:\t%s\n", BUILD_TIME)
		fmt.Printf("Git ID:\t%s\n", GO_VERSION)
		fmt.Println(`-----🎉🎉🎉👍💁👌⚽🎍😍🎉🎉🎉------`)
		os.Exit(0)
	}
}
