SET CHARSET utf8mb4;
DROP DATABASE IF EXISTS get5_go;
CREATE DATABASE IF NOT EXISTS get5_go DEFAULT CHARACTER SET utf8mb4;

use get5_go;
BEGIN;
SOURCE /docker-init-sqlc-definitions.d/users_schema.sql;
SOURCE /docker-init-sqlc-definitions.d/gameserver_schema.sql;
SOURCE /docker-init-sqlc-definitions.d/teams_schema.sql;
SOURCE /docker-init-sqlc-definitions.d/matches_schema.sql;
SOURCE /docker-init-sqlc-definitions.d/mapstats_schema.sql;
SOURCE /docker-init-sqlc-definitions.d/players_schema.sql;
COMMIT;
