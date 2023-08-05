CREATE TABLE `teams` (
  `id` BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `user_id` BIGINT NOT NULL,
  `name` varchar(40) NOT NULL,
  `flag` varchar(4) NOT NULL,
  `logo` varchar(10) NOT NULL,
  `tag` varchar(40) NOT NULL,
  `public_team` BOOLEAN DEFAULT FALSE,
  FOREIGN KEY (`user_id`) REFERENCES users(`id`)
);