var host = "http://" + window.location.host;

function lottoinit(){
    $("#lottotbody").empty();
    var req = "";
    $.ajax({
        type: 'POST',
        url: "lottery",
        url: host + "/lottery",
        data: req,
        dataType: "json",
        contentType: "application/json",
        success: function (result) {
            var items = result.items;
            if(items == null) return;
            for(var i = 0,len = items.length; i < len; i++){
                var tr = "<tr><td>" + items[i].name + "</td></tr>";
                $("#lottotbody").append(tr);
            }
        },
        error: function (err) {
            alert("初始化错误!");
        }

    });
}

$(document).ready(function(){
    lottoinit();
});

function clickToLotto(){
    var req = "";
    $.ajax({
        type: 'POST',
        url: host + "/lottery/drawing",
        data: req,
        dataType: "json",
        contentType: "application/json",
        success: function (result) {
            if(result == "kong"){
                alert("未中奖");
            }else{
                alert("恭喜获得：" + result);

            }
        },
        error: function (err) {
            alert("错误" + err);
        }

    });
}

function flushLotto(){
    alert("haha")
    //lottoinit();
}