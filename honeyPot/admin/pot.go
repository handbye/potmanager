package admin

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"honeypot/admin/tools"
	"net/http"
)

type PotData struct {
	Name     string
	Pottype  string
	State    int
	Url      string
	Configid int
}

func Pot(c *gin.Context) {
	islogin := tools.GetSession(c)
	//name = c.PostForm("name")
	//pottype = c.PostForm("pottype")
	//state = c.PostForm("state")
	//configid = c.PostForm("configid")
	res := PotRead()
	if islogin {
		if c.Request.Method == "GET" {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"IsLogin": islogin,
				"res":     res,
			})
		}
	} else {
		c.HTML(http.StatusOK, "login.html", nil)
	}
}

func PotRead() (Pot []PotData) {
	db, err := sql.Open("sqlite3", tools.DbPath)
	checkErr(err)
	// 查询数据
	rows, err := db.Query("SELECT name, pottype, state, url,configid FROM pot order by id ASC")
	checkErr(err)
	var result = make([]PotData, 0)
	for rows.Next() {
		var (
			name, pottype string
			state         int
			url           string
			configid      int
		)
		err = rows.Scan(&name, &pottype, &state, &url, &configid)
		if err == nil {
			result = append(result, PotData{name, pottype, state, url, configid})
		} else {
			result = append(result, PotData{"", "", 0, "#", 0})
		}
		db.Close()

	}
	return result
}

func Exit() {
	db, err := sql.Open("sqlite3", tools.DbPath)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	querysql := "SELECT count(1) FROM sqlite_master WHERE type='table' AND name='pot'"
	row, err := db.Query(querysql)
	if err != nil {
		fmt.Println(err)
	}

	var num int
	for row.Next() {
		err = row.Scan(&num)
		if err != nil {
			fmt.Println(err)
		}
	}

	sqlStr := "DROP TABLE IF EXISTS 'goose_db_version';"
	if num == 1 {
		sqlStr = "UPDATE pot SET state = 0;DROP TABLE IF EXISTS 'goose_db_version';"
	}
	_, err = db.Exec(sqlStr)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("数据库状态恢复成功！")
}
