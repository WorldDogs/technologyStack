# Overiew

* 本篇刻画Docker的学习目标
* 原理层面
  * 是什么
  * 历史
  * 核心原理
  * 设计初衷
* 使用层面
  * 熟练使用常用的命令
  * 不常用的命令知道如何去查



# Content

## 核心原理

### LXC

* docker基于LXC(Linux Container)

### namespace

* 通过**namespace**实现资源的隔离

### control group

* 通过cgroup进行资源限制

### Copy On Write (写时复制|COW)

* 通过写时赋值进行高效的文件操作
  * 写时复制的概念：A复制得到B时，并没有真正复制A，而是让B指向A，只有当A或者B进行写操作时，才进行对应块的复制。linux的fork也使用了写时复制技术,fock时将内存标记为Read Only，当进行写操作时，发生中断然后进行页复制。

# 是什么

* 其实就是



