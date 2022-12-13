/*
 * @Author: Sugar 45682h@gmail.com
 * @Date: 2022-11-21 20:50:42
 * @Describe: 
 */
'use stript'

var pushBtn = document.getElementById("pushBtn")

pushBtn.addEventListener("click", startPush)

var uid = $("#uid").val()
var streamName = $("#streamName").val()
var audio = $("#audio").val()
var video = $("#video").val()

function startPush() {
    console.log("send push: /signaling/push")

    $.post("/signaling/push",
        {"uid": uid, "streamName": streamName, "audio": audio, video: video},
        function(data, textStatus) {
            console.log("push response: " + JSON.stringify(data))
            if (textStatus == "success" && data.errNo == 0) {
                console.log("success")
                $("#tips1").html("<font color='blue'>推流请求成功!</font>");
            } else {
                $("#tips1").html("<font color='red'>推流请求失败!</font>");
            }
        },
        "json" 
    )
}