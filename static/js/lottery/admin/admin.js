var host = "http://" + window.location.host;

// 添加物品
function addItem() {
    var item = $("#item").val();
    var percent = $("#percent").val();
    if (item != "" && percent != "") {
        var tr = "<tr><td>" + item + "</td>" + "<td>" + percent + "</td>"
            + "<td><button type=\"button\" class=\"btn btn-danger\">删除</button></td></tr>";
        $("#tbody").append(tr);
    }
}

function init(){
    var req = "";
    $.ajax({
        type: 'GET',
        url: host + "/v1/admin/",
        data: "",
        cache: false,
        contentType: "application/json",
        success: function (result) {
            var prompt = result.lottoPrompt;
            var notice = result.notice;
            $("#prompt").val(prompt);
            $("#notice").val(notice);

            var items = result.items;

            var styleTr = ['class="success"', 'class="warning"', 'class="danger"'];

            for(var i = 0,len = items.length; i < len; i++){
                var tr = "<tr " + styleTr[i%3] + "><td>" + items[i].name + "</td>" + "<td>" + items[i].percent + "</td>"
                    + "</tr>";
                $("#lottoTBody").append(tr);
            }
        },
        error: function (err) {
            alert("err: " + err);
        }

    });
}

$(function(){
    host = "http://" + window.location.host;
    init();

    $("#tbody").on("click", "button",function(){
        $(this).parent().parent().remove();
    });
});

// 提交
function submitItems() {
    var items = [];
    var i = 0; 
    $("#tbody").find("tr").each(function(){
        var tdArr = $(this).children();
        var name = tdArr.eq(0).text();
        var percent = parseFloat(tdArr.eq(1).text());
        
        var item = {"name" : name, "percent" : percent};
        items[i++] = item;
    });

    var jvdata = { "param": { "type": "login", "items": items}};

    var req = JSON.stringify(jvdata);
    console.log("req: ", req);
    $.ajax({
        type: 'POST',
        url: host + "/lottery/admin/submit",
        //url: "http://localhost:8080/lottery/admin/submit",
        data: req,
        dataType: "json",
        contentType: "application/json",
        success: function (result) {
            if(result == 200){
                alert("提交成功!");
            }else if(result == 100){
                window.location.href = host + "/lottery/admin";
                //window.location.href = "http://localhost:8080/lottery/admin";
            }else{
                alert(result);
            }
        },
        error: function (err) {
            alert("错误: " + err);
        }

    });
}

function login(){
    var account = $("#account").val();
    var pwd = $("#password").val();
    var req = {"account" : account, "password" : pwd};
    $.ajax({
        type: 'POST',
        //url: "lottery/admin/login",
        url: host + "/lottery/admin/login",
        data: JSON.stringify(req),
        dataType: "json",
        contentType: "application/json",
        success: function (result) {
            if(result == 200){
                // window.location.href = "http://localhost:8080/lottery/admin/main";
                window.location.href = host + "/lottery/admin/main";
            }else{
                alert("账号或密码错误!");
            }
        },
        error: function (err) {
            alert("登录失败!");
        }

    });
}
