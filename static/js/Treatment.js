//Ajax提交
function AjaxPost(Url,JsonData,LodingFun,ReturnFun) {
    $.ajax({
        type: "post",
        url: Url,
        data: JsonData,
        dataType: 'json',
        async: 'false',
        beforeSend: LodingFun,
        error: function () { AjaxErro({ "Status": "Erro", "Erro": "500" }); },
        success: ReturnFun
    });
}
//示例
//AjaxPost("ajax调用路徑", ajax傳参,
//                function () {
//                     //ajax加載中
//                },
//                function (data) {
//                    //ajax返回 
//                    //AjaxErro(data);
//                })


//弹出
function ErroAlert(e) {
    var index = layer.alert(e, { icon: 5, time: 2000, offset: 't', closeBtn: 0, title: '錯誤信息', btn: [], anim: 2, shade: 0 });
    layer.style(index, {
        color: '#777'
    }); 
}

//Ajax 錯誤返回处理
function AjaxErro(e) {
    if (e.Status == "Erro") {
        switch (e.Erro) {
            case "500":
                top.location.href = '/Erro/Erro500';
                break;
            case "100001":
                ErroAlert("錯誤 : 錯誤代碼 '10001'");
                break;
            default:
                ErroAlert(e.Erro);
        }
    } else {
        layer.msg("未知錯誤！");
    }
}


//生成验证碼
var code = "";
function createCode(e) {
    code = "";
    var codeLength = 4;
    var selectChar = new Array(1, 2, 3, 4, 5, 6, 7, 8, 9, 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'J', 'K', 'L', 'M', 'N', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z');
    for (var i = 0; i < codeLength; i++) {
        var charIndex = Math.floor(Math.random() * 60);
        code += selectChar[charIndex];
    }
    if (code.length != codeLength) {
        createCode(e);
    }
	if(canGetCookie == 1){
    	setCookie(e, code, 60 * 60 * 60, '/');
	}else{
		return code;
	}
}


//hours為空字符串時,cookie的生存期至浏览器会话结束。  
//hours為數字0時,建立的是一個失效的cookie,  
//這個cookie会覆盖已经建立過的同名、同path的cookie（如果這個cookie存在）。     
function setCookie(name, value, hours, path) {
    var name = escape(name);
    var value = escape(value);
    var expires = new Date();
    expires.setTime(expires.getTime() + hours * 3600000);
    path = path == "" ? "" : ";path=" + path;
    _expires = (typeof hours) == "string" ? "" : ";expires=" + expires.toUTCString();
    document.cookie = name + "=" + value + _expires + path;
}
//cookie名获取值  
function getCookieValue(name) {
    var name = escape(name);
    //读cookie属性，這將返回文檔的所有cookie     
    var allcookies = document.cookie;
    //查找名為name的cookie的开始位置     
    name += "=";
    var pos = allcookies.indexOf(name);
    //如果找到了具有该名字的cookie，那么提取並使用它的值     
    if (pos != -1) {    //如果pos值為-1则说明搜索"version="失败     
        var start = pos + name.length;   //cookie值开始的位置     
        var end = allcookies.indexOf(";", start); //从cookie值开始的位置起搜索第一個";"的位置,即cookie值结尾的位置     
        if (end == -1) end = allcookies.length; //如果end值為-1说明cookie列表里只有一個cookie     
        var value = allcookies.substring(start, end);  //提取cookie的值     
        return unescape(value);       //对它解碼           
    }
    else return "-1";    //搜索失败，返回-1  
}    