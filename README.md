[![Build Status](https://github.com/FlowingSPDG/get5-web-go/actions/workflows/go_build/badge.svg)
[![Downloads](https://img.shields.io/github/downloads/flowingspdg/get5-web-go/total?style=flat-square)](https://github.com/FlowingSPDG/get5-web-go/releases)
[![LICENSE](https://img.shields.io/github/license/flowingspdg/get5-web-go?style=flat-square)](https://github.com/FlowingSPDG/get5-web-go/blob/master/LICENSE)

GET5-WEB GO
===========================
**Status: UNDER DEVELOPMENT**

## Author
Shugo Kawamura  
Github : [**FlowingSPDG**](http://github.com/FlowingSPDG)  
Twitter : [**@FlowingSPDG**](http://twitter.com/FlowingSPDG)

## About
This is recreation of [get5 web panel](https://github.com/splewis/get5-web) (Python2.7) with Go and Vue.  
Front-end looks pretty same with original get5-web. but API logic is not exactly the same. but most functions should be compatible.


## WHY
1. Python2.7,which is used in original get5-web is not supported anymore  
2. Current get5-web needs so many steps to launch(DB migration,python2.7 install,pip package management and venv,etc...). this webpanel may need fewer steps to launch.
3. Cloud Native System.
4. Go has better performance than Python in some case
5. To support SPA and better UI/UX design with React

## How to use
1. Login by your SteamID.
2. Register your CS:GO servers on the "Add a server" section.
3. Register teams on the "Create a Team" section with steamids.
4. Go to the "Create a Match" page.

API Server will send rcon command to load match config( ``get5_loadmatch_url <webserver>/api/v1/match/<matchid>/config`` ) Then game server loads match and wait for players.

## ScreenShots
![Matches](/screenshots/Matches.PNG?raw=true "Matches list page")
![Match Stats Page](/screenshots/Match.PNG?raw=true "Match Stats Page")

## Requirements
- Open HTTP access to access API.
- Setup environment variables.
- Migrate database.

## Requirements(Developers)
- Docker
- Go v1.20
- NodeJS and Yarn(Volta)
- MySQL DB(came
- CSGO Server with GET5 v0.15.0 [GET5](https://github.com/splewis/get5/releases)
- Steam WebAPI Token for handling Steam-Login. ([here](https://steamcommunity.com/dev/apikey))

## Setup(Developers)
- ``git clone https://github.com/FlowingSPDG/get5-web-go.git $GOPATH/src/github.com/FlowingSPDG/get5-web-go`` (you can fork your own)  
- ``cd $GOPATH/src/github.com/FlowingSPDG/get5-web-go && make deps``
- You're good to Go! edit each `.go` files to fix/add something nice!
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


## Release
I'm [releasing](https://github.com/FlowingSPDG/get5-web-go/releases) compiled-files for people who feel lazy to build for each major update.

## Deploy and Launch
- Setup environment variables
- Start your compiled binary
- Now it's up!

## License
・[MIT](https://github.com/FlowingSPDG/get5-web-go/blob/master/LICENSE)
