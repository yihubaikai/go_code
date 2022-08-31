//生成从minNum到maxNum的随机数
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

var gARR;
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
}