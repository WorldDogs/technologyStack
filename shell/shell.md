# SHELL、以及Linux常用命令

## 网络测试

### ping

- 通过发送ICMP包来探测目标ip

### host

- 查询dns记录
- 使用

	- host google.com
	- host www.google.com

### 进程管理

### ps

- 显示的当前瞬时的进程状态
- 使用

	- -A -e 所有进程
	- -a 不和本终端相关的所有进程
	- -w 显示加宽，可以显示更多的信息
	- -u 有效使用者的相关进程

- 显示参数

	- VSZ 占用虚拟内存大小
	- RSS 占用内存
	- STAT 进程状态

		- D 不可中断
		- R 运行
		- S 休眠
		- T 暂停
		- Z 僵尸
		- W 没有足够内存分配
		- < 高优先级
		- N 低优先级

### top

- 实时进程信息
- 显示参数

	- CPU行

		- us：用户空间
		- sy 系统内核
		- ni 改变过优先级的占用
		- id 空闲
		- wa I/O等待百分比
		- hi 系统中断
		- si 软件中断
		- ：

- 使用

	- top 1 显示各个逻辑cpuu
	- f 可以选择显示更多字段，并显示字段解释
	- o 选择排序列

### kill

- 常用参数

	- l 获取所有能发给内核的信号

		- HUP 重启
		- KILL 强制杀死
		- TERM 正常结束

### killall

- 通过进程名杀进程 killall httpd

### lsof

- 查询进程打开的文件
- 常用命令

	- lsof -i [46]  [protocol] [@hostname|hostaddr][:service|port]

## 网络管理

### ifconfig

- 网卡信息管理
- 使用

	- ifconfig interface [up|down] 打开关闭网卡
	- 

### route

- 路由设备管理

## 记录常用的shell命令以及一些参数

## 字符串处理

### grep

- 使用

  grep [-ivnc] 'target  content' filename
  
  

- 参数

	- i 忽略大小写
	- v 反向匹配
	- n 输出行号
	- c 统计行数

### sed

- 特点

	- steam editor
	- 以行为单位
	- 对行中数据进行，删除、查找替换、添加、插入

- 使用

	- sed [options] `command` file
	- 文本替换

		- sed -e 's/this/this-/g'

			- 将this替换为this-,如果多个替换就多个-e

	- 删除

		- sed ‘1,2d’ file

			- 删除1~2行，d之前可以只有一个参数
			- $标识最后一行

		- sed '1!d' file

			- 只不删除第1行

		- sed `/Empty/d` file 

			- 删除所有空行

		- 子主题 4

	- 查找替换

		- 默认情况只替换第一个被查找到的
		- sed 's/find/target/' file
		- sed 's/find/target/2' file

			- 匹配行 最多替换2个

		- sed 's/find/target/g' file

			- 替换行中所有匹配的

	- 字符替换

		- 将查找中的每一个字母转换为target
		- sed 'y/OLD/target' file

			- 将每一个O，L，D替换为target

	- 插入

		- sed ‘2 i insert’

			- 在第二行前插入insert

		-  sed ‘2 a insert’

			- 在第二行后插入insert

		- sed ‘2 i insert’

			- 在第二行前插入insert

		- sed '/find/i\Insert' file

			- 在出现find的行后插入Insert

	- 读入

		- sed '/^$/r /etc/passwd' file

			- 将/etc/password的内容读入file的空行后

	- 将sed脚本应用到文件中

		- sed -f se.rules file

### awk

- 特点

	- 按列处理

- 概念

	- $0,$1，$2分别代表整行、第一列、第二列
	- 默认使用空格作为分隔符

- 内置变量及函数

	- NF

		- 列数，$NF可以标识最后一列·

	- NR

		- 行数

	- substr(列，开始字符位置，结束字符位置)

		- 子主题 1

	- length

		- 字符串长度

- 参数

	- F

		- 分隔符

- 计算

	- cat file | awk 'BEGIN{total=0} {total+=$3}END{print total}'

### read

- 读取输入

	- read -p “ll” n
	- read n

## 网络相关

### nc

- 简介（以下基于man简要翻译）

	- 任意的TCP,UDP连接和监听
	- 打开TCP连接，发送UDP包，监听任意的TCP和UDP端口
	- 支持IPv4，ipv6
	- 可以做端口扫描

- 常用命令

	- 端口扫描

		- nc -vzw 2 192.168.1.54 8888-9999

	- 网络连通性测试

		- nc -vvv baidu.com 443

### curl

- 简介

	- 强大的HTTP请求工具

- 常用命令

	- curl [--request <[POST|GET|DELETE...]>] -H "user-argent:test"

		- --request 请求方法
		- -H 请求头
		- --data body数据，
		- --data @x.json 引入外部json
		- -i 显示response header
		- -k 不使用ssl
		- -X  method 请求方法
		- --cert  使用client certificate
		- --output filename 获取文件
		- -u username:password 登录用户密码

