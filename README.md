# 海康威视综合安防信息泄露工具

## 简介

*该工具用于海康威视信息泄露漏洞的检测，若目标redis对外开放会结合redis写入计划任务反弹shell。*

本工具只做学习交流用，代码开源，请勿用做违法用途。

## 使用

 **参数**

```shell
  -c    check is vuln(default)
  -e    reverse a shell
  -u string
        target url
```

**检测**

```shell
./Hikvision_Info_Leak -u http://xxxx.xx.xxx.xx
```

**利用**

```shell
./Hikvision_Info_Leak -e -u http://xxxx.xx.xxx.xx
```

## 注意


1.加入了自动解密.感谢棱角社区的在线解密功能(https://forum.ywhack.com/decrypt.php)

2.自动解密失败会用到的工具

https://github.com/wafinfo/Hikvision