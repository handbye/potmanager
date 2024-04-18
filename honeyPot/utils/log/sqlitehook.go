package mylog

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"honeypot/admin/tools"
	"time"
)

type SQLiteHook struct {
	db      *sql.DB
	tablename string
	timeout time.Duration
}

// NewSQLiteHook - create new SQLite3 logrus hook
func NewSQLiteHook(db *sql.DB, tablename string, timeout time.Duration) (hook *SQLiteHook, err error) {
	if err = db.Ping(); err != nil {
		return
	}
	hook = &SQLiteHook{
		db:      db,
		tablename: tablename,
		timeout: timeout,
	}
	return
}

func (hook *SQLiteHook) Fire(entry *logrus.Entry) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), hook.timeout)
	defer cancel()

	str, err := entry.String()
	if err != nil {
		err = errors.Wrap(err, "unable to read logrus entry")
		return
	}
	t := entry.Time.Format("2006-01-02 15:04:05")
	var rowLen int
	hook.db.QueryRow(fmt.Sprintf("SELECT 1 FROM %s WHERE time = '%s' and clientIP = '%s'",hook.tablename,t,entry.Data["clientIP"])).Scan(&rowLen)

	//插入语句
	if rowLen == 0{
		if tools.In(hook.tablename, tools.Config("httplog")){
			query := fmt.Sprintf("INSERT INTO %s(time,clientIP,statusCode,reqMethod,reqUri,full_message) VALUES (?,?,?,?,?,?)",hook.tablename)
			smt, err := hook.db.PrepareContext(ctx, query)
			if err != nil {
				err = errors.Wrap(err, "unable to insert log entry")
			} else {
				smt.ExecContext(ctx,
					t,
					entry.Data["clientIP"],
					entry.Data["statusCode"],
					entry.Data["reqMethod"],
					entry.Data["reqUri"],
					str,
				)
			}
		}
		if tools.In(hook.tablename, tools.Config("nohttplog")){
			query := fmt.Sprintf("INSERT INTO %s(time,msg) VALUES (?,?)",hook.tablename)
			smt, err := hook.db.PrepareContext(ctx, query)
			if err != nil {
				err = errors.Wrap(err, "unable to insert log entry")
			} else {
				smt.ExecContext(ctx,
					t,
					str,
				)
			}
		}
	}
	return
}

func (hook *SQLiteHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
