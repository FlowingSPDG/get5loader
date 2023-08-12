[![Go Build](https://github.com/FlowingSPDG/get5loader/actions/workflows/go_build.yaml/badge.svg)](https://github.com/FlowingSPDG/get5loader/actions/workflows/go_build.yaml)
[![Go Test](https://github.com/FlowingSPDG/get5loader/actions/workflows/go_test.yaml/badge.svg)](https://github.com/FlowingSPDG/get5loader/actions/workflows/go_test.yaml)
[![Downloads](https://img.shields.io/github/downloads/flowingspdg/get5-web-go/total?style=flat-square)](https://github.com/FlowingSPDG/get5loader/releases)
[![LICENSE](https://img.shields.io/github/license/flowingspdg/get5-web-go?style=flat-square)](https://github.com/FlowingSPDG/get5loader/blob/master/LICENSE)

get5loader
===========================
**Status: UNDER DEVELOPMENT**

## Author
Shugo Kawamura  
Github : [**FlowingSPDG**](http://github.com/FlowingSPDG)  
Twitter : [**@FlowingSPDG**](http://twitter.com/FlowingSPDG)

## About
This is match management system for [get5](https://github.com/splewis/get5).  
Inspired by [get5-web](https://github.com/splewis/get5-web).  


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
- Setup database.

## Requirements(Developers)
- Docker
- Go v1.21
- NodeJS and Yarn(Volta)
- MySQL DB
- CSGO Server with GET5 v0.15.0 [GET5](https://github.com/splewis/get5/releases)
- Steam WebAPI Token for handling Steam-Login. ([here](https://steamcommunity.com/dev/apikey))

## Setup(Developers)
- ``git clone https://github.com/FlowingSPDG/get5loader.git $GOPATH/src/github.com/FlowingSPDG/get5loader`` (you can fork your own)  
- ``cd $GOPATH/src/github.com/FlowingSPDG/get5loader && make deps``
- You're good to Go! edit each `.go` files to fix/add something nice!
- You can test your server by ``go run ./cmd/main.go``,and build them by ``make``.You may get binary files in ``./build``.

## Release
I'm [releasing](https://github.com/FlowingSPDG/get5loader/releases) compiled-files for people who feel lazy to build for each major update.

## Deploy and Launch
- Setup environment variables
- Start your compiled binary
- Now it's up!

## License
・[MIT](https://github.com/FlowingSPDG/get5loader/blob/master/LICENSE)
