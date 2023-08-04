CREATE TABLE `game_servers` (
  `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `user_id` BIGINT NOT NULL,
  `in_use` BOOLEAN NOT NULL DEFAULT FALSE,
  `ip` varchar(32) NOT NULL,
  `port` int(11) NOT NULL DEFAULT 27015,
  `rcon_password` varchar(32) NOT NULL,
  `display_name` varchar(32) NOT NULL,
  `is_public` BOOLEAN NOT NULL DEFAULT FALSE,
  FOREIGN KEY (`user_id`) REFERENCES users(`id`)
);