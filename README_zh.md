# CtxTimeout middleware

[English](https://github.com/hertz-contrib/ctxtimeout/blob/main/README.md) | 中文

Ctx 超时中间件可以通过 ctx 控制 client 的超时。

## 注意
1. 使用时需要手动在 client 上 Use 该中间件
2. 该中间件和 `client.DoDeadline`, `client.DoTimeout` 功能重复，使用其中一个控制超时即可
3. HTTP2 只能通过使用该中间件控制超时，使用 `client.DoDeadline`, `client.DoTimeout` 不会生效

## 使用
参考 [example](https://github.com/hertz-contrib/ctxtimeout/blob/main/examples)