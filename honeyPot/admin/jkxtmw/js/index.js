//登录检查
function check() {
    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;
    if (username === "" || password === "") {
        alert("用户名密码不能为空")
    } else {
        $.ajax({
            type: 'post',
            url: "login",
            data: {
                username: username,
                password: password
            },
            dataType: "json",
            success: function (data) {
                if (data.code === 0) {
                    alert("用户名或密码错误")
                    window.location.href = "login";
                } else {
                    alert("登录成功")
                    window.location.href = "./";
                }
            }
        })
    }
}

//取消配置
function cancel() {
    window.location.href="./";
}

function isValidIP(ip) {
    var reg = /^(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])$/
    return reg.test(ip);
}

//修改配置
function potConfigSave() {
    configid = $('#configid').val();
    username = $('#username').length >0 ? $('#username').val() : "";
    password = $('#password').length >0 ? $('#password').val() : "";
    port = $('#port').length >0 ? $('#port').val() : "0";
    filelist = $('#filelist').length >0 ? $('#filelist').val() : "";
    payload = $('#payload').length >0 ? $('#payload').val() : "";
    fileexists = $('#fileexists').length >0 ? $('#fileexists').val() : "0";
    ip = $('#ip').length >0 ? $('#ip').val() : "";

    if (port > 65535 || port < 1) {
        alert("端口号设置错误！");
        return;
    }

    if (configid === 1003 && !isValidIP(ip)) {
        alert("ip格式错误！");
        return;
    }

    exitflag = false;

    if ($('#file').length > 0 && $('#file').val().length > 0) {
        $.ajax({
            url: 'uploadfile',
            type: 'POST',
            cache: false,
            data: new FormData($('#uploadForm')[0]),
            processData: false,
            contentType: false,
            async: false,
            success: function(data) {
                if (data.code === 0) {
                    fileexists = "1";
                } else {
                    alert("文件上传失败！请检查网络设置！");
                    exitflag = true;
                }
            }
        })
    }

    if (exitflag) {return;}

    $.ajax({
        type: 'post',
        url: "potconfig",
        data: {
            username: username,
            password: password,
            port: port,
            filelist: filelist,
            configid: configid,
            payload: payload,
            fileexists: fileexists,
            ip: ip
        },
        dataType: "json",
        success: function (data) {
            if (data.code === 0) {
                alert("修改成功！请重新启动对应蜜罐使配置生效!");
                window.location.href = "./";
            } else {
                alert("修改失败！请检查输入格式是否正确！");
            }
        }
    })
}

//改变蜜罐状态（0：关闭；1：运行）
function potcontrol(configid, state) {
    //弹出确认框,提示用户先对蜜罐进行配置
    var msg = state === 0 ? "您确认启动吗？(启动之前请确认是否对蜜罐进行了配置)": "您确认关闭吗？";
    if (confirm(msg) === false) {
        return;
    }

    $.ajax({
        type: 'post',
        url: "potcontrol",
        data: {
            configid: configid,
            state: state
        },
        dataType: "json",
        success: function (data) {
            if (data.code === 0) {
                setTimeout(function(){ if (state === 0) {
                    alert("启动成功！");
                } else {
                    alert("关闭成功！");
                }
                    window.location.href = "./";}, 2000);
            } else {
                alert("系统错误！请稍后重试！");
            }
        }
    })
}

//修改密码
function pwd() {
    const oldpass = document.getElementById("oldpass").value;
    const password1 = document.getElementById("password1").value;
    const password2 = document.getElementById("password2").value;
    if (oldpass === "") {
        alert("旧密码不能为空")
    }
    if (password1 === "" || password2 === "") {
        alert("输入的密码不能为空")
    }
    if (password1 !== "" && password2 !== "" && password1 !== password2) {
        alert("两次输入的密码不一致")
    } else {
        $.ajax({
            type: 'post',
            url: "changepass",
            data: {
                oldpass:oldpass,
                password1: password1,
                password2: password2
            },
            dataType: "json",
            success: function (data) {
                if (data.code === -1) {
                    alert("旧密码输入错误")
                }
                if (data.code === 0) {
                    alert("两次输入的密码不一致，请重新输入")
                }
                if (data.code === 3) {
                    alert("密码长度必须大于8位，并且必须包含大小写字母,数字和特殊符号")
                }
                if (data.code === 1)  {
                    alert("密码修改成功，请重新登录")
                    window.location.href = "logout";
                }
                if (data.code === 2)  {
                    alert("密码修改失败，请重新修改")
                }
            }
        })
    }
}

function countSubstr(str, substr) {
    const regex = new RegExp(substr, 'g');
    const result = str.match(regex);
    return !result ? 0 : result.length
}

$(function () {
    $('.logout').click(function() {
        let count = countSubstr(window.location.pathname, "/");
        path = ""
        for (let i = 0; i < count -2; i++) {
            path += "../"
        }
        if (confirm('确定退出？')) {
            window.location.href = path + "logout"
        }
    });

    $("#fileinput").fileinput({
        language: 'zh',
        dropZoneTitle: '将license文件拖拽到这里进行上传',
        showUpload: true,
        maxFileSize: 1024,
        uploadUrl: "license",
        uploadAsync: true,
        allowedFileExtensions: ['dat']
    });

    //上传成功后执行
    $("#fileinput").on("fileuploaded", function (event, data,) {
        if (data.response.msg === "upload success"){
            alert("license文件上传成功,请重启平台完成授权!")
        }
    });
})