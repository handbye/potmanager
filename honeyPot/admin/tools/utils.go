package tools

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"io"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

// GetSession 获取session
func GetSession(c *gin.Context) bool {
	session := sessions.Default(c)
	session.Options(sessions.Options{MaxAge: 120 * 60})
	loginuser := session.Get("secure")
	if loginuser == "admin" {
		return true
	} else {
		return false
	}
}

func IndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "nginx.html", nil)
}

// PageNotFound 404页面全部转到nginx默认页
func PageNotFound(engine *gin.Engine) {
	engine.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})
}

// SafeDate 全局过滤
func SafeDate(s string) {
	strings.TrimSpace(s)
	strings.Trim(s, "\"")
	strings.Trim(s, "'")
	strings.Trim(s, "%")
	strings.Trim(s, "#")
	strings.Trim(s, "(")
	strings.Trim(s, ")")
	strings.Trim(s, "-")
}

// In 判断元素是否在数组中
func In(target string, str_array []string) bool {
	for _, element := range str_array {
		if target == element {
			return true
		}
	}
	return false
}

//生成随机字符串用于后台地址
// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// Generate a random string of A-Z chars with len = l
func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(97, 122))
	}
	return string(bytes)
}
func RandomAdminUrl() string {
	rand.Seed(time.Now().UnixNano())
	return randomString(8)
}

// Strip 去掉字符串中空格和换行符
func Strip(old_string string) string {
	new_string := strings.Replace(old_string, " ", "", -1)
	new_string = strings.Replace(new_string, "\n", "", -1)
	return new_string
}

// ZeroToNull 0与空字符串转化
func ZeroToNull(s1 string) string {
	if s1 == "0" {
		return ""
	}
	if s1 == "" {
		return "0"
	}
	return s1
}

/*
   判断文件或文件夹是否存在
   如果返回的错误为nil,说明文件或文件夹存在
   如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
   如果返回的错误为其它类型,则不确定是否在存在
*/

func PathExists(path string) (bool, error) {

	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func WriteFile(filename, data string) {
	var (
		err error
	)
	// 拿到一个文件对象
	// file对象肯定是实现了io.Reader,is.Writer
	fileObj, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	writer := bufio.NewWriter(fileObj)
	defer writer.Flush()
	defer fileObj.Close()
	_, err = writer.WriteString(data)
	if err != nil {
		fmt.Println(err)
	}
}

// 计算密码复杂度

const (
	levelD = iota
	LevelC
	LevelB
	LevelA
	LevelS
)

func CheckPass(minLength, maxLength, minLevel int, pwd string) bool {
	if len(pwd) < minLength {
		//fmt.Printf("密码长度必须大于 %d", minLength)
		return false
	}
	if len(pwd) > maxLength {
		//fmt.Printf("密码长度必须小于 %d", maxLength)
		return false
	}

	var level int = levelD
	patternList := []string{`[0-9]+`, `[a-z]+`, `[A-Z]+`, `[~!@#$%^&*?_-]+`}
	for _, pattern := range patternList {
		match, _ := regexp.MatchString(pattern, pwd)
		if match {
			level++
		}
	}

	if level < minLevel {
		//fmt.Println("密码复杂度太低,必须包含大小写数字和字母")
		return false
	}
	return true
}

// Md5File 计算文件的hash
func Md5File(filepath string) string {
	file, err := os.Open(filepath)
	if err != nil {
		return ""
	}
	defer file.Close()

	m := md5.New()
	_, err = io.Copy(m, file)
	if err != nil {
		return ""
	}

	return hex.EncodeToString(m.Sum(nil))
}

func CreateUploadDic() {
	dicname := Cwd + string(os.PathSeparator) + "upload"
	_, err := os.Stat(dicname);
	if os.IsNotExist(err) {
		os.Mkdir(dicname, os.FileMode(0660))
		os.Chmod(dicname, os.FileMode(0660))
	}
	_, err = PathExists(dicname)
	if err != nil {
		fmt.Println("请检查当前用户是否具有创建目录的权限！")
		os.Exit(1)
	}
}