# 基于WebRTC的分布式实时音视频系统

## 编译环境
go version go1.19.3 linux/amd64

理论上低版本也可编译，具体支持的最低版本需自行实验

## 文件目录
### src文件
源代码：包括本项目的framework模块、action模块、log模块（基于开源项目[glog](https://github.com/golang/glog)修改而来）
### static
静态网站文件，包括html以及相对应的js代码
### build.sh
编译生成服务器二进制执行文件