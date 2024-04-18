package admin

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"honeypot/admin/tools"
	"net/http"
	"strconv"
	"strings"
)

type PotConfigs struct {
	Configid   int
	Port       int
	Payload    string
	Fileexists int
	Username   string
	Password   string
	Filelist   string
	Ip         string
}

// PotConfig 重启系统后需初始化相关表
func PotConfig(c *gin.Context) {
	islogin := tools.GetSession(c)
	if islogin {
		if c.Request.Method == "GET" {
			configid, _ := strconv.Atoi(c.Query("configid"))
			res := ConfigRead(configid)
			res.Filelist = tools.ZeroToNull(res.Filelist)
			res.Payload = tools.ZeroToNull(res.Payload)
			res.Username = tools.ZeroToNull(res.Username)
			res.Password = tools.ZeroToNull(res.Password)
			res.Ip = tools.ZeroToNull(res.Ip)
			filename := ""

			if configid == 1002 {
				_, err := tools.PathExists(tools.VpnFile)
				if err == nil {
					ss := strings.Split(tools.VpnFile, "/")
					filename = ss[len(ss)-1]
				}
			}

			c.HTML(http.StatusOK, "potconfig.html", gin.H{
				"filename": filename,
				"res":      res,
			})
		}
		if c.Request.Method == "POST" {
			username := tools.Strip(tools.ZeroToNull(c.PostForm("username")))
			password := tools.Strip(tools.ZeroToNull(c.PostForm("password")))
			port, _ := strconv.Atoi(tools.Strip(c.PostForm("port")))
			filelist := tools.ZeroToNull(c.PostForm("filelist"))
			configid, _ := strconv.Atoi(tools.Strip(c.PostForm("configid")))
			payload := tools.ZeroToNull(c.PostForm("payload"))
			fileexists, _ := strconv.Atoi(c.PostForm("fileexists"))
			ip := tools.Strip(tools.ZeroToNull(c.PostForm("ip")))

			config1 := PotConfigs{configid, port, payload, fileexists, username, password, filelist, ip}
			code := ConfigSave(config1)
			c.JSON(http.StatusOK, gin.H{"code": code})
		}
	} else {
		c.HTML(http.StatusOK, "login.html", nil)
	}
}

func ConfigRead(configid int) (PotConfig1 PotConfigs) {
	db, err := sql.Open("sqlite3", tools.DbPath)
	checkErr(err)
	defer db.Close()

	querystring := fmt.Sprintf("SELECT * FROM pot_config WHERE configid = %d", configid)
	rows, err := db.Query(querystring)
	checkErr(err)
	result := PotConfigs{configid, 0, "", 0, "", "", "", ""}
	for rows.Next() {
		var (
			port       int
			payload    string
			fileexists int
			username   string
			password   string
			filelist   string
			ip         string
		)
		err = rows.Scan(&configid, &port, &payload, &fileexists, &username, &password, &filelist, &ip)
		if err == nil {
			result = PotConfigs{configid, port, payload, fileexists, username, password, filelist, ip}
		}

	}
	return result
}

func ConfigSave(PotConfig1 PotConfigs) (code int) {
	db, err := sql.Open("sqlite3", tools.DbPath)
	checkErr(err)
	defer db.Close()

	sqlStr := "UPDATE pot_config SET port = ?,payload = ?,fileexists = ?,username = ?,password = ?,filelist = ?,ip = ? WHERE configid = ?"
	stmt, err := db.Prepare(sqlStr)
	checkErr(err)

	_, err = stmt.Exec(PotConfig1.Port, PotConfig1.Payload, PotConfig1.Fileexists, PotConfig1.Username, PotConfig1.Password, PotConfig1.Filelist, PotConfig1.Ip, PotConfig1.Configid)
	if err != nil {
		checkErr(err)
	} else {
		return 0
	}

	return 1
}
