var host = "http://" + window.location.host;

var FLUSH_HTML = "可能session过期，请试着刷新页面"

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

function init() {
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

            for (var i = 0, len = items.length; i < len; i++) {
                var tr = "<tr " + styleTr[i % 3] + "><td>" + items[i].name + "</td>" + "<td>" + items[i].percent + "</td>"
                    + "</tr>";
                $("#lottoTBody").append(tr);
            }
        },
        error: function (err) {
            alert("err: " + err);
        }

    });

    $("#lottoTBodyAdd").on("click", "button", function () {
        $(this).parent().parent().remove();
    });
}

$(function () {
    host = "http://" + window.location.host;
    init();

});

// 提交item
function submitItems() {
    var items = [];
    var i = 0;
    $("#lottoTBodyAdd").find("tr").each(function () {
        var tdArr = $(this).children();
        var name = tdArr.eq(0).text();
        var percent = parseFloat(tdArr.eq(1).text());

        var item = { "name": name, "percent": percent };
        items[i++] = item;
    });

    var jvdata = { "items": items };

    var req = JSON.stringify(jvdata);
    console.log("req: ", req);
    $.ajax({
        type: 'POST',
        url: host + "/v1/admin/lotto_item",
        data: req,
        dataType: "json",
        contentType: "application/json",
        success: function (result) {
            if (result == true) {
                init();
            }
        },
        error: function (err) {
            alert("错误: " + err + FLUSH_HTML);
        }

    });
}

// 提交notice
function submitNoticeFunc() {
    var notice = $("#submitNotice").val()
    var jvdata = { "notice": notice };

    var req = JSON.stringify(jvdata);
    console.log("req: ", req);
    $.ajax({
        type: 'POST',
        url: host + "/v1/admin/notice",
        data: req,
        dataType: "json",
        contentType: "application/json",
        success: function (result) {
            if (result == true) {
                init();
            }
        },
        error: function (err) {
            alert("错误: " + err + FLUSH_HTML);
        }

    });
}

// 提交prompt
function submitPromptFunc() {
    var prompt = $("#submitPrompt").val()
    var jvdata = { "lottoPrompt": prompt };

    var req = JSON.stringify(jvdata);
    $.ajax({
        type: 'POST',
        url: host + "/v1/admin/prompt",
        data: req,
        dataType: "json",
        contentType: "application/json",
        success: function (result) {
            if (result == true) {
                init();
            }
        },
        error: function (err) {
            alert("错误: " + err + FLUSH_HTML);
        }

    });
}

function lookLotInfo() {

    var identityNum = $("#identityNum").val()
    var jvdata = { "prefix": identityNum };

    var req = JSON.stringify(jvdata);
    $.ajax({
        type: 'POST',
        url: host + "/v1/admin/lotinfo",
        data: req,
        dataType: "json",
        contentType: "application/json",
        success: function (result) {
            if (result == "") {
                alert("未找到!!!");
            } else {
                var confirm = (result.Confirm == 0) ? "未确认" : "已确认";
                var time = timestampToTime(result.CreateTime);
                var td = $("#lotInfoBody").find("tr").eq(0).find("td");
                td.eq(0).attr("lotid", result.ID);
                td.eq(0).text(result.Name);
                td.eq(1).text(result.Prefix);
                td.eq(2).text(time);
                $("#confirm").val(confirm);
                $("#Memo").val(result.Memo);
            }
        },
        error: function (err) {
            alert("错误: " + err + FLUSH_HTML);
        }

    });
}

function submitCheck() {
    var confirm = parseInt($("#confirm").val());
    if(confirm != 1 &&  confirm != 0){
        alert("请输入0 或者 1");
        return;
    }
    var td = $("#lotInfoBody").find("tr").eq(0).find("td");
    var id = td.eq(0).attr("lotid");
    var jvdata = {
        "ID" : id,
        "Confirm": confirm,
        "Memo" : $("#Memo").val()
    }

    var req = JSON.stringify(jvdata);
    $.ajax({
        type: 'POST',
        url: host + "/v1/admin/lotinfo_update",
        data: req,
        dataType: "json",
        contentType: "application/json",
        success: function (result) {
            alert("提交成功");
        },
        error: function (err) {
            alert("错误: " + err);
        }

    });
}

function timestampToTime(timestamp) {
    var date = new Date(timestamp * 1000);//时间戳为10位需*1000，时间戳为13位的话不需乘1000
    var Y = date.getFullYear() + '-';
    var M = (date.getMonth() + 1 < 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1) + '-';
    var D = date.getDate() + ' ';
    var h = date.getHours() + ':';
    var m = date.getMinutes() + ':';
    var s = date.getSeconds();
    return Y + M + D + h + m + s;
}

$(function () {
    $('#checkModel').on('hide.bs.modal',
        function () {
            $("#identityNum").val("")
            var td = $("#lotInfoBody").find("tr").eq(0).find("td");
            td.eq(0).text("");
            td.eq(1).text("");
            td.eq(2).text("");
            $("#confirm").val("");
            $("#Memo").val("");
        })
});