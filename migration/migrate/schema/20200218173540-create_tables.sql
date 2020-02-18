
-- +migrate Up
ALTER TABLE `round_stats`
  ADD `winner` int(11) DEFAULT NULL ,
  ADD `winner_side` varchar(40) DEFAULT NULL ;

-- +migrate Down

ALTER TABLE `round_stats`
  drop `winner` ,
  drop `winner_side` ;