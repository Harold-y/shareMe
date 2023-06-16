## ShareME

ShareMe是一个轻量级的，方便、快速地共享文件、URL链接和笔记的应用程序。

### 介绍

- 只需要 shareCode（访问代码） (长度为4，例如AU8I)
- 100%全站程序员匠心制造, 0% AI添加成分
- 后端: Go, leveldb
- 前端: Vue3 + Vite
- 语言支持: <a href="https://github.com/Harold-y/shareMe/blob/main/README.md">English</a>, 中文

#### 分享文件

文件共享是我们的主要目标之一，通过使用ShareME，将你的文件分享给你的朋友，或者其他设备的步骤非常简单。支持多文件上传，并提供简单易用的漂亮交互界面。
<br/><br/><img src="https://raw.githubusercontent.com/Harold-y/shareMe/main/images/2.png" style="width:650px"/>

用一个长度为4的shareCode获取你的文件，支持压缩下载，支持预览。

<br/><img src="https://raw.githubusercontent.com/Harold-y/shareMe/main/images/7.png" style="width:650px"/>

我们支持许多格式的预览，包括图像、文本、程序代码。

<br/><img src="https://raw.githubusercontent.com/Harold-y/shareMe/main/images/8.png" style="width:650px"/>

#### 分享URL链接

有些URLs太长了？深有体会! 试试我们的URL共享，把它们变成一个长度为4的shareCode，按下按钮自动重定向到目标URL。

<br/><img src="https://raw.githubusercontent.com/Harold-y/shareMe/main/images/3.png" style="width:650px"/>

#### 分享笔记

你心中有一些想法，或者想与你的朋友分享一些笔记？我们帮你解决! 我们的笔记分享功能可以让你上传文字，你的朋友可以预览它们，也可以下载它们!

<br/><img src="https://raw.githubusercontent.com/Harold-y/shareMe/main/images/4.png" style="width:650px"/>

<br/><img src="https://raw.githubusercontent.com/Harold-y/shareMe/main/images/6.png" style="width:650px"/>

#### shareCode索取

在主页上检索文件/URL/笔记! 超级简单和方便! 程序将根据类型作出不同的反应。如果是一个URL，将启动自动重定向。

<br/><img src="https://raw.githubusercontent.com/Harold-y/shareMe/main/images/1.png" style="width:650px"/>

仅需四位！

<br/><img src="https://raw.githubusercontent.com/Harold-y/shareMe/main/images/5.png" style="width:650px"/>

#### 部署

你的服务器/主机需要有: node.js (>=v18), Go (>=v1.16)

前端（默认 port 4173, 通过添加--port <port> 参数在执行 npm run preview 的时候）:
- 更改 main.js "app.config.globalProperties.BASE_URL" 成你自己的后端URL
- npm i
- npm run build
- npm run preview -- --host (或者其它Vite部署方式)

后端（默认 port 8569, 在config.txt更改）:
- go build
- 执行build好的文件（根据机器不同而不同）

后端设置 (config.txt):
- AdminToken: 用于手动清理文件和DB记录的令牌（默认为March7th。三月七真的很好看！）
- Delim: 用于分隔文件名的分隔符，任何文件名都不能包含这些符号（默认 ||<_>-<_>||）
- UploadFolder：存储文件的文件夹（默认为./upload）
- MaxUploadSize：处理文件的最大内存（默认为50M）
- DestroyOnStart：是否在启动时执行清理文件和DB记录（默认为No，改成Yes即可启用）
- DailyDestroyHour：每小时执行清理的时间（默认为20）
- Port：运行后台的端口（默认为8569）。
  
#### Docker

尚待开发
