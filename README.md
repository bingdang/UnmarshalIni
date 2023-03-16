# UnmarshalIni
Unmarshalini：反射方式实现初始化配置文件反序列化
```
felix@MacBook-Pro 2023Project % cat ./config.ini
#this is comment
;this a comment
;[]表示一个section
[server]
ip = 10.238.2.2
port = 8080

[mysql]
username = root
passwd = admin
database = test
host = 192.168.10.10
port = 8000
timeout = 1.2
default-character-set = utf8mb4

felix@MacBook-Pro 2023Project % go run 2023Project
ini.Config{SvcCfg:ini.Server{Ip:"10.238.2.2", Port:8080}, DbCfg:ini.Mysql{Username:"root", Passwd:"admin", Database:"test", Host:"192.168.10.10", Port:8000, Timeout:1.2, DefaultCharacterSet:"utf8mb4"}}
```
