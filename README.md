# wechat_motion_cheat（微信运动步数作弊）

## 前言
为了在微信运动步数上霸榜，最近在网上看了很多刷微信运动步数的代码实现，感觉很多人都刻意将简单的事情复杂化，代码写得冗长又晦涩，有种生怕别人看懂的感觉，无语。
因此自己用golang实现了一版刷步数脚本，并且对刷步数的脚本实现原理进行剖析，让后来者更清晰更易于理解。
本质上就是3次http请求，有代码基础的朋友也完全可以自己去实现。

## 原理
本质上是利用`Zepp life`对外暴露的接口漏洞，通过脚本将步数刷到第三方App服务器，再由第三方App同步到微信运动实现
<br>
![image](https://github.com/zhutiancheng1997/wechat_motion_cheat/blob/main/img/1.png)

## 用途
- 朋友圈装b ^^
- 做公益捐步数 👍

## 刷步步骤
1.下载 `Zepp life` App
<img src="https://github.com/zhutiancheng1997/wechat_motion_cheat/blob/main/img/step/zepp_life_app.png" height="240px">

2.用手机号或者邮箱号注册账号
<img src="https://github.com/zhutiancheng1997/wechat_motion_cheat/blob/main/img/step/register1.png" height="240px">
<img src="https://github.com/zhutiancheng1997/wechat_motion_cheat/blob/main/img/step/register2.png" height="240px">

3.绑定 `Zepp life` 账号到微信运动
<img src="https://github.com/zhutiancheng1997/wechat_motion_cheat/blob/main/img/step/bind_wx1.png" height="240px">
<img src="https://github.com/zhutiancheng1997/wechat_motion_cheat/blob/main/img/step/bind_wx2.png" height="240px">

4.通过脚本更新步数到`Zepp life`服务端，步数会自动同步到微信运动
确认你已经安装golang环境，clone本项目代码，在cmd目录下执行
```go build```
然后执行
```./cheater -user {你的zepplife账号} -password {你的zepp life密码} -step {你想刷的步数}```
刷步数。
(如成功无法刷步数，建议重新注册一个新账号并重新绑定)

## 脚本实现细节
绿色部分为脚本实现的内容，分别是三次http请求。
- 第一次请求根据user+password获取access code
- 第二次根据access code获取app token和uid
- 第三次push步数到服务器
![image](https://github.com/zhutiancheng1997/wechat_motion_cheat/blob/main/img/2.png)

## 注意⚠️
- 脚本一天内可多次执行，但是每次更新步数服务端会做check，只能比上次更新的更多，不能更少，更少则默认无效。
- 短时间不要多次执行脚本，会触发服务端限流，表现在第一个http请求会返回429响应。

### 如果觉得好用，请给一个免费的[star](https://github.com/zhutiancheng1997/wechat_motion_cheat/)吧
### 有问题可以在issue中提问交流

