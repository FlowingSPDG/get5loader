
-- +migrate Up
ALTER TABLE `round_stats` RENAME COLUMN `fisrt_victim_steamid` TO `first_victim_steamid`
-- +migrate Down
ALTER TABLE `round_stats` RENAME COLUMN `first_victim_steamid` TO `fisrt_victim_steamid`