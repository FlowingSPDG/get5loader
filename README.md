[![Build Status](https://travis-ci.org/FlowingSPDG/get5-web-go.svg?branch=master)](https://travis-ci.org/FlowingSPDG/get5-web-go)  
GET5-WEB Go
===========================
**Status: Experimental, Not Supported**

This is recreation of [get5 web panel](https://github.com/splewis/get5-web) with Go API and Vue Front-end.  
PRs are welcome!

## Author
Shugo [**FlowingSPDG**](http://github.com/FlowingSPDG) Kawamura

## WHY
1. Python2.7,which is used in original get5-web is not supported anymore  
2. Current get5-web needs so many steps to launch(DB migration,python2.7 install,pip package management and venv,etc...). this webpanel may need fewer steps to launch.
3. GOLANG has better performance than Python in some case
4. To support local file-DB insted of MySQL DB for better performance and easy to deploy(this would be optional).
5. To support SPA and better UI/UX design
6. To support Get5 HTTP/gRPC Streaming API for developers

## How to use
0. Login by your SteamID.
1. Register your CS:GO servers on the "Add a server" section.
2. Register teams on the "Create a Team" section with steamids.
3. Go to the "Create a Match" page.

API Server will send rcon command to load match config( ``get5_loadmatch_url <webserver>/api/v1/match/<matchid>/config`` ) Then game server loads match and wait for players.

## ScreenShots
![Matches](/screenshots/Matches.PNG?raw=true "Matches list page")
![Match Stats Page](/screenshots/Match.PNG?raw=true "Match Stats Page")

## Requirements
- Open HTTP access to access web-panel
- Steam WebAPI Token for handling Steam-Login.
- original MySQL [get5-web](https://github.com/splewis/get5-web) DB

## Requirements(Developers)
- Golang v1.13.5
- NodeJS v10.18.0
- original MySQL [get5-web](https://github.com/splewis/get5-web) DB
- CSGO Server with GET5 v0.7.1 [GET5](https://github.com/splewis/get5/releases)
- Yarn v1.16.0
- Steam WebAPI Token for handling Steam-Login.

## Setup(Developers)
- ``git clone https://github.com/FlowingSPDG/get5-web-go.git $GOPATH/src/github.com/FlowingSPDG/get5-web-go`` (you can fork your own)  
- ``cd $GOPATH/src/github.com/FlowingSPDG/get5-web-go && make deps``
- You're good to GO! edit each `.go` files to fix/add something nice!
- You can test your server by ``go run ./main.go``,and build them by ``make``.You may get binary files in ``./build``.
- To test Vue rendering,``cd ./web/ && yarn run serve`` and open http://localhost:8081/# by your browser.  


## Build
- ``cd $GOPATH/src/github.com/FlowingSPDG/get5-web-go && make``
- You'll get compiled files in ``build`` directly.  
You can use following scripts as your needs :
- ``make build-all`` (or simply, ``make``) Builds Vue and binaries for all supported platforms
- ``make build-linux`` Builds Vue and binaries for Linux
- ``make build-linux-server-only`` Builds binaries for Linux
- ``make build-mac`` Builds Vue and binaries for Mac(darwin)
- ``make build-mac-server-only`` Builds binaries for Mac(darwin)
- ``make build-windows`` Builds Vue and binaries for Windows
- ``make build-windows-server-only`` Builds binaries for Windows
- ``make build-web`` Builds Vue frontend

I'm [releasing](https://github.com/FlowingSPDG/get5-web-go/releases) compiled-files for people who feel lazy to build for each major update.

## Deploy and Launch
- Copy `config.ini.template` to `config.ini` and edit it for your MySQL DB and SteamAPI keys
- `./get5`
- Now it's up!
