// JavaScript Document
function qiehuantap(tapnum) {
    for (var i = 1; i <= 3; i++) {
        document.getElementById("divtap" + i).style.display = "none";
        document.getElementById("limenu" + i).className = "";
    }
    document.getElementById("divtap" + tapnum).style.display = "block";
    document.getElementById("limenu" + tapnum).className = "select";
};
function jilutap(tapnum) {
    for (var i = 1; i <= 2; i++) {
        document.getElementById("xfjl" + i).style.display = "none";
        document.getElementById("tapli" + i).className = "";
    }
    document.getElementById("xfjl" + tapnum).style.display = "block";
    document.getElementById("tapli" + tapnum).className = "select";
};

//选择样式
function listnum(tapnum) {
    for (var i = 1; i <= 36; i++) {
        document.getElementById("listli" + i).className = "";
    }
    document.getElementById("listli" + tapnum).className = "listfont";
};