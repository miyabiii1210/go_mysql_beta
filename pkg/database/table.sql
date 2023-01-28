CREATE TABLE `user` (
	`uid` INT (10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`name` VARCHAR (20) CHARACTER SET utf8mb4 COLLATE `utf8mb4_bin` NOT NULL,
	`email` VARCHAR (50) CHARACTER SET utf8mb4 COLLATE `utf8mb4_bin` NOT NULL,
	`tel_number` VARCHAR (16) CHARACTER SET utf8mb4 COLLATE `utf8mb4_bin` NOT NULL,
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (`uid`)
) AUTO_INCREMENT = 10000001 COMMENT = 'ユーザ管理テーブル';

CREATE TABLE `notification` (
	`nid` INT (10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`uid` INT (10) UNSIGNED NOT NULL,
	`json` json,
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	INDEX `uid` (`uid`),
	PRIMARY KEY (`notice_id`)
) AUTO_INCREMENT = 1001 COMMENT = '通知管理テーブル';