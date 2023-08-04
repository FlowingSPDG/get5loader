CREATE TABLE `matches` (
  `id` BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `user_id` BIGINT NOT NULL,
  `server_id` BIGINT NOT NULL,
  `team1_id` BIGINT NOT NULL,
  `team2_id` BIGINT NOT NULL,
  `winner` BIGINT DEFAULT NULL,
  `cancelled` BOOLEAN NOT NULL DEFAULT FALSE,
  `start_time` datetime DEFAULT NULL,
  `end_time` datetime DEFAULT NULL,
  `max_maps` int(11) NOT NULL,
  `title` varchar(60) NOT NULL,
  `skip_veto` BOOLEAN NOT NULL,
  `api_key` varchar(32) NOT NULL,
  `team1_score` int(11) NOT NULL,
  `team2_score` int(11) NOT NULL,
  `forfeit` tinyint(1) DEFAULT NULL,
  FOREIGN KEY (`user_id`) REFERENCES users(`id`),
  FOREIGN KEY (`server_id`) REFERENCES servers(`id`),
  FOREIGN KEY (`team1_id`) REFERENCES teams(`id`),
  FOREIGN KEY (`team2_id`) REFERENCES teams(`id`),
  FOREIGN KEY (`winner`) REFERENCES teams(`id`)
);