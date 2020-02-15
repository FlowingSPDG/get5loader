
-- +migrate Up

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


ALTER TABLE `match`
  ADD `cvars` varchar(40) DEFAULT NULL ,
  ADD `side_type` varchar(40) DEFAULT "standard" ,
  ADD `is_pug` tinyint(1) DEFAULT NULL ;

ALTER TABLE `map_stats` ADD `dem_path` varchar(40) DEFAULT NULL ;

ALTER TABLE `team` ADD `steamids` varchar(400) DEFAULT NULL ;


CREATE TABLE `round_stats` (
  `id` int(11) NOT NULL,
  `match_id` int(11) DEFAULT NULL,
  `map_id` int(11) DEFAULT NULL,
  `first_killer_steamid` varchar(40) DEFAULT NULL,
  `fisrt_victim_steamid` varchar(40) DEFAULT NULL,
  `second_killer_steamid` varchar(40) DEFAULT NULL,
  `second_victim_steamid` varchar(40) DEFAULT NULL,
  `third_killer_steamid` varchar(40) DEFAULT NULL,
  `third_victim_steamid` varchar(40) DEFAULT NULL,
  `fourth_killer_steamid` varchar(40) DEFAULT NULL,
  `fourth_victim_steamid` varchar(40) DEFAULT NULL,
  `fifth_killer_steamid` varchar(40) DEFAULT NULL,
  `fifth_victim_steamid` varchar(40) DEFAULT NULL,
  `sixth_killer_steamid` varchar(40) DEFAULT NULL,
  `sixth_victim_steamid` varchar(40) DEFAULT NULL,
  `seventh_killer_steamid` varchar(40) DEFAULT NULL,
  `seventh_victim_steamid` varchar(40) DEFAULT NULL,
  `eighth_killer_steamid` varchar(40) DEFAULT NULL,
  `eighth_victim_steamid` varchar(40) DEFAULT NULL,
  `ninth_killer_steamid` varchar(40) DEFAULT NULL,
  `ninth_victim_steamid` varchar(40) DEFAULT NULL,
  `tenth_killer_steamid` varchar(40) DEFAULT NULL,
  `tenth_victim_steamid` varchar(40) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE `round_stats`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `round_stats`
  ADD CONSTRAINT `round_stats_ibfk_1` FOREIGN KEY (`match_id`) REFERENCES `match` (`id`),
  ADD CONSTRAINT `round_stats_ibfk_2` FOREIGN KEY (`map_id`) REFERENCES `map_stats` (`id`);
COMMIT;

-- +migrate Down
SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";



ALTER TABLE `match`
  drop `cvars` ,
  drop `side_type` ,
  drop `is_pug` ;

ALTER TABLE `map_stats`
  drop `dem_path` ;

DROP TABLE `round_stats`;