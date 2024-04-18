package utils

import (
	"context"
	"fmt"
	"honeypot/admin"
	"os"
	"os/signal"
	"time"
)

func Stop() {
	//优雅关闭服务
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	// quit 信道是同步信道，若没有信号进来，处于阻塞状态
	// 反之，则执行后续代码
	<-quit
	fmt.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 调用 srv.Shutdown() 完成优雅停止
	// 调用时传递了一个上下文对象，对象中定义了超时时间
	//系统关闭时统一结束蜜罐服务
	if err := admin.BurpPot.Shutdown(ctx); err != nil {
		fmt.Println("burpsuite Shutdown:", err)
	}
	if err := admin.VpnPot.Shutdown(ctx); err != nil {
		fmt.Println("vpn Shutdown:", err)
	}
	if err := admin.GobyPot.Shutdown(ctx); err != nil {
		fmt.Println("goby Shutdown:", err)
	}
	if err := Sys.Shutdown(ctx); err != nil {
		fmt.Println("Server Shutdown:", err)
	}
	fmt.Println("Server exiting")
}
