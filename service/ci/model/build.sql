USE e5Code;
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `build`;
CREATE TABLE `build` (
    `id` varchar(255) NOT NULL,
    `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `name` varchar(255) NOT NULL,
    `project_id` varchar(255) NOT NULL,
    `building_detail` varchar(255) NOT NULL,
    `desc` varchar(255),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;
