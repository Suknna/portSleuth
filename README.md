# Portsleuth
Portsleuth 是一个用于检测当前主机与目标主机之间端口连接情况的工具。
## 使用方法
```
bash复制代码运行
./portsleuth [flags]
./portsleuth [command]
```
### 可用命令

- completion: 为指定的 shell 生成自动补全脚本。
- help: 关于任何命令的帮助。
- version: 显示版本信息。
### 标志

- -g, --goroutine int: 输入 goroutine 的数量，默认为 10。
- -h, --help: 显示 portsleuth 的帮助信息。
- -i, --ip string: 输入 Ipv4 地址，格式如：{192.168.1.2|192.168.1.2-192.168.1.222|192.168.1.3,192.168.3.2,192.168.4.5}。
- -p, --port string: 输入端口，格式如：{80|80-8080|80,22,39,60}。
- -s, --timeout string: 输入超时时间（以秒为单位），默认为 "0.2"。
- -t, --toggle: 切换帮助信息。

使用 "portsleuth [command] --help" 获取更多关于命令的信息。
## 示例
```
bash复制代码运行
./portsleuth -p 22 -i 192.168.1.10
```
输出：
```
复制代码运行
Ipv4 192.168.1.10 :
    opened:
                []
    closed:
                [22]
```
```
bash复制代码运行
./portsleuth -p 22 -i 192.168.1.11
```
输出：
```
复制代码运行
Ipv4 192.168.1.11 :
    opened:
                [22]
    closed:
                []
```
