[![Build Status](https://travis-ci.org/FlowingSPDG/get5-web-go.svg?branch=master)](https://travis-ci.org/FlowingSPDG/get5-web-go)  
get5-web-go
===========================
**Status: Work-In-Progress!!!**

This is recreation of [get5 web panel](https://github.com/splewis/get5-web) with Go API and Vue Front-end.  
Still Work-In-Progress project. PRs are welcome!

## Author
Shugo [**FlowingSPDG**](http://github.com/FlowingSPDG) Kawamura

## WHY
1. Python2.7,which is used in original get5-web, ~~will be no longer supported after end of 2019.~~ **is not supported anymore!!**  
2. Current get5-web needs so many steps to launch(DB migration,python2.7 install,pip package management and venv,etc...). this webpanel may need fewer steps to launch.
3. GOLANG has better performance than Python in some case
4. To support local file-DB insted of MySQL DB for better performance and easy to deploy(this would be optional).
5. To support SPA and better UI/UX design
6. To support get5 HTTP API for developers

## How to use it:
1. Register your CS:GO servers on the "Add a server" section.
2. Register teams on the "Create a Team" section with steamids.
3. Go to the "Create a Match" page.

API Server will send rcon command to load match config(``get5_loadmatch_url <webserver>/match/<matchid>/config``) Then game server loads match and wait for players.

## ScreenShots
![Matches](/screenshots/Matches.PNG?raw=true "Matches list page")
![Match Stats Page](/screenshots/Match.PNG?raw=true "Match Stats Page")

## Requirements
- Open port 8081 to access web-panel and accept RCON connection
- CSGO Server with get5 v0.7.1 and get5_apistats
- MySQL

## Requirements(Developers)
- GO v1.13.5
- original get5-web DB
- CSGO Server with GET5 v0.7.1
- Yarn v1.16.0

## Setup(Developers)
- ``git clone https://github.com/FlowingSPDG/get5-web-go.git $GOPATH/src/github.com/FlowingSPDG/get5-web-go`` (you can fork your own)  
- ``cd $GOPATH/src/github.com/FlowingSPDG/get5-web-go && make deps``
- You're good to GO! edit each `.go` files to fix/add something nice!
- You can test your server by ``go run ./main.go``,and build them by ``make build-all``.You may get binary files.
- To test Vue rendering,``cd ./web/``,``yarn run serve`` and open http://localhost:8081/#.  


## Build(get5-web-go itself doesnt work yet!)
- ``git clone $GOPATH/src/github.com/FlowingSPDG/get5-web-go``  
- ``cd $GOPATH/src/github.com/FlowingSPDG/get5-web-go && make deps``
- ``make build-all``
- You'll get compiled files in ``build`` directly.  

I'm planning to [release](https://github.com/FlowingSPDG/get5-web-go/releases) binary-file for people who feel lazy to build. :P

## Deploy and Launch
- Copy `config.ini.template` to `config.ini` and edit it for your MySQL DB and SteamAPI keys
- `./get5`
- Now it's up!
