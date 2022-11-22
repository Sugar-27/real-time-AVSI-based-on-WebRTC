# 基于WebRTC的分布式实时音视频系统——信令模块

## 编译环境
go version go1.19.3 linux/amd64

理论上低版本也可编译，具体支持的最低版本需自行实验

## 文件目录
### src文件
源代码部分：
1. main.go - 主程序
2. router.go - action的路由注册
3. framework模块 - 程序的主要框架模块，包括http服务器、日志封装api、xrtc、配置读取以及初始化设置（将前面几项一同初始化好供`main.go`中调用）
4. action模块 - 推拉流模块（包括静态资源调用的推拉流以及xrtc服务器的推拉流）
5. glog模块 - 日志模块（基于开源项目[glog](https://github.com/golang/glog)修改而来）
6. goconfig模块 - 系统配置模块（基于开源项目[goconfig](https://github.com/unknwon/goconfig)）
7. comerrors模块 - 推拉流请求url中解析参数的异常处理
### static
静态网站文件，包括html以及相对应的js代码
### conf
服务器配置文件夹，服务器启动时加载配置，同时还包括https的证书文件`fullchain.pem`和`privkey.pem`，使用时请将自己的https证书放进去
### build.sh
编译生成服务器二进制执行文件