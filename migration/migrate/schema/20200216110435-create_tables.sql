
-- +migrate Up
ALTER TABLE `match` MODIFY `cvars` varchar(512) DEFAULT NULL ;
-- +migrate Down
ALTER TABLE `match` MODIFY `cvars` varchar(40) DEFAULT NULL ;