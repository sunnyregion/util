package version

import (
	"flag"
	"fmt"
	"os"
)

var (
	//VERSION  版本信息
	VERSION string
	//BuildTime 创建时间
	BuildTime string
	//BuildName build的名字
	BuildName string
	//GitID git版本
	GitID string
	//GoVersion 版本
	GoVersion string
	//EMOJI 表情
	EMOJI string
	//GitBranch 版本分支
	GitBranch string
)

func init() {
	var showVer bool
	flag.BoolVar(&showVer, "v", false, "show build version")
	flag.Parse()
	if showVer {
		fmt.Printf("Build Name:\t%s\n", BuildName)
		fmt.Printf("Build Version:\t%s\n", VERSION)
		fmt.Printf("Build Time:\t%s\n", BuildTime)
		fmt.Printf("Git ID:\t\t%s\n", GitID)
		fmt.Printf("Git Branch:\t%s\n", GitBranch)
		fmt.Printf("Go Version:\t%s\n", GoVersion)
		//fmt.Println(`-----🎉🎉🎉👍💁👌⚽🎍😍🎉🎉🎉------`)
		fmt.Println(EMOJI)
		os.Exit(0)
	}
}
