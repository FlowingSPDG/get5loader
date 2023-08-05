CREATE TABLE `game_servers` (
  `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `user_id` BIGINT NOT NULL,
  `ip` varchar(32) NOT NULL,
  `port` SMALLINT UNSIGNED NOT NULL DEFAULT 27015,
  `rcon_password` varchar(32) NOT NULL,
  `display_name` varchar(32) NOT NULL,
  `is_public` BOOLEAN NOT NULL DEFAULT FALSE,
  `status` TINYINT NOT NULL,
  FOREIGN KEY (`user_id`) REFERENCES users(`id`),
  UNIQUE KEY `unique_ip_port` (`ip`, `port`)
);