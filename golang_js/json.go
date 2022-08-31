package main

import (
    "encoding/json"
    "fmt"
    "github.com/robertkrimen/otto"
    "github.com/yihubaikai/gopublic/net"
)
 //go get -v -u github.com/robertkrimen/otto
 //go get -v -u github.com/yihubaikai/gopublic
/*
{
    "state":"1",
    "msg":"",
    "title":"",
    "content":"",
    "auth":"",
    "url":"",
    "imgurl":"",
    "time":"",
    "remark":""
}
*/

type NewsItem struct{
    State   string   `"state"`
    Msg     string   `"msg"`
    Title   string   `"title"`
    Content string   `"content"`
    Auth    string   `"auth"`
    Url     string   `"url"`
    ImgUrl  string   `"imgurl"`
    Time    string   `"time"`
    Remark  string   `"remark"`
}
var iVm = 0;
var vm = otto.New()

func GetNews(Text string) map[string]string{

jscode := `//生成从minNum到maxNum的随机数
function GetRand(minNum,maxNum){
    switch(arguments.length){ 
        case 1: 
            return parseInt(Math.random()*minNum+1,10); 
        break; 
        case 2: 
            return parseInt(Math.random()*(maxNum-minNum+1)+minNum,10); 
        break; 
            default: 
                return 0; 
            break; 
    } 
}

function GetRandomNum(Min,Max)
{
  return (new Date()).getTime()%Max;
} 


function GetNews(jsontext){  
  var aRet = {"state":"1", "msg":"", "title":"", "content":"","auth":"", "url":"", "imgurl":"","time":"", "remark":""};
  var arr, iCount=0;
  try {
      arr = JSON.parse(jsontext);  
    } catch (e) {
        aRet["title"]   = "1.数据错误，请联系管理员";
        aRet["content"] = aRet["title"];
        aRet["auth"]    = aRet["title"];
        aRet["imgurl"]  = aRet["title"];
        return JSON.stringify(aRet); 
    }

    try {
        iCount = arr["data"]["data"].length;
    } catch (e) {
        aRet["msg"] = "2.获取失败，请刷新!";
        aRet["title"]   = aRet["msg"] ;
        aRet["content"] = aRet["title"];
        aRet["auth"]    = aRet["title"];
        aRet["imgurl"]  = aRet["title"];
        return JSON.stringify(aRet); 
    }
    if(iCount == 0){
      aRet["msg"] = "3.无新闻内容";
      aRet["title"]   = aRet["msg"] ;
      aRet["content"] = aRet["title"];
      aRet["auth"]    = aRet["title"];
      aRet["imgurl"]  = aRet["title"];
      return JSON.stringify(aRet);
    }


 try {
    var iRnd = GetRandomNum(0, iCount);
    var tArr = arr["data"]["data"][iRnd];
    if("title" in tArr){
      aRet["title"] = tArr["title"];
    }
    if("url" in tArr){
      aRet["url"] = tArr["url"];
    }
    if("imgurl_https" in tArr){
      aRet["imgurl"] = tArr["imgurl_https"];
    }
    if("intro" in tArr){
      aRet["content"] = iRnd;//tArr["intro"];
    }
    if("source_from" in tArr){
      aRet["auth"] = tArr["source_from"];
    }
    if("published_at" in tArr){
      aRet["time"] = tArr["published_at"];
    }
    aRet["state"] = "0";
    aRet["msg"] = "Rand:" + iRnd;
    } catch (e) {
        aRet["msg"] = "4.键值改变,请刷新";
        aRet["title"]   = aRet["msg"] ;
        aRet["content"] = aRet["title"];
        aRet["auth"]    = aRet["title"];
        aRet["imgurl"]  = aRet["title"];
        return JSON.stringify(aRet); 
    }
   return JSON.stringify(aRet); 
}`


sRet := make(map[string]string)
if(iVm == 0){
    //创建虚拟机
    //vm := otto.New()

    //执行虚拟机
    _, err := vm.Run(jscode) //value, err := vm.Call("encodeInp2", nil, data) 
    if err!=nil {
        sRet["state"] = "1"
        sRet["msg"] = "1.执行JS代码失败"
        return sRet
    }
    iVm = 1
}

if(iVm == 0){
    sRet["state"] = "1"
    sRet["msg"] = "2.执行JS代码失败"
    return sRet
}


//解析传入数据
 value, err := vm.Call("GetNews", nil, Text)
 if err != nil {
    sRet["state"] = "1"
    sRet["msg"] = "3.执行JS代码失败"
    return sRet
}

    sRet["state"] = "0"
    sRet["msg"] = "SUCC"
    sRet["text"] = value.String()
    return sRet 
}

func Get_News_Item(){
       //获取json字符串
    s := make(map[string]string)
    b := hNet.Httpget("https://api2.firefoxchina.cn/home/news_cnxh.json?v=20211221155704", s)

    s = GetNews(b)
    fmt.Println(s["text"])
    if(s["state"] == "0"){
        data := NewsItem{}
        if err := json.Unmarshal([]byte(s["text"]), &data); err != nil {
            fmt.Println(err)
        }
        fmt.Println( data.Title, data.Content, data.Auth, data.Url, data.ImgUrl);
    }
}

func main() {  
    Get_News_Item()
}
