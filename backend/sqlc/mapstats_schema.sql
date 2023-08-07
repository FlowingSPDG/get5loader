CREATE TABLE `map_stats` (
  `id` VARCHAR(36) NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `match_id` VARCHAR(36) NOT NULL,
  `map_number` TINYINT UNSIGNED NOT NULL,
  `map_name` varchar(64) NOT NULL,
  `start_time` datetime DEFAULT NULL,
  `end_time` datetime DEFAULT NULL,
  `winner` VARCHAR(36) DEFAULT NULL,
  `team1_score` TINYINT UNSIGNED NOT NULL,
  `team2_score` TINYINT UNSIGNED NOT NULL,
  `forfeit` BOOLEAN DEFAULT NULL,
  FOREIGN KEY (`match_id`) REFERENCES matches(`id`),
  FOREIGN KEY (`winner`) REFERENCES teams(`id`),
  UNIQUE KEY `match_id_map_number` (`match_id`, `map_number`)
);