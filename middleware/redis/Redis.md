# Redis

## 备份与恢复

### RDB

### AOF

- 定义

	- Append Only File
	- 是一种写后日志，即当命令执行完成后才记录日志

## 数据结构

### string

### hash

- 相关命令

	- hget

		-  hget dev01:temperature 202012170914

	- hmget

	  获取多个hash键

		-  HMGET dev01:temperature 202012170813 202012170914
		- 获取hash 多个键对应的值

	- hmset

	  增加多个hash键值

		- hmset wd 1 23 2 24 3 25
		- 设置多个hash 键值

	- hgetall

	  获取hash中所有键值对

		- 获取hash 所有的键值

	- HINCRBY

	  给key的某个字段增加***

		- 字段增加值

	- hkeys

		- 获取所有字段

	- hvals

		- 获取所有值

	- hsetnx

		- 字段不存在时设置

	- hscan

		- 使用场景

			- 海量数据查询时，会导致阻塞主线程，所以不能直接hkeys,hvals这种操作，hscan是基于游标的查询

### list

### set

- 集合

	- 支持各种集合操作

### sorted set

- zadd
- zcount

	- 根据score 统计

### hyperloglog

### GEO

### steam

### set

- sorted list
- 集合，元素唯一

### 数据结构与底层实现 - 图

- 

## 发布订阅

### 支持pattern 订阅

## 事务

### multi

- 开始原子操作，以后输入的命令都会进入队列，但不执行

### exec

- 执行内部队列的所有命令

### tips

- multi过后，如果出现错误，还是能正长执行后续命令，但是exec的时候会报错

## 实践

### 时序数据存储

- sorted set 实现

	- 优点

		- 可以排序 聚合运算

- hash实现

	- 优点
	- 缺点

- redisTimeSeries

## 优先级

### 使用

### 运维

### 原理

