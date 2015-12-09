// JavaScript Document
//<!--教师列表轮换js-->
function setDetailMsgRowText(rowID, btn) {
    var row = document.getElementById(rowID);
    for (var i = 0; i < 4; i++)//隐藏所有行
    {
        var rownone = document.getElementById("row" + (i + 1));
        rownone.style.display = "none";
        document.getElementById("dian" + (i + 1)).src = "images/leve.png";
    }
    //if (row.style.display == (document.all ? "none" : "table-row")) {
    row.style.display = "block";
    //}
    //else {
    //  row.style.display = (document.all ? "none" : "table-row");
    //}
    document.getElementById(btn).src = "images/now.png";
};

//<!--弹出层js--> 
function showDiv() {
    document.getElementById('popDiv').style.display = 'block';
    document.getElementById('bg').style.display = 'block';
};

function closeDiv() {
    document.getElementById('popDiv').style.display = 'none';
    document.getElementById('bg').style.display = 'none';
};

function showzcDiv() {
    document.getElementById('popzcDiv').style.display = 'block';
    document.getElementById('bg').style.display = 'block';
};

//<!--登录选项卡js-->
function closezcDiv() {
    document.getElementById('popzcDiv').style.display = 'none';
    document.getElementById('bg').style.display = 'none';
};

function selectTag(showContent, selfObj) {
    // 操作标签
    var tag = document.getElementById("tags").getElementsByTagName("li");
    var taglength = tag.length;
    for (i = 0; i < taglength; i++) {
        tag[i].className = "";
    }
    selfObj.parentNode.className = "selectTag";
    // 操作内容
    for (i = 0; j = document.getElementById("tagContent" + i) ; i++) {
        j.style.display = "none";
    }
    document.getElementById(showContent).style.display = "block";
};

function setTitleOnSelect(titleindex) {
    if (titleindex == 1) { //选中首页
        document.getElementById("shouye").className = "menuasty";
        document.getElementById("fontshouye").style.color = "#FF9A3A";
    } else if (titleindex == 2) {//选中老师
        document.getElementById("laoshi").className = "menuasty";
        document.getElementById("fontlaoshi").style.color = "#FF9A3A";
    } else if (titleindex == 3) {//选中问答中心
        document.getElementById("wenda").className = "menuasty";
        document.getElementById("fontwenda").style.color = "#FF9A3A";
    } else if (titleindex == 4) {//选中注册/用户姓名
        document.getElementById("lizhuce").className = "menuasty";
        document.getElementById("lizhuce").children[0].style.color = "#FF9A3A";
    } else if (titleindex == 0) {//什么都不选
        //document.getElementById("shouye").className = "";
        //document.getElementById("laoshi").className = "";
        //document.getElementById("wenda").className = "";
        //document.getElementById("lizhuce").className = "";
        //document.getElementById("fontshouye").style.color = "#9d9d9d";
        //document.getElementById("fontlaoshi").style.color = "#9d9d9d";
        //document.getElementById("fontwenda").style.color = "#9d9d9d";
        //document.getElementById("lizhuce").style.color = "#9d9d9d";
    }
    getUserInfo();
}

function getUserInfo() {
    var userid = getCookie("userid");
    if (userid > 0) {
        var website = getCookie("OnlineUrl");
        var getstudent = "http://" + website + "/orange/userinformation/GetUserinformationById/" + userid;
        $.getJSON(getstudent, function (data) {
            if (data != null && data["Id"] > 0) {
                document.getElementById("loginusername").innerHTML = data["UserName"];
                return data["UserName"];
            }
        });
    }
}


//删除左右两端的空格
function trim(str) {
    return str.replace(/(^\s*)|(\s*$)/g, "");
}

//获取当前系统时间
function getNowFormatDate() {
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
    var currentdate = date.getFullYear() + seperator1 + month + seperator1 + strDate
            + "T" + date.getHours() + seperator2 + date.getMinutes()
            + seperator2 + date.getSeconds() + "" + "Z";
    return currentdate;
}

function selectTagpass(showContent, selfObj) {
    document.getElementById("li1").className = "";
    document.getElementById("li2").className = "";
    selfObj.parentNode.className = "selectTag";
    if (showContent == "tagContent0") {
        document.getElementById("divcoutent").innerHTML = "";
        var tupian = '<div class="tagContents" id="tagContent1">' +
                                '<img style="width: 90%;" src="images/stuprocess.png" />' +
                            '</div>';
        document.getElementById("divcoutent").innerHTML = tupian;
    } else if (showContent = "tagContent1") {
        document.getElementById("divcoutent").innerHTML = "";
        var tupian = '<div class="tagContents" id="tagContent1">'+
                                '<img style="width: 100%;" src="images/teacherprocess.png" />'+
                            '</div>';
        document.getElementById("divcoutent").innerHTML = tupian;
    }
};

//密码验证方法
function isPasswd(s) {
    var patrn = /^(\w){6,20}$/;
    if (!patrn.exec(s)) return false
    return true
}