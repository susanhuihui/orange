var allcookies = document.cookie;
function getCookie(cookie_name) {
    var allcookies = document.cookie;
    var cookie_pos = allcookies.indexOf(cookie_name);   //索引的长度  
    // 如果找到了索引，就代表cookie存在，    
    // 反之，就说明不存在。    
    if (cookie_pos != -1) {
        // 把cookie_pos放在值的开始，只要给值加1即可。    
        cookie_pos += cookie_name.length + 1;      //这里我自己试过，容易出问题，所以请大家参考的时候自己好好研究一下。。。    
        var cookie_end = allcookies.indexOf(";", cookie_pos);
        if (cookie_end == -1) {
            cookie_end = allcookies.length;
        }
        var value = unescape(allcookies.substring(cookie_pos, cookie_end)); //这里就可以得到你想要的cookie的值了。。。    
    }
    return value;
}

//清除cookie  
function clearCookie(name) {
    setCookie(name, "", -1);
}

//设置cookie  
function setCookie(name, value, seconds) {
    seconds = seconds || 0;   //seconds有值就直接赋值，没有为0，这个根php不一样。  
    var expires = "";
    if (seconds != 0) {      //设置cookie生存时间  
        var date = new Date();
        date.setTime(date.getTime() + (seconds * 1000));
        expires = "; expires=" + date.toGMTString();
    }
    document.cookie = name + "=" + escape(value) + expires + "; path=/";   //转码并赋值  
}

function getCookie2(name) {
    var cookies = document.cookie.split(";");
    for (var i = 0; i < cookies.length; i++) {
        var cookie = cookies[i];
        var cookieStr = cookie.split("=");
        if (cookieStr && cookieStr[0].trim() == name) {
            return decodeURI(cookieStr[1]);
        }
    }
}

String.prototype.trim = function () {
    return this.replace(/^(\s*)|(\s*)$/g, "");
}



//删除左右两端的空格
function trim(str) {
    return str.replace(/(^\s*)|(\s*$)/g, "");
}

//根据字符串（2015-11-01T09:30:50+08:00）获取简单时间2015-5-5
function getSimpDate(time) {
    var starttime = new Date(time);
    var strstr = starttime.getFullYear() + "-" + (starttime.getMonth() + 1) + "-" + starttime.getDate();
    return strstr;
}

//根据字符串（2015-11-01T09:30:50+08:00）获取简单时分
function getSimpTime(time) {
    var starttime = new Date(time);
    var hours = starttime.getHours();
    if (hours >= 0 && hours <= 9) {
        hours = "0" + hours;
    }
    var minutes = starttime.getMinutes();
    if (minutes >= 0 && minutes <= 9) {
        minutes = "0" + minutes;
    }
    var stardate = hours + ":" + minutes;
    return stardate;
}

//将时间转换为可添加或修改的样子
function getInsertDate(time) {
    var date = new Date(time);
    var seperator1 = "-";
    var seperator2 = ":";
    var month = date.getMonth() + 1;
    var strDate = date.getDate();
    if (month >= 1 && month <= 9) {
        month = "0" + month;
    }
    if (strDate >= 0 && strDate <= 9) {
        strDate = "0" + strDate;
    }
    var hours = date.getHours();
    if (hours >= 0 && hours <= 9) {
        hours = "0" + hours;
    }
    var minutes = date.getMinutes();
    if (minutes >= 0 && minutes <= 9) {
        minutes = "0" + minutes;
    }
    var second = date.getSeconds();
    if (second >= 0 && second <= 9) {
        second = "0" + second;
    }
    var currentdate = date.getFullYear() + seperator1 + month + seperator1 + strDate
            + "T" + hours + seperator2 + minutes
            + seperator2 + second + "" + "Z";
    return currentdate;
}

//获取现在时间，用于数据库添加时间时使用
function getInsertNowDate() {
    var date = new Date();
    var seperator1 = "-";
    var seperator2 = ":";
    var month = date.getMonth() + 1;
    var strDate = date.getDate();
    if (month >= 1 && month <= 9) {
        month = "0" + month;
    }
    if (strDate >= 0 && strDate <= 9) {
        strDate = "0" + strDate;
    }
    var hours = date.getHours();
    if (hours >= 0 && hours <= 9) {
        hours = "0" + hours;
    }
    var minutes = date.getMinutes();
    if (minutes >= 0 && minutes <= 9) {
        minutes = "0" + minutes;
    }
    var second = date.getSeconds();
    if (second >= 0 && second <= 9) {
        second = "0" + second;
    }
    var currentdate = date.getFullYear() + seperator1 + month + seperator1 + strDate
            + "T" + hours + seperator2 + minutes
            + seperator2 + second + "" + "Z";
    return currentdate;
}

//获取今天之前的几个日期，不包括今天
function getdate(daycount) {
    //设置日期，当前日期的前七天
    var myDate = new Date(); //获取今天日期
    myDate.setDate(myDate.getDate() - daycount);
    var dateArray = [];
    var dateTemp;
    var flag = 1;
    for (var i = 0; i < daycount; i++) {
        dateTemp = myDate.getFullYear() + "-" + (myDate.getMonth() + 1) + "-" + myDate.getDate();
        dateArray.push(dateTemp);
        myDate.setDate(myDate.getDate() + flag);
    }
    return dateArray;
}
//获取日期集合第一个参数0代表从今天开始1代表从明天开始获取日期，第二个参数往后获取几个日期
function getdate2(startday, daycount) {
    //设置日期，当前日期的前七天
    var myDate = new Date(); //获取今天日期
    myDate.setDate(myDate.getDate() + startday);
    var dateArray = [];
    var dateTemp;
    var flag = 1;
    for (var i = 0; i < daycount; i++) {
        dateTemp = myDate.getFullYear() + "-" + (myDate.getMonth() + 1) + "-" + myDate.getDate();
        dateArray.push(dateTemp);
        myDate.setDate(myDate.getDate() + flag);
    }
    return dateArray;
}

//获取日期的月和日，传入参数格式：2015-12-12，传出格式为：12-12
function getmonthday(time) {
    time += " 00:00:00";
    time = time.replace(/-/g, "/");
    var myDate = new Date(time); //获取今天日期
    var strdate = (myDate.getMonth() + 1) + "-" + myDate.getDate();
    return strdate;
}


//将某个时间段的小时，减去固定值 传入：2015-12-12 15：00：00 传出：2015-12-12 03：00：00
function setHours(date, value) {
    date.setHours(date.getHours() - value);
    return date;
}