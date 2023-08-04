CREATE TABLE `users` (
  `id`   BIGINT  NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `steam_id` varchar(40) NOT NULL UNIQUE,
  `name` varchar(40) NOT NULL,
  `admin` BOOLEAN NOT NULL
);