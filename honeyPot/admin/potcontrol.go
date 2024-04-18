package admin

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"honeypot/admin/tools"
	"honeypot/plugins/burpsuite"
	"honeypot/plugins/goby"
	"honeypot/plugins/mysql"
	"honeypot/plugins/vpn"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	G errgroup.Group
	GobyPot *http.Server
	VpnPot *http.Server
	BurpPot *http.Server
)

func PotControl(c *gin.Context){
	islogin := tools.GetSession(c)
	if islogin {
		if c.Request.Method == "POST" {
			configid, _ := strconv.Atoi(tools.Strip(c.PostForm("configid")))
			state, _ := strconv.Atoi(c.PostForm("state"))

			//确认状态一致
			if 0 == checkState(configid, state) {
				//执行实际的关闭开启操作
				code := changePot(configid, state)

				if code == 0 {
					code = changeState(configid, state)
					c.JSON(http.StatusOK, gin.H{"code": code})
					return
				}
			}

			c.JSON(http.StatusOK, gin.H{"code": 5})
		}
	}else {
		c.HTML(http.StatusOK, "login.html", nil)
	}
}

//检查蜜罐状态,port不为0且状态一致
func checkState(configid, state int) (code int) {
	db, err := sql.Open("sqlite3", tools.DbPath)
	checkErr(err)
	defer db.Close()
	code = 1

	querystring := fmt.Sprintf("SELECT p.state,pc.port FROM pot as p, pot_config AS pc WHERE p.configid = %d and p.configid = pc.configid", configid)
	rows, err := db.Query(querystring)
	checkErr(err)
	for rows.Next() {
		var (
			state1 int
			port int
		)
		_ = rows.Scan(&state1, &port)
		if port > 0 && state == state1 {
			code = 0
		}

	}
	return code
}

//改变蜜罐状态,四个状态：0：启动成功；1：启动失败；2：关闭成功；3：关闭失败
func changeState(configid, state int) (code int) {
	db, err := sql.Open("sqlite3", tools.DbPath)
	checkErr(err)
	defer db.Close()

	if state == 0 {
		state = 1
	} else {
		state = 0
	}

	sqlStr := "UPDATE pot SET state = ? WHERE configid = ?"
	stmt, err := db.Prepare(sqlStr)
	checkErr(err)

	_, err = stmt.Exec(state, configid)
	if err != nil {
		checkErr(err)
	} else {
		return 0
	}

	return 1
}

//实际蜜罐状态改变；0：成功；1：失败
func changePot(configid, state int) (code int) {
	potconfig := ConfigRead(configid)
	port := strconv.Itoa(potconfig.Port)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//mysql蜜罐
	if configid == 1004 {
		//启动操作
		if state == 0 {
			go mysql.Mysql(mysql.PotConfigs(potconfig))
		}
		if state == 1 {
			mysql.Shutdown()
		}
	}

	//goby蜜罐
	if configid == 1003 {
		if state == 0 {
			payload := strings.Replace(potconfig.Payload, "'", "\\'", -1)
			api := tools.GobyApi1 + payload + tools.GobyApi2

			tools.WriteFile(tools.GobyFile, api)
			_, err :=  tools.PathExists(tools.GobyFile)
			if err != nil {
				return 1
			}

			goby.SetInfo(potconfig.Ip, port)

			GobyPot = &http.Server{
				Addr:         ":" + port,
				Handler:      goby.Goby(),
				ReadTimeout:  2 * time.Second,
				WriteTimeout: 5 * time.Second,
			}
			GobyPot.SetKeepAlivesEnabled(false)

			G.Go(func() error {
				return GobyPot.ListenAndServe()
			})
		}
		if state == 1 {
			if err := GobyPot.Shutdown(ctx); err != nil {
				fmt.Println("goby Shutdown:", err)
				return 1
			}
		}
	}

	//vpn蜜罐
	if configid == 1002 {
		if state == 0 {
			_, err :=  tools.PathExists(tools.VpnFile)
			if err != nil {
				return 1
			}

			VpnPot = &http.Server{
				Addr:         ":" + port,
				Handler:      vpn.VPN(),
				ReadTimeout:  2 * time.Second,
				WriteTimeout: 5 * time.Second,
			}
			VpnPot.SetKeepAlivesEnabled(false)

			G.Go(func() error {
				return VpnPot.ListenAndServe()
			})
		}
		if state == 1 {
			if err := VpnPot.Shutdown(ctx); err != nil {
				fmt.Println("vpn Shutdown:", err)
				return 1
			}
		}
	}

	//burp蜜罐
	if configid == 1001 {
		if state == 0 {
			api := tools.BurpApi1 + potconfig.Payload + tools.BurpApi2

			tools.WriteFile(tools.BurpFile, api)

			_, err :=  tools.PathExists(tools.BurpFile)

			if err != nil {
				return 1
			}

			BurpPot = &http.Server{
				Addr:         ":" + port,
				Handler:      burpsuite.BurpSuite(),
				ReadTimeout:  2 * time.Second,
				WriteTimeout: 5 * time.Second,
			}
			BurpPot.SetKeepAlivesEnabled(false)

			G.Go(func() error {
				return BurpPot.ListenAndServe()
			})
		}
		if state == 1 {
			if err := BurpPot.Shutdown(ctx); err != nil {
				fmt.Println("burp Shutdown:", err)
				return 1
			}
		}
	}

	return 0
}


