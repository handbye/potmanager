# potManager

这个项目是由[这个项目](https://github.com/handbye/SimpleHoneyPot)演化而来的，加了web管理和详细的日志记录。

> 此项目是2021年编写的，其中用到的蜜罐技术现在已经过时了，仅作为学习使用～～。

特点：

- 跨平台
- 编译后仅一个二进制文件，无任何依赖
- 方便使用可自行扩展插件支持其他蜜罐

使用方法：

使用`main.exe -h`即可查看启动帮助

![image-20240418211709340](images/image-20240418211709340.png)

**初次使用时必须使用 `-init`参数初始化数据。**

启动完毕后会提示后台地址：

![image-20240418211803916](images/image-20240418211803916.png)

然后拼接此地址即可登录后台：

例如：

```txt
http://localhost:8080/bvovxsli/
```

平台默认用户名密码：venus/venus@2021

![image-20240418211636177](images/image-20240418211636177.png)

登录后务必修改密码！！！
