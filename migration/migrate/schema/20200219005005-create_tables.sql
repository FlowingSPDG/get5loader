
-- +migrate Up
ALTER TABLE `round_stats` CHANGE `fisrt_victim_steamid` `first_victim_steamid` varchar(40);
ALTER TABLE `round_stats` MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

-- +migrate Down
ALTER TABLE `round_stats` CHANGE `first_victim_steamid` `fisrt_victim_steamid` varchar(40);