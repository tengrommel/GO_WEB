# TCP

抓包工具wireshark

# 代理

golang 编写代理服务
- 监听地址
- 接受连接
- 建立到目标服务的连接
- 数据拷贝
- 关闭连接

# socks5协议
> 指定一个端口作为代理端口
    
    ssh -D 8021 teng@127.0.0.1    
    teng@teng-PC:~$ netstat -nat | grep 8021
    tcp        0      0 127.0.0.1:8021          0.0.0.0:*               LISTEN     
    tcp6       0      0 ::1:8021                :::*                    LISTEN     
    teng@teng-PC:~$ 

# 对称加密和非对称加密
>非对称加密　公钥和私钥都可以加密但必须通过对方解密

rc4加密