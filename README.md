# cronTask
------------

一个定时任务管理器，基于Go语言和Gin框架开发。用于统一管理项目中的定时任务，提供可视化配置界面、执行日志记录、邮件通知等功能， 

## 项目背景

使用crontab不好管理的问题,很多时候没有机器权限，有些场景需要秒级别的的定时任务，crontab 颗粒度只到分钟级别

## 功能特点

* 统一管理多种定时任务。
* 秒级定时器，使用crontab的时间表达式。
* 可随时暂停任务。
* 记录每次任务的执行结果。
* 执行结果邮件通知。

## 界面截图

![cronTask](https://github.com/tiangyinglun/cronTask/blob/main/assets/images/1653984296407.jpg)


![cronTask](https://github.com/tiangyinglun/cronTask/blob/main/assets/images/1653984290840.jpg)


## 安装说明

系统需要安装Go和MySQL。

获取源码

	$ go get  github.com/tiangyinglun/cronTask

 

创建数据库cron_task，再导入cron_task.sql

$ mysql -u username -p -D cron_task < cron_task.sql

运行

	$ ./main
	或
	$ nohup ./webcron 2>&1 > error.log &
	设为后台运行

访问：

http://localhost:8000

帐号：admin
密码：123456
 