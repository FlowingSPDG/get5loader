get5-web-go
===========================
**Status: Work-In-Progress!!!**

This is recreation of [get5 web panel](https://github.com/splewis/get5-web) with GOLANG / VueJS.  
Still Work-In-Progress project. PRs are welcome!

## Author:
Shugo **FlowingSPDG** Kawamura

## WHY?
1. Python2.7,which is used in original get5-web, will be no longer supported after Jan.2020. very soon!!
2. Current get5-web needs so many steps to launch(DB migration,python install,pip package management and venv,etc...). this webpanel may need fewer steps to launch.
3. GOLANG has better performance than Python(I guess...??)
4. To support local file-DB insted of MySQL DB for better performance and easy to deploy(this would be optional).

## Requirements:
- Open port 8081 to access web-panel and accept RCON connection
- MySQL

## Requirements(Developers):
- GO v1.13.5
- original get5-web DB

## Setup(Developers)
- ``git clone $GOPATH/src/github.com/FlowingSPDG/get5-web-go``  
- ``cd $GOPATH/src/github.com/FlowingSPDG/get5-web-go && go get``
- You're good to GO! edit each .go files to fix/add something nice!
- You can test your server by ``go run ./main.go``,and build them by ``go build ./main.go``.You may get binary files for your OS.
- To test Vue rendering,``cd ./web/``,``yarn run dev`` and open http://localhost:8081/#.  


## Build(get5-web-go itself doesnt work yet!)
- ``git clone $GOPATH/src/github.com/FlowingSPDG/get5-web-go``  
- ``cd $GOPATH/src/github.com/FlowingSPDG/get5-web-go && go get``
- ``go build -o get5``
- You'll get `get5` file(depends on your Operating System).  
I'm planning to release binary-file for people who feel lazy to build. :P

## Deploy and Launch
- Edit `config.ini` to your MySQL DB and SteamAPI keys
- `./get5`
- Now it's up!
