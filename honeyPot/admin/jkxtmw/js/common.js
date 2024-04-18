//日期组件
$(function () {
    $(".daterangepicker").datetimepicker({
        fontAwesome: 'font-awesome', //解决图标缺失问题，直接用font-awesome代替
        forceParse: 0, //设置为0，时间不会跳转1899，会显示当前时间。
        language: 'zh-CN', //显示中文
        format: 'yyyy-mm-dd hh:ii:ss', //日期格式化
        weekStart: 1, //每周的第一天是
        initialDate: new Date(), //初始化当前日期
        endDate: new Date(), //结束日期，后面的不可选
    });

    //设置endDate的最小值，不能小于startDate
    $("#startDate").datetimepicker().on('changeDate', function (e) {
        $('#endDate').datetimepicker('setStartDate', e.date);
    })

    //设置startDate的最大值，不能大于endDate
    $("#endDate").datetimepicker().on('changeDate', function (e) {
        $('#startDate').datetimepicker('setEndDate', e.date);
    })

    // 加载分页组件
    setpage()

    //获取url参数
    function GetQueryString(name)
    {
        var reg = new RegExp("(^[a-z]{5}/)"+ name +"/([^&]*)(&|$)");
        var r = window.location.pathname.substr(1).match(reg);
        if(r!=null)return  unescape(r[2]); return null;
    }

    //条件查询
    $("#btn1").click(function () {
        let tablename =  GetQueryString("log")
        let startDate = $("#startDate").val()
        let endDate = $("#endDate").val()
        let ip = $("#ip").val()
        let reqmethod = $("#m option:selected").val()
        if (startDate !== "" && endDate === "") {
            alert("结束日期不能为空")
        }
        if (endDate !== "" && startDate === "") {
            alert("开始日期不能为空")
        }
        setpage()
        $.ajax({
            type: 'post',
            url: 'logsearch',
            data: {
                tablename: tablename,
                startDate: startDate,
                endDate: endDate,
                ip: ip,
                reqmethod: reqmethod
            },
            dataType: 'json',
            success: function (data) {
                $.each(data.res, function (index, obj) {

                    $("#tr-" + index).show()

                    $("#time-" + index).text(obj.Time)
                    $("#ip-" + index).text(obj.ClientIP)
                    $("#code-" + index).text(obj.StatusCode)
                    $("#method-" + index).text(obj.ReqMethod)
                    $("#uri-" + index).text(obj.ReqUri)
                    $("#message-" + index).text(obj.Full_message)

                    if (index === data.res.length - 1) {
                        for (let i = index + 1; i < 10; i++) {
                            $("#tr-" + i).hide()
                        }
                    }
                })
            }
        })

    })

    //获取分页数量
    function getnum() {
        let tablename =  GetQueryString("log")
        let startDate = $("#startDate").val()
        let endDate = $("#endDate").val()
        let ip = $("#ip").val()
        let reqmethod = $("#m option:selected").val()
        let num = 0;
        $.ajax({
            type: 'post',
            url: 'logcount',
            data: {
                tablename: tablename,
                startDate: startDate,
                endDate: endDate,
                ip: ip,
                reqmethod: reqmethod
            },
            dataType: "json",
            async: false,
            success: function (data) {
                num = data.datanum
            }
        })
        return num
    }

    // 分页函数
    function setpage() {
        $(".myPagination").Pagination({
            page: 1,
            count: getnum(),
            groups: 5,
            onPageChange: function (page) {
                let startDate = $("#startDate").val()
                let endDate = $("#endDate").val()
                let ip = $("#ip").val()
                let reqmethod = $("#m option:selected").val()
                $.ajax({
                    type: 'post',
                    url: 'log',
                    data: {
                        page: page,
                        startDate: startDate,
                        endDate: endDate,
                        ip: ip,
                        reqmethod: reqmethod
                    },
                    dataType: "json",
                    success: function (data) {
                        $.each(data.res, function (index, obj) {

                            $("#tr-" + index).show()

                            // console.log(obj.ReqUri)
                            $("#time-" + index).text(obj.Time)
                            $("#ip-" + index).text(obj.ClientIP)
                            $("#code-" + index).text(obj.StatusCode)
                            $("#method-" + index).text(obj.ReqMethod)
                            $("#uri-" + index).text(obj.ReqUri)
                            $("#message-" + index).text(obj.Full_message)
                            if (index === data.res.length - 1) {
                                for (let i = index + 1; i < 10; i++) {
                                    $("#tr-" + i).hide()
                                }
                            }
                        })
                    }
                })
            }
        });
    }

})