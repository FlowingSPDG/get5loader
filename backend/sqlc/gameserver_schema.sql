CREATE TABLE `game_servers` (
  `id` VARCHAR(36) NOT NULL PRIMARY KEY,
  `user_id` VARCHAR(36) NOT NULL,
  `ip` BINARY(4) NOT NULL,
  `port` SMALLINT UNSIGNED NOT NULL DEFAULT 27015,
  `rcon_password` VARCHAR(32) NOT NULL,
  `display_name` VARCHAR(32) NOT NULL,
  `is_public` BOOLEAN NOT NULL DEFAULT FALSE,
  `status` TINYINT NOT NULL,
  FOREIGN KEY (`user_id`) REFERENCES users(`id`),
  UNIQUE KEY `unique_ip_port` (`ip`, `port`),
  CHECK(is_ipv4(ip))
);