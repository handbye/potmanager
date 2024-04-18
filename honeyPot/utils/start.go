package utils

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"honeypot/admin"
	"log"
	"net/http"
	"time"
)

var (
	g errgroup.Group
)

var Vip *string
var Vport *string

var Sys = &http.Server{}

func Start() {

	fmt.Println("平台正在启动,请稍等...")

	fmt.Printf("后台地址是：【 %s 】 ,请保存，此地址每次启动都会更改！\n", admin.Adminurl)

	Sys = &http.Server{
		Addr:         *Vip + ":" + *Vport,
		Handler:      admin.Admin(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	admin.G.Go(func() error {
		return Sys.ListenAndServe()
	})

	if err := admin.G.Wait(); err != nil {
		log.Fatal(err)
	}
}
