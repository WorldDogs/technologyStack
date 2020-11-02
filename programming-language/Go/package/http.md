# go-http package

## 直接请求

### http.get

### http.post

### http.postform

## 注意事项

### resp.Body.Close()

- 使用完后关闭响应body

### client和transport 是并发安全的 应该创建一次然后重复利用

## http.client

控制请求头、重定向和其他一些设置 使用 clent

## 建立http web server

### http.handle

### http.handlefunc

### http.ListenAndServe

## http.transport

关于传输层、会话层的控制
代理、TLS、keep-alive,compression

