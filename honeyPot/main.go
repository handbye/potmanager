package main

import (
	"flag"
	"fmt"
	"honeypot/admin"
	"honeypot/admin/db"
	"honeypot/admin/tools"
	"honeypot/utils"
	"os"
)

var Vinit *bool
var Vstart *bool

func init() {
	tools.CreateUploadDic()
	admin.Exit()
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	Vinit = flag.Bool("init", false, "初始化数据")
	Vstart = flag.Bool("start", false, "启动平台")
	utils.Vip = flag.String("ip","0.0.0.0","平台启动IP")
	utils.Vport = flag.String("port","80","平台启动端口")
	flag.Parse()
}

func main() {
	_, err := os.Stat(tools.DbPath)
	if err != nil {
		fmt.Printf("\"数据库文件不存在,请进行初始化操作,使用 %s -h 查看操作方法\\n\",os.Args[0]")
		os.Exit(1)
	}
	if *Vinit {
		db.DbInit()
	}
	if *Vstart {
		utils.Start()
		defer utils.Stop()
	}
}
