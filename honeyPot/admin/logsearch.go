package admin

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"honeypot/admin/tools"
	"net/http"
)

func LogSearch(c *gin.Context){
	islogin := tools.GetSession(c)
	starttime := c.PostForm("startDate")
	endtime := c.PostForm("endDate")
	ip := c.PostForm("ip")
	method := c.PostForm("reqmethod")
	tablename := c.PostForm("tablename")
	if tablename == ""{
		tablename = "log"
	}
	if tools.In(tablename, tools.Config("httplog")){
		res, err := HttpLogSearchQuery(tablename,starttime,endtime,ip,method)
		if islogin{
			if err !=nil || len(res)==0{
				res = append(res, HttpLogData{"","","","","",""})
			}
			if len(res)!=0{
				if c.Request.Method == "POST"{
					c.JSON(http.StatusOK,gin.H{
						"res" : res,
					})
				}
			}
		}
	}
	if tools.In(tablename, tools.Config("nohttplog")){
		res, err := NoHttpLogSearchQuery(tablename,starttime,endtime)
		if islogin{
			if err !=nil || len(res)==0{
				res = append(res, NohttpLogData{"",""})
			}
			if len(res)!=0{
				if c.Request.Method == "POST"{
					c.JSON(http.StatusOK,gin.H{
						"res" : res,
					})
				}
			}
		}
	}
}

func HttpLogSearchQuery(tablename string,starttime string ,endtime string ,ip string, method string) (data []HttpLogData, e error){
	tools.SafeDate(starttime)
	tools.SafeDate(endtime)
	tools.SafeDate(ip)
	tools.SafeDate(method)
	db, err := sql.Open("sqlite3", tools.DbPath)
	checkErr(err)
	// 查询数据
	var querystring string
	var result = make([]HttpLogData, 0)
	if starttime == "" && endtime =="" && ip == "" && method == ""{
		querystring = fmt.Sprintf("SELECT time,clientIp,statusCode,reqMethod,reqUri,full_message FROM %s ORDER BY id DESC",tablename)
	}
	if starttime == "" && endtime =="" && ip != "" && method == ""{
		querystring = fmt.Sprintf("SELECT time,clientIp,statusCode,reqMethod,reqUri,full_message FROM %s  WHERE clientIP='%s' ORDER BY id DESC",tablename,ip)
	}
	if starttime == "" && endtime =="" && ip == "" && method != ""{
		querystring = fmt.Sprintf("SELECT time,clientIp,statusCode,reqMethod,reqUri,full_message FROM %s  WHERE reqMethod='%s' ORDER BY id DESC",tablename,method)
	}
	if starttime == "" && endtime =="" && ip != "" && method != ""{
		querystring = fmt.Sprintf("SELECT time,clientIp,statusCode,reqMethod,reqUri,full_message FROM %s  WHERE clientIP='%s' and reqMethod='%s' ORDER BY id DESC",tablename,ip,method)
	}
	if starttime != "" && endtime !="" && ip == "" && method == ""{
		querystring = fmt.Sprintf("SELECT time,clientIp,statusCode,reqMethod,reqUri,full_message FROM %s  WHERE time >='%s' AND Time <='%s' ORDER BY id DESC",tablename,starttime,endtime)
	}
	if starttime != "" && endtime !="" && ip != "" && method != ""{
		querystring = fmt.Sprintf("SELECT time,clientIp,statusCode,reqMethod,reqUri,full_message FROM %s  WHERE time >='%s' AND Time <='%s' AND clientIP='%s' and reqMethod='%s' ORDER BY id DESC",tablename,starttime,endtime,ip,method)
	}
	if starttime != "" && endtime !="" && ip == "" && method != ""{
		querystring = fmt.Sprintf("SELECT time,clientIp,statusCode,reqMethod,reqUri,full_message FROM %s  WHERE time >='%s' AND Time <='%s' AND reqMethod='%s' ORDER BY id DESC",tablename,starttime,endtime,method)
	}
	if starttime != "" && endtime !="" && ip != "" && method == ""{
		querystring = fmt.Sprintf("SELECT time,clientIp,statusCode,reqMethod,reqUri,full_message FROM %s  WHERE time >='%s' AND Time <='%s' AND clientIP='%s' ORDER BY id DESC",tablename,starttime,endtime,ip)
	}
	rows, err := db.Query(querystring)
	if err != nil{
		fmt.Printf("数据库查询数据失败：%s\n", err)
	}else {
		for rows.Next() {
			var (
				time,clientIp string
				statusCode string
				reqMethod,reqUri,full_message string
			)
			err = rows.Scan(&time, &clientIp, &statusCode ,&reqMethod, &reqUri, &full_message)
			if err ==nil{
				result = append(result, HttpLogData{time,clientIp,statusCode,reqMethod,reqUri,full_message})
			}else {
				result = append(result, HttpLogData{"","","","","",""})
			}
		}
		db.Close()
	}

	return result, nil
}

func NoHttpLogSearchQuery(tablename string,starttime string ,endtime string ) (data []NohttpLogData, e error){
	tools.SafeDate(tablename)
	tools.SafeDate(starttime)
	tools.SafeDate(endtime)
	db, err := sql.Open("sqlite3", tools.DbPath)
	checkErr(err)
	// 查询数据
	var querystring string
	var result = make([]NohttpLogData, 0)
	if starttime == "" && endtime ==""{
		querystring = fmt.Sprintf("SELECT time,msg FROM %s ORDER BY id DESC",tablename)
	}
	if starttime != "" && endtime !=""{
		querystring = fmt.Sprintf("SELECT time,msg FROM %s  WHERE time >='%s' AND Time <='%s' ORDER BY id DESC",tablename,starttime,endtime)
	}
	rows, err := db.Query(querystring)
	if err != nil{
		fmt.Printf("数据库查询数据失败：%s\n", err)
	}else {
		for rows.Next() {
			var (
				time,full_message string
			)
			err = rows.Scan(&time, &full_message)
			if err ==nil{
				result = append(result, NohttpLogData{time,full_message})
			}else {
				result = append(result, NohttpLogData{"",""})
			}
		}
		db.Close()
	}

	return result, nil
}