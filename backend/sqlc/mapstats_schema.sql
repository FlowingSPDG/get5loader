CREATE TABLE `map_stats` (
  `id` BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `match_id` BIGINT NOT NULL,
  `map_number` int(11) NOT NULL,
  `map_name` varchar(64) NOT NULL,
  `start_time` datetime DEFAULT NULL,
  `end_time` datetime DEFAULT NULL,
  `winner` BIGINT DEFAULT NULL,
  `team1_score` int(11) NOT NULL,
  `team2_score` int(11) NOT NULL,
  `forfeit` BOOLEAN DEFAULT NULL,
  FOREIGN KEY (`match_id`) REFERENCES matches(`id`)
);