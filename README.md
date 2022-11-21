# 基于WebRTC的分布式实时音视频系统

## 编译环境
go version go1.19.3 linux/amd64

理论上低版本也可编译，具体支持的最低版本需自行实验

## 文件目录
### src文件
源代码部分：
1. main.go - 主程序
2. router.go - action的路由实现
3. framework模
4. action模块
5. log模块（基于开源项目[glog](https://github.com/golang/glog)修改而来）
6. 系统配置模块(基于开源项目[goconfig](https://github.com/unknwon/goconfig))
### static
静态网站文件，包括html以及相对应的js代码
### conf
服务器配置文件夹，服务器启动时加载配置，同时还包括https的证书文件`fullchain.pem`和`privkey.pem`，使用时请将自己的https证书放进去
### build.sh
编译生成服务器二进制执行文件