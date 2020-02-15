# Migration from Original get5-web(and get5-web-go =< v0.1.2)

## Preparation
If you gonna migrate from original get5-web database, you'll need to execute preparation before start migration.  
**This step is not required if you gonna create new database from empty MySQL database.**  
- go to migration/migrate_from_original folder. Execute `migrate` binary with following args :
  - -db - Database name. e.g. `"get5"`
  - -host - Database host destination. e.g. `"localhost"`.
  - -port - Database password `3306` by default.
  - -user Database user name e.g. `"get5_user"`.
  - -password - Database user password e.g. `"MyGet5Password"`.
- Example: `./migrate -db get5 -host localhost -port 3306 -user get5_user -password MyGet5Password`

## Migration
**This step is required for all users!**  
This will Add/Modiby Database tables to support new incoming features in get5-web-go.
- goto migration/v0.1.2 folder. Execute `migrate` binary with following args :
  - -direction - `"up"` or `"down"`. in this case, chose `"up"` or leave it by default value `"up"`.
  - -db - Database name. e.g. `"get5"`
  - -host - Database host destination. e.g. `"localhost"`.
  - -port - Database password `3306` by default.
  - -user Database user name e.g. `"get5_user"`.
  - -password - Database user password e.g. `"MyGet5Password"`.
- Example: `./migrate -direction up -db get5 -host localhost -port 3306 -user get5_user -password MyGet5Password`