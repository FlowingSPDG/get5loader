get5-web-go
===========================
**Status: Work-In-Progress!!!**

This is recreation of [get5 web panel](https://github.com/splewis/get5-web) with GOLANG.

## WHY?
1. Python2.7 will be no longer supported after Jan.2020. very soon!!
2. Current get5-web needs so many steps to launch(DB migration,python install,venv,etc...)
3. GOLANG has better performance than Python(I guess...??)
4. To support local file-DB insted of MySQL DB for better performance and easy to manage(this would be optical).

## Requirements:
- Open port 8081 to access web-panel and accept RCON connection

## Requirements(developers):
- GO v1.13.3
- original get5-web DB
