## ShareME

ShareMe is a light-weighted application that enable convenient, quick sharing of files, URLs, and Notes.

### Detailed Introduction

- shareCode (length of 4, e.g., AU8I) is all you need
- 100% "Handwritten", 0% of AI-enabled tools
- Backend: Go, levelDB
- Frontend: Vue3 + Vite
- Language Support: English, <a href="https://github.com/Harold-y/shareMe/blob/main/README_CN.md">中文</a>

#### File Sharing

File sharing is one of our main goals, and the steps to share your file to your friends and other devices of your own could be very simple by using ShareME. We support multi-file upload and easy-use and beautiful interface.
<br/><br/><img src="https://raw.githubusercontent.com/Harold-y/shareMe/main/images/2.png" style="width:650px"/>

Retrieve your files with a length-4 shareCode, zip-download them, or preview them to prevent from virus.

<br/><img src="https://raw.githubusercontent.com/Harold-y/shareMe/main/images/7.png" style="width:650px"/>

We support many formats' preview, including images, texts, programming scripts.

<br/><img src="https://raw.githubusercontent.com/Harold-y/shareMe/main/images/8.png" style="width:650px"/>

#### URL Sharing

Sonme URLs are too long? We can certainly relate to that! Try our URL sharing and transform them into a length-4 shareCode with auto-redirection enabled.

<br/><img src="https://raw.githubusercontent.com/Harold-y/shareMe/main/images/3.png" style="width:650px"/>

#### Note Sharing

Have some thoughts in your mind or want to share some quick notes with your friends? We got you covered! Our note sharing functionality can let you upload the text and your friend can preview them and also download them!

<br/><img src="https://raw.githubusercontent.com/Harold-y/shareMe/main/images/4.png" style="width:650px"/>

<br/><img src="https://raw.githubusercontent.com/Harold-y/shareMe/main/images/6.png" style="width:650px"/>

#### All in One Place

Retrieve File/URL/Note in the home page! Super Easy and Convenient! The application will react differently based on the type. If it is an URL, an auto-redirection will be launched.

<br/><img src="https://raw.githubusercontent.com/Harold-y/shareMe/main/images/1.png" style="width:650px"/>

4-digit/chars, as agreed.

<br/><img src="https://raw.githubusercontent.com/Harold-y/shareMe/main/images/5.png" style="width:650px"/>

#### Deploy

You need to have: node.js (>=v18), Go (>=v1.16)

Frontend (default port 4173, change by adding --port <port> in the npm run preview):
- change main.js "app.config.globalProperties.BASE_URL" to your URL
- npm i
- npm run build
- npm run preview -- --host (or other vite endoresed methods)

Backend (default port 8569, change by editing config file):
- go build
- execute the built file (varies depend on the machine)

Settings (config.txt):
- AdminToken: Token used to manually kill the files and DB record (destroy) (default March7th 三月七真的很好看！)
- Delim: delimiter used to separate file names, no file name can contain these symbol (default ||<_>-<_>||)
- UploadFolder: the folder to store the files (default ./upload)
- MaxUploadSize: the max memory to process the file (default 50M)
- DestroyOnStart: whether to execute destroy on start (default No, change to Yes to enable)
- DailyDestroyHour: how long to execute destroy per hour (default 20)
- Port: the port to run the backend (default 8569)
  
#### Docker

Under development
