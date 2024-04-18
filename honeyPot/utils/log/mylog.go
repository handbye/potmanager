package mylog

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"honeypot/admin/tools"
	"io"
	"os"
	"path"
	"strings"
	"time"
)

// 更改日志格式


func Logger(filepath string) *logrus.Logger {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/" + filepath
	}
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
	}
	logFileName := now.Format("2006-01-02") + ".log"
	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println(err.Error())
		}
	}
	//写入控制台
	writer1 := os.Stdout
	//写入文件
	writer2, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	//实例化
	logger := logrus.New()

	//设置输出
	logger.SetOutput(io.MultiWriter(writer1, writer2))

	//设置日志级别
	logger.SetLevel(logrus.InfoLevel)

	//设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return logger
}

func HttpLog( filepath string, tablename string) gin.HandlerFunc {
	logger := Logger(filepath)
	return func(c *gin.Context) {
		// 处理请求
		c.Next()

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		if !strings.Contains(reqUri,"jkxtmw")  {
			requestlog := logger.WithFields(logrus.Fields{"reqMethod":reqMethod,"statusCode":statusCode,"clientIP":clientIP,"reqUri":reqUri})
			requestlog.Info()
			SqlHook(logger,tablename)
		}
	}
}


func NoHttpLog(filepath string, tablename string) *logrus.Logger{
	logger := Logger(filepath)
	SqlHook(logger,tablename)
	return logger
}

func SqlHook( logger *logrus.Logger, tablename string){
	timeout := time.Second * 10
	db, err := sql.Open("sqlite3", tools.DbPath)
	if err != nil {
		fmt.Printf("Unable to open database: %s\n", err)
	}
	//fmt.Println(("Database opened successful"))
	logger.SetLevel(logrus.InfoLevel)
	hook, err := NewSQLiteHook(db, tablename,timeout)
	if err != nil {
		fmt.Printf("Unable to initialize hook: %s\n", err)
	}
	//fmt.Println("Hook initialized successful")
	logger.AddHook(hook)
}
