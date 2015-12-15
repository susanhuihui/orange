var allcookies = document.cookie;
function getCookie(cookie_name) {
    var allcookies = document.cookie;
    var cookie_pos = allcookies.indexOf(cookie_name);   //�����ĳ���  
    // ����ҵ����������ʹ���cookie���ڣ�    
    // ��֮����˵�������ڡ�    
    if (cookie_pos != -1) {
        // ��cookie_pos����ֵ�Ŀ�ʼ��ֻҪ��ֵ��1���ɡ�    
        cookie_pos += cookie_name.length + 1;      //�������Լ��Թ������׳����⣬�������Ҳο���ʱ���Լ��ú��о�һ�¡�����    
        var cookie_end = allcookies.indexOf(";", cookie_pos);
        if (cookie_end == -1) {
            cookie_end = allcookies.length;
        }
        var value = unescape(allcookies.substring(cookie_pos, cookie_end)); //����Ϳ��Եõ�����Ҫ��cookie��ֵ�ˡ�����    
    }
    return value;
}

//���cookie  
function clearCookie(name) {
    setCookie(name, "", -1);
}

//����cookie  
function setCookie(name, value, seconds) {
    seconds = seconds || 0;   //seconds��ֵ��ֱ�Ӹ�ֵ��û��Ϊ0�������php��һ����  
    var expires = "";
    if (seconds != 0) {      //����cookie����ʱ��  
        var date = new Date();
        date.setTime(date.getTime() + (seconds * 1000));
        expires = "; expires=" + date.toGMTString();
    }
    document.cookie = name + "=" + escape(value) + expires + "; path=/";   //ת�벢��ֵ  
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



//ɾ���������˵Ŀո�
function trim(str) {
    return str.replace(/(^\s*)|(\s*$)/g, "");
}

//�����ַ�����2015-11-01T09:30:50+08:00����ȡ��ʱ��2015-5-5
function getSimpDate(time) {
    var starttime = new Date(time);
    var strstr = starttime.getFullYear() + "-" + (starttime.getMonth() + 1) + "-" + starttime.getDate();
    return strstr;
}

//�����ַ�����2015-11-01T09:30:50+08:00����ȡ��ʱ��
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

//��ʱ��ת��Ϊ����ӻ��޸ĵ�����
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

//��ȡ����ʱ�䣬�������ݿ����ʱ��ʱʹ��
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

//��ȡ����֮ǰ�ļ������ڣ�����������
function getdate(daycount) {
    //�������ڣ���ǰ���ڵ�ǰ����
    var myDate = new Date(); //��ȡ��������
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
//��ȡ���ڼ��ϵ�һ������0����ӽ��쿪ʼ1��������쿪ʼ��ȡ���ڣ��ڶ������������ȡ��������
function getdate2(startday, daycount) {
    //�������ڣ���ǰ���ڵ�ǰ����
    var myDate = new Date(); //��ȡ��������
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

//��ȡ���ڵ��º��գ����������ʽ��2015-12-12��������ʽΪ��12-12
function getmonthday(time) {
    time += " 00:00:00";
    time = time.replace(/-/g, "/");
    var myDate = new Date(time); //��ȡ��������
    var strdate = (myDate.getMonth() + 1) + "-" + myDate.getDate();
    return strdate;
}


//��ĳ��ʱ��ε�Сʱ����ȥ�̶�ֵ ���룺2015-12-12 15��00��00 ������2015-12-12 03��00��00
function setHours(date, value) {
    date.setHours(date.getHours() - value);
    return date;
}