
-- +migrate Up
ALTER TABLE `round_stats` CHANGE `fisrt_victim_steamid` `first_victim_steamid` varchar(40);
ALTER TABLE `round_stats` MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
ALTER TABLE `round_stats` DROP `winner`;
ALTER TABLE `round_stats` ADD `winner` varchar(32) NOT NULL;
ALTER TABLE `round_stats` DROP FOREIGN KEY `round_stats_ibfk_2`;
ALTER TABLE `round_stats` DROP `map_id`;
ALTER TABLE `round_stats` ADD `map_number` int(11) DEFAULT NULL ;

-- +migrate Down
ALTER TABLE `round_stats` CHANGE `first_victim_steamid` `fisrt_victim_steamid` varchar(40);