<h1 align="center">
  <br>subscribe2clash<br>
</h1>


<h4 align="center">Clash配置转换</h4>

<p align="center">
  <a href="https://github.com/whoisix/subscribe2clash/actions">
    <img src="https://img.shields.io/github/workflow/status/whoisix/subscribe2clash/Go" alt="Github Actions">
  </a>
  <a href="https://goreportcard.com/report/github.com/whoisix/subscribe2clash">
    <img src="https://goreportcard.com/badge/github.com/whoisix/subscribe2clash">
  </a>
  <a href="https://github.com/whoisix/subscribe2clash/releases">
    <img src="https://img.shields.io/github/release/whoisix/subscribe2clash/all.svg">
  </a>
</p>


## 简介

Clash配置转换，默认自动获取[ACL4SSR](https://github.com/ACL4SSR/ACL4SSR)路由规则。  

支持v2ray\trojan\ss\ssr\ssd订阅转换。**暂不支持单节点转换**  

支持多订阅一起转换，多个订阅连接用英文逗号隔开。

## 启动服务

screen -R s2c
DIR:
104: /home/wwwroot
149: /usr

### 二进制

- [release](https://github.com/whoisix/subscribe2clash/releases)下载对应的版本
- 解压后执行`./subscribe2clash`
- 访问http://localhost:8162/?sub_link=你的订阅链接

### 源码

- 安装Go 1.16+
- `go get github.com/whoisix/subscribe2clash`
- `export GO111MODULE=on`
- 编译 `make build`
- 启动 `./main`
- 访问http://localhost:8162/?sub_link=你的订阅链接

## 命令

- 如果只想生成clash配置文件（没有节点数据），不启用api服务，可使用命令

  ```
  ./main -gc
  ```

- 指定自定义基础配置文件，可在里面添加自定义的路由规则，程序将按照这个文件写入路由信息。可参考[internal/acl/config/default_base_config.yaml](https://github.com/whoisix/subscribe2clash/blob/master/internal/acl/config/default_base_config.yaml) ，`{{.}}`将被程序替换为ACL的路由规则。

  ```
  ./main -b ./yourfile.yaml
  ```

- 指定自定义碎片路由配置文件，可在该配置文件中自定义添加ACL碎片规则文件。可参考[internal/acl/config/default_rules.ini](https://github.com/whoisix/subscribe2clash/blob/master/internal/acl/config/default_rules.ini)

  ```
  ./main -r ./yourfile.ini
  ```

- 指定输出的配置文件。默认情况下配置文件会输出为`./config/acl.yaml`，可以通过以下命令来重新指定。

  ```
  ./main -o ./yourconfig.yaml
  ```


- 启用http代理。由于网络原因，ACL的github源可能连接不上，你可能需要配合代理一起食用。

  ```
  ./main -proxy http://127.0.0.1:7890
  ```

- 指定服务监听地址，默认监听`127.0.0.1:8162`端口。

  ```
  ./main -l 127.0.0.1:8162
  ```

- 指定更新规则频率，单位小时，默认每6小时拉取一次。

  ```
  ./main -t 6
  ```

  

## 参考

- https://github.com/ne1llee/v2ray2clash

## 引用

- https://github.com/ACL4SSR/ACL4SSR

## 测试地址
[http://47.106.211.213:8162/?sub_link=yourlink](http://47.106.211.213:8162/?sub_link=)
