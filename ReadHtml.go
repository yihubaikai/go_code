package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

//me.exe https://www.zusms.com/messages/601fe5b6b9ff681aa1bb3fb6 科技
//不停访问这个页面， 当出现科技2个关键字的时候就退出+保存
//********************************************************
//get请求
//调用demo
//	s := make(map[string]string)
//	s["wd"] = "牛魔王之红孩儿诞生"
//	s["act"] = "我是get请求"
//	r := httpget("http://pay.ggpaygg.com/debug/recvtest.php", s)
//	fmt.Println(r)

func Httpgetz(desurl string, para_kv map[string]string) string {
	fullurl := desurl
	if para_kv != nil {
		u := url.Values{}
		for k, v := range para_kv {
			//fmt.Println(k, "*", v)
			u.Set(k, v)
		}
		fullurl = desurl + "?" + u.Encode()
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(fullurl)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close() //not ok
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(body)
}

/*保存文件（优化版）*/
func SaveLog(m_FilePath string, val string) {
	var dir, filename string
	filename = filepath.Base(m_FilePath)
	if len(m_FilePath) > 1 && string([]byte(m_FilePath)[1:2]) == ":" {
		filename = filepath.Base(m_FilePath)
		dir = strings.TrimSuffix(m_FilePath, filename)
		fmt.Println("abspath:filename:" + filename + "\n" + "dir:" + dir + "\n")
	} else {
		dir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
		dir = dir + "/" + m_FilePath
		filename = filepath.Base(m_FilePath)
		dir = strings.TrimSuffix(dir, filename)
		fmt.Println("noptabspath:filename:" + filename + "\n" + "dir:" + dir + "\n")
	}

	p := dir + "/" + filename
	p = strings.Replace(p, "\\", "/", -1)
	p = strings.Replace(p, "//", "/", -1)
	fmt.Println("fullpath:" + p + "\n")

	_, err := os.Stat(dir)
	if err != nil {
		if !os.IsExist(err) {
			os.MkdirAll(dir, os.ModePerm)
		}
	}
	fl, err := os.OpenFile(p, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	defer fl.Close()

	if err != nil {
		fmt.Println("SaveLog:error")
	} else {
		fmt.Println("SaveLog:SUCC")
		io.WriteString(fl, val)
	}
}
func main() {
	if len(os.Args) != 3 {
		fmt.Println("使用方法: \nme.exe https://www.zusms.com/messages/601fe5b6b9ff681aa1bb3fb6 科技")
		return
	}

	for {
		//r := Httpgetz("https://www.zusms.com/messages/600fb6ccd5f46474bb8b4d14", nil) //16517528229
		//r := Httpgetz("https://www.zusms.com/messages/5fc5178f598e4e6768abb92b", nil) //16535533188
		//r := Httpgetz("https://www.zusms.com/messages/601fe5b6b9ff681aa1bb3fb6", nil) //16521555240
		r := Httpgetz(os.Args[1], nil)
		//fmt.Println("https://www.zusms.com/messages/601fe5b6b9ff681aa1bb3fb6 : 16521555240")
		ifind := strings.Index(r, os.Args[2]) //经略
		if ifind > 10 {
			txt := r[ifind-9 : ifind+200]
			fmt.Println(txt)
			SaveLog("log.txt", txt+"\n")
			break
		} else {
			fmt.Println(".")
		}
	}
}
