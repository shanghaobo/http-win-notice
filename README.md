# HttpWinNotice


<!-- PROJECT LOGO -->
<br />

<p align="center">
  <a href="https://github.com/shaojintian/Best_README_template/">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a>

<h3 align="center">HTTP Windows消息通知</h3>
  <p align="center">
  
   

</p>

> 一个运行在Windows本机的服务，监听http接口请求触发windows消息通知


## 功能

- [x] http接口请求
- [x] Windows 消息弹窗
- [x] 历史消息记录
- [x] 开机启动
- [x] 托盘控制
- [x] 端口配置

Feture：
- [ ] Server端

## 架构

![](images/jiagou.png)

其中server部分还未开发完成

## 使用

1. 直接下载编译后的exe文件
2. 双击exe启动
3. 右击托盘图标勾选开机启动
4. 测试通知
   - 方式1
     浏览器输入链接`http://127.0.0.1:19000/api/toast?msg=哈喽，我是新的消息通知`
   - 方式2 (python调用示例)
     ```python
     import requests
        
     title = "自定义通知标题"
     msg = "我是测试消息通知啦啦啦啦啦"
     res = requests.get(f"http://127.0.0.1:19000/api/toast?msg={msg}&title={title}")
     print(res.json())
       ```
5. 可右击托盘图标修改端口号
6. 可打开web界面查看历史消息记录

## 效果

- 程序运行
  ![](images/demo3.png)


- 调用http消息通知接口
  ![](images/demo2.png)

- 消息通知记录
  ![](images/demo1.png)



