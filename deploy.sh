#! /bin/sh

# 杀掉进程
kill -9 $(pgrep webServer)
cd /home/yezi60/web/NewGoWeb/
git pull https://github.com/yezi60/NewGoWeb.git
cd webServer/
./webServer & 

