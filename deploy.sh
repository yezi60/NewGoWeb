#! /bin/sh

# ζζθΏη¨
kill -9 $(pgrep webServer)
cd /home/yezi60/web/NewGoWeb/
git pull https://github.com/yezi60/NewGoWeb.git
cd webServer/
./webServer & 

