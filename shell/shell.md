# overiew

* 记录常用的shell命令以及一些参数

# 网络相关

## nc

### 简介（以下基于man简要翻译）

* 任意的TCP,UDP连接和监听
* 打开TCP连接，发送UDP包，监听任意的TCP和UDP端口
* 支持IPv4，ipv6
* 可以做端口扫描

### 常用命令

* 端口扫描 
  * nc -vzw 2 192.168.1.54 8888-9999
* 网络连通性测试
  * nc -vvv baidu.com 443



## curl

### 简介

* 强大的HTTP请求工具

### 常用命令

* curl [--request <[POST|GET|DELETE...]>] -H "user-argent:test" 
  * --request 请求方法
  * -H 请求头
  * --data body数据，
  * --data @x.json 引入外部json
  * -i 显示response header
  * -k 不使用ssl
  * -X  method 请求方法
  * --cert  使用client certificate
  * --output filename 获取文件
  * -u username:password 登录用户密码

# 字符串处理

## grep

### 使用

grep [-ivnc] 'target  content' filename

### 参数

* i 忽略大小写
* v 反向匹配
* n 输出行号
* c 统计行数

