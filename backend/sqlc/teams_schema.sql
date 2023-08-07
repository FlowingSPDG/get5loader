CREATE TABLE `teams` (
  `id` VARCHAR(36) NOT NULL PRIMARY KEY,
  `user_id` VARCHAR(36) NOT NULL,
  `name` varchar(40) NOT NULL,
  `flag` varchar(4) NOT NULL,
  `logo` varchar(10) NOT NULL,
  `tag` varchar(40) NOT NULL,
  `public_team` BOOLEAN DEFAULT FALSE,
  FOREIGN KEY (`user_id`) REFERENCES users(`id`)
);