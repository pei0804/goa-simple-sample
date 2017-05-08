-- +migrate Up

CREATE TABLE accounts
(
	id bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
	name varchar(30) NOT NULL COMMENT '名前',
	email text NOT NULL COMMENT 'メールアドレス',
	created_at datetime NOT NULL COMMENT '作成日時',
	updated_at datetime NOT NULL COMMENT '更新日時',
	deletead_at datetime COMMENT '削除日時',
	PRIMARY KEY (id)
) COMMENT = 'アカウント';


CREATE TABLE bottles
(
	id bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
	name varchar(30) COMMENT 'ボトル名',
	quantity int unsigned DEFAULT 0 NOT NULL COMMENT '数量',
	account_id bigint unsigned NOT NULL COMMENT 'アカウントID',
	PRIMARY KEY (id)
) COMMENT = 'ボトル';

CREATE TABLE categories
(
	id bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
	name varchar(30) NOT NULL COMMENT 'カテゴリ名',
	created_at datetime COMMENT '作成日時',
	updated_at datetime COMMENT '更新日時',
	deletead_at datetime COMMENT '削除日時',
	PRIMARY KEY (id)
) COMMENT = 'カテゴリ';

CREATE TABLE bottle_categories
(
	id bigint unsigned NOT NULL COMMENT 'ID',
	created_at datetime COMMENT '作成日時',
	updated_at datetime COMMENT '更新日時',
	deletead_at datetime COMMENT '削除日時',
	bottle_id bigint unsigned NOT NULL COMMENT 'ボトルID',
	category_id bigint unsigned NOT NULL COMMENT 'カテゴリID',
	PRIMARY KEY (id)
) COMMENT = 'ボトルカテゴリ';


ALTER TABLE bottles
	ADD FOREIGN KEY (account_id)
	REFERENCES accounts (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE bottle_categories
	ADD FOREIGN KEY (bottle_id)
	REFERENCES bottles (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE bottle_categories
	ADD FOREIGN KEY (category_id)
	REFERENCES categories (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;

-- +migrate Down

DROP TABLE IF EXISTS bottle_categories;
DROP TABLE IF EXISTS bottles;
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS categories;