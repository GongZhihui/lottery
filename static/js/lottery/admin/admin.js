var host = "http://" + window.location.host;

// 添加物品
function addItem() {
    var item = $("#itemName").val();
    var percent = $("#itemPercent").val();
    if (item != "" && percent != "") {
        var tr = "<tr><td>" + item + "</td>" + "<td>" + percent + "</td>"
            + "<td><button type=\"button\" class=\"btn btn-danger\">删除</button></td></tr>";
        $("#lottoTBodyAdd").append(tr);
    }
}

function init(){
    // 清空table
    $("#lottoTBody").empty();

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

    $("#lottoTBodyAdd").on("click", "button",function(){
        $(this).parent().parent().remove();
    });
}

$(function(){
    host = "http://" + window.location.host;
    init();

    
});

// 提交item
function submitItems() {
    var items = [];
    var i = 0; 
    $("#lottoTBodyAdd").find("tr").each(function(){
        var tdArr = $(this).children();
        var name = tdArr.eq(0).text();
        var percent = parseFloat(tdArr.eq(1).text());
        
        var item = {"name" : name, "percent" : percent};
        items[i++] = item;
    });

    var jvdata = {"items": items};

    var req = JSON.stringify(jvdata);
    console.log("req: ", req);
    $.ajax({
        type: 'POST',
        url: host + "/v1/admin/lotto_item",
        data: req,
        dataType: "json",
        contentType: "application/json",
        success: function (result) {
            if(result == true){
                init();
            }
        },
        error: function (err) {
            alert("错误: " + err);
        }

    });
}

// 提交notice
function submitNoticeFunc() {
    var notice = $("#submitNotice").val()
    var jvdata = {"notice": notice};

    var req = JSON.stringify(jvdata);
    console.log("req: ", req);
    $.ajax({
        type: 'POST',
        url: host + "/v1/admin/notice",
        data: req,
        dataType: "json",
        contentType: "application/json",
        success: function (result) {
            if(result == true){
                init();
            }
        },
        error: function (err) {
            alert("错误: " + err);
        }

    });
}

// 提交prompt
function submitPromptFunc() {
    var prompt = $("#submitPrompt").val()
    var jvdata = {"lottoPrompt": prompt};

    var req = JSON.stringify(jvdata);
    console.log("req: ", req);
    $.ajax({
        type: 'POST',
        url: host + "/v1/admin/prompt",
        data: req,
        dataType: "json",
        contentType: "application/json",
        success: function (result) {
            if(result == true){
                init();
            }
        },
        error: function (err) {
            alert("错误: " + err);
        }

    });
}