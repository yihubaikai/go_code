package main

import(
    "net/http"
    "fmt"
    "path/filepath"
    "strings"
    "os"
    "time"
    "io"
    "net"
    "errors"

)

/*保存文件（优化版）*/
func SaveLog(m_FilePath string, val string) {
	var dir, filename string
	filename = filepath.Base(m_FilePath)
	if len(m_FilePath) > 1 && string([]byte(m_FilePath)[1:2]) == ":" {
		filename = filepath.Base(m_FilePath)
		dir = strings.TrimSuffix(m_FilePath, filename)
		//fmt.Println("abspath:filename:" + filename + "\n" + "dir:" + dir + "\n")
	} else {
		dir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
		dir = dir + "/" + m_FilePath
		filename = filepath.Base(m_FilePath)
		dir = strings.TrimSuffix(dir, filename)
		//fmt.Println("noptabspath:filename:" + filename + "\n" + "dir:" + dir + "\n")
	}

	p := dir + "/" + filename
	p = strings.Replace(p, "\\", "/", -1)
	p = strings.Replace(p, "//", "/", -1)
	//fmt.Println("fullpath:" + p + "\n")

	//fmt.Println(p)
	_, err := os.Stat(dir)
	if err != nil {
		if !os.IsExist(err) {
			os.MkdirAll(dir, os.ModePerm)
		}
	}
	fl, err := os.OpenFile(p, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	defer fl.Close()

	if err == nil {
		io.WriteString(fl, val)
	}
}


func ReadLog(m_FilePath string, iSeek int64) string{
	var dir, filename string
	filename = filepath.Base(m_FilePath)
	if len(m_FilePath) > 1 && string([]byte(m_FilePath)[1:2]) == ":" {
		filename = filepath.Base(m_FilePath)
		dir = strings.TrimSuffix(m_FilePath, filename)
		//fmt.Println("abspath:filename:" + filename + "\n" + "dir:" + dir + "\n")
	} else {
		dir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
		dir = dir + "/" + m_FilePath
		filename = filepath.Base(m_FilePath)
		dir = strings.TrimSuffix(dir, filename)
		//fmt.Println("noptabspath:filename:" + filename + "\n" + "dir:" + dir + "\n")
	}

	p := dir + "/" + filename
	p = strings.Replace(p, "\\", "/", -1)
	p = strings.Replace(p, "//", "/", -1)
	//fmt.Println("fullpath:" + p + "\n")
	
	_, err := os.Stat(dir)
	if err != nil {
		return ""
	}


	//fmt.Println(p)
	fl, err := os.OpenFile(p, os.O_RDWR, os.ModePerm)
	defer fl.Close()
	if(err!=nil){
		fmt.Println("File is Not Found")
		return ""
	}

	//读取最后1K
	iStart   := int64(0)   //读取文本的开始位置
	stat,err := os.Stat(p) //读取文本大小
	iSize    := stat.Size()//申请内存空间大小

	if(stat.Size()  > iSeek){
		iStart = stat.Size() - iSeek
		iSize  = iSeek 
	}

	sLen,err1 := fl.Seek(iStart, io.SeekStart)
	t := "SUCC"
	if(err1 != nil){
		t = "Seek Error!"
	}
    fmt.Println("读取文件开始位置:", iStart, "申请内存大小:", iSize, "实际文件大小:", stat.Size(), "Seek文件大小:", sLen, t)
    buf := make([]byte, iSize)
    fl.Read(buf)
    ret := string(buf)//fmt.Sprintf("%s",string(buf))
	//fmt.Println(ret)
	return ret;


	/* 全部读取
	ret := ""
	scanner := bufio.Newscanner(fl)
	for scanner.Scan(){
		ret = ret + scanner.Text()
	}
	return ret
	*/
}

/*获取当前时间*/
func Gettime() string {
	Year := time.Now().Year()     //年[:3]
	Month := time.Now().Month()   //月
	Day := time.Now().Day()       //日
	Hour := time.Now().Hour()     //小时
	Minute := time.Now().Minute() //分钟
	Second := time.Now().Second() //秒
	//Nanosecond:=time.Now().Nanosecond()//纳秒
	var timestr string
	timestr = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", Year, Month, Day, Hour, Minute, Second)
	return timestr
}

func GetIP(r *http.Request) (string, error) {
	ip := r.Header.Get("X-Real-IP")
	if net.ParseIP(ip) != nil {
		return ip, nil
	}

	ip = r.Header.Get("X-Forward-For")
	for _, i := range strings.Split(ip, ",") {
		if net.ParseIP(i) != nil {
			return i, nil
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	if net.ParseIP(ip) != nil {
		return ip, nil
	}
	
	return "", errors.New("no valid ip found")
}


func main() {

    http.HandleFunc("/zc",hello)
    http.HandleFunc("/",Home)
    http.HandleFunc("/show",Show)
    http.ListenAndServe(":8080",nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w,"Hello Docker Form Golang!")
}


func Home(w http.ResponseWriter, r *http.Request) {
    ip, err := GetIP(r)
    txt := fmt.Sprintf("[%v] : %v SUCC!\n", Gettime(), ip)
    if(err != nil){
		txt = fmt.Sprintf("[%v] : %v Error:%\n", Gettime(), ip, err)
    }    
    SaveLog("ip.txt", txt)
    fmt.Println("Home:", txt)
    fmt.Fprintf(w,txt)
}
func Show(w http.ResponseWriter, r *http.Request) {
	
    txt := ReadLog("ip.txt", 400)
    fmt.Println("Show:", txt)
    txt = "<pre>\n" + txt + "\n</pre>" 
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprintf(w,txt)
}


/*
执行指令:
docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)
docker ps

docker stop $(docker ps -a -q)
docker rmi $(docker images -q)
docker images

docker build -t zcdocker .
docker run -p 8080:8080 -d zcdocker


Dockerfile 文件   211M

FROM centos:7
ADD ./http /http
EXPOSE 8080
CMD ["/http", "8080"]

FROM gcr.io/distroless/static-debian11	7M
FROM gcr.io/distroless/base-debian11

centos:7 : 211
JAVA   : 217
GOLANG : 972

*/
