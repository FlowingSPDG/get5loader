CREATE TABLE `players` (
  `id` BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `team_id` BIGINT NOT NULL,
  `steam_id` varchar(40) NOT NULL,
  `name` varchar(40) NOT NULL,
  FOREIGN KEY (`team_id`) REFERENCES teams(`id`)
);