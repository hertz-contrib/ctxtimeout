# CtxTimeout middleware

English | [中文](https://github.com/hertz-contrib/ctxtimeout/blob/main/README_zh.md)

The ctx timeout middleware allows client timeouts to be controlled through Context.


## Notes
1. When using this middleware, it is necessary to manually add it on the client.
2. The functionality of this middleware overlaps with that of client.DoDeadline and client.DoTimeout, so only one of them is needed to control the timeout.
3. For HTTP2, this middleware is the only way to control the timeout, and `client.DoDeadline` and `client.DoTimeout` don't work.
## Usage
Please refer to the example code: https://github.com/hertz-contrib/ctxtimeout/blob/main/examples