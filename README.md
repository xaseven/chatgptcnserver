# ChatGPTCnServer

基于golang的ChatGPT反向代理服务器，项目起的名字有点夸张，见笑 

## 部署方式

1. 从本站下载好 [二进制](https://github.com/xaseven/chatgptcnserver/releases)（或者自己编译）文件，附带config.yaml配置文件即可.
2. nginx配置：  
    location /api/ {  
       add_header Access-Control-Allow-Origin * always;  
     add_header Access-Control-Allow-Credentials false always;  
     add_header Access-Control-Allow-Methods * always;  
     add_header Access-Control-Allow-Headers * always;  
     if ($request_method = 'OPTIONS') {  
        return 204;  
      }  
      proxy_set_header Host $http_host;  
      proxy_set_header  X-Real-IP $remote_addr;  
      proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;  
      proxy_set_header X-Forwarded-Proto $scheme;  
      rewrite ^/api/(.*)$ /$1 break;  
      proxy_pass http://127.0.0.1:8899/;  
    } 
3. config.yml说明详见本文件内部说明
4. 前端html代码详见 https://github.com/xaseven/chatgptcnhtml  
   主要使用接口：/api/base/quest (配合nginx做api目录重写)

## Build from source

1. Clone the repo。go get github.com/xaseven/chatgptcnserver
2. Install dependencies with `go mod tidy`
3. `GOOS=linux GO111MODULE=on GOARCH=amd64 go build`
## 视频教程
https://www.youtube.com/@qilaoxi/videos  

https://www.bilibili.com/video/BV1MR4y1r7xQ/  

QQ群：816577341
## 广告
<p>赚钱项目：<a href="https://sg1-1309278490.cos-website.ap-nanjing.myqcloud.com/?zid=742349">游鱼赚钱项目平台</a></p>
<p>服务器购买①:<a href="https://url.cn/uBTxO4Gm">【腾讯云】云服务器等爆品抢先购，低至4.2元/月</a></p>
<p>服务器购买②:<a href="https://url.cn/XD0oefym">【腾讯云】热卖云产品年终特惠，2核2G轻量应用服务器6.58元/月起</a></p>

