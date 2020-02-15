
-- +migrate Up
SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;


CREATE TABLE `alembic_version` (
  `version_num` varchar(32) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `game_server` (
  `id` int(11) NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  `in_use` tinyint(1) DEFAULT NULL,
  `ip_string` varchar(32) DEFAULT NULL,
  `port` int(11) DEFAULT NULL,
  `rcon_password` varchar(32) DEFAULT NULL,
  `display_name` varchar(32) DEFAULT NULL,
  `public_server` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `map_stats` (
  `id` int(11) NOT NULL,
  `match_id` int(11) DEFAULT NULL,
  `map_number` int(11) DEFAULT NULL,
  `map_name` varchar(64) DEFAULT NULL,
  `start_time` datetime DEFAULT NULL,
  `end_time` datetime DEFAULT NULL,
  `winner` int(11) DEFAULT NULL,
  `team1_score` int(11) DEFAULT NULL,
  `team2_score` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `match` (
  `id` int(11) NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  `server_id` int(11) DEFAULT NULL,
  `team1_id` int(11) DEFAULT NULL,
  `team2_id` int(11) DEFAULT NULL,
  `winner` int(11) DEFAULT NULL,
  `cancelled` tinyint(1) DEFAULT NULL,
  `start_time` datetime DEFAULT NULL,
  `end_time` datetime DEFAULT NULL,
  `max_maps` int(11) DEFAULT NULL,
  `title` varchar(60) DEFAULT NULL,
  `skip_veto` tinyint(1) DEFAULT NULL,
  `api_key` varchar(32) DEFAULT NULL,
  `veto_mappool` varchar(500) DEFAULT NULL,
  `team1_score` int(11) DEFAULT NULL,
  `team2_score` int(11) DEFAULT NULL,
  `team1_string` varchar(32) DEFAULT NULL,
  `team2_string` varchar(32) DEFAULT NULL,
  `forfeit` tinyint(1) DEFAULT NULL,
  `plugin_version` varchar(32) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `player_stats` (
  `id` int(11) NOT NULL,
  `match_id` int(11) DEFAULT NULL,
  `map_id` int(11) DEFAULT NULL,
  `team_id` int(11) DEFAULT NULL,
  `steam_id` varchar(40) DEFAULT NULL,
  `name` varchar(40) DEFAULT NULL,
  `kills` int(11) DEFAULT NULL,
  `deaths` int(11) DEFAULT NULL,
  `roundsplayed` int(11) DEFAULT NULL,
  `assists` int(11) DEFAULT NULL,
  `flashbang_assists` int(11) DEFAULT NULL,
  `teamkills` int(11) DEFAULT NULL,
  `suicides` int(11) DEFAULT NULL,
  `headshot_kills` int(11) DEFAULT NULL,
  `damage` int(11) DEFAULT NULL,
  `bomb_plants` int(11) DEFAULT NULL,
  `bomb_defuses` int(11) DEFAULT NULL,
  `v1` int(11) DEFAULT NULL,
  `v2` int(11) DEFAULT NULL,
  `v3` int(11) DEFAULT NULL,
  `v4` int(11) DEFAULT NULL,
  `v5` int(11) DEFAULT NULL,
  `k1` int(11) DEFAULT NULL,
  `k2` int(11) DEFAULT NULL,
  `k3` int(11) DEFAULT NULL,
  `k4` int(11) DEFAULT NULL,
  `k5` int(11) DEFAULT NULL,
  `firstdeath_Ct` int(11) DEFAULT NULL,
  `firstdeath_t` int(11) DEFAULT NULL,
  `firstkill_ct` int(11) DEFAULT NULL,
  `firstkill_t` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `team` (
  `id` int(11) NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  `name` varchar(40) DEFAULT NULL,
  `flag` varchar(4) DEFAULT NULL,
  `logo` varchar(10) DEFAULT NULL,
  `auths` blob,
  `tag` varchar(40) DEFAULT NULL,
  `public_team` tinyint(1) DEFAULT '1'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `user` (
  `id` int(11) NOT NULL,
  `steam_id` varchar(40) DEFAULT NULL,
  `name` varchar(40) DEFAULT NULL,
  `admin` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


ALTER TABLE `game_server`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

ALTER TABLE `map_stats`
  ADD PRIMARY KEY (`id`),
  ADD KEY `match_id` (`match_id`),
  ADD KEY `winner` (`winner`);

ALTER TABLE `match`
  ADD PRIMARY KEY (`id`),
  ADD KEY `server_id` (`server_id`),
  ADD KEY `team1_id` (`team1_id`),
  ADD KEY `team2_id` (`team2_id`),
  ADD KEY `user_id` (`user_id`),
  ADD KEY `winner` (`winner`);

ALTER TABLE `player_stats`
  ADD PRIMARY KEY (`id`),
  ADD KEY `map_id` (`map_id`),
  ADD KEY `match_id` (`match_id`),
  ADD KEY `team_id` (`team_id`);

ALTER TABLE `team`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`),
  ADD KEY `ix_team_public_team` (`public_team`);

ALTER TABLE `user`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `steam_id` (`steam_id`);


ALTER TABLE `game_server`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

ALTER TABLE `map_stats`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

ALTER TABLE `match`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

ALTER TABLE `player_stats`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

ALTER TABLE `team`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

ALTER TABLE `user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;


ALTER TABLE `map_stats`
  ADD CONSTRAINT `map_stats_ibfk_1` FOREIGN KEY (`match_id`) REFERENCES `match` (`id`),
  ADD CONSTRAINT `map_stats_ibfk_2` FOREIGN KEY (`winner`) REFERENCES `team` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;


-- +migrate Down
DROP TABLE IF EXISTS `alembic_version`,`game_server`,`map_stats`,`match`,`player_stats`,`team`,`user`;