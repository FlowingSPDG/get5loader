CREATE TABLE `players` (
  `id` VARCHAR(36) NOT NULL PRIMARY KEY,
  `team_id` VARCHAR(36) NOT NULL,
  `steam_id` BIGINT UNSIGNED NOT NULL,
  `name` varchar(40) NOT NULL,
  FOREIGN KEY (`team_id`) REFERENCES teams(`id`)
);