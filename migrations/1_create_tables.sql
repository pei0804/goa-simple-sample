-- +migrate Up

CREATE TABLE accounts
(
	id bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
	name varchar(30) NOT NULL COMMENT '名前',
	email text NOT NULL COMMENT 'メールアドレス',
	created_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '作成日',
	updated_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '更新日',
  deleted_at datetime COMMENT '削除日',
	PRIMARY KEY (id)
) COMMENT = 'アカウント';


CREATE TABLE bottles
(
	id bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
	name varchar(30) COMMENT 'ボトル名',
	quantity int unsigned DEFAULT 0 NOT NULL COMMENT '数量',
	created_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '作成日',
	updated_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '更新日',
  deleted_at datetime COMMENT '削除日',
	account_id bigint unsigned NOT NULL COMMENT 'アカウントID',
	PRIMARY KEY (id)
) COMMENT = 'ボトル';

CREATE TABLE categories
(
	id bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
	name varchar(30) NOT NULL COMMENT 'カテゴリ名',
	created_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '作成日',
	updated_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '更新日',
  deleted_at datetime COMMENT '削除日',
	PRIMARY KEY (id)
) COMMENT = 'カテゴリ';

CREATE TABLE bottle_categories
(
	id bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
	created_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '作成日',
	updated_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '更新日',
  deleted_at datetime COMMENT '削除日',
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

--
-- テーブルのデータのダンプ `accounts`
--

INSERT INTO `accounts` (`id`, `name`, `email`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'ユーザー1', 'example1@gmail.com', '2017-05-10 20:04:31', '2017-05-10 20:04:31', NULL),
(2, 'ユーザー2', 'example2@gmail.com', '2017-05-10 20:04:31', '2017-05-10 20:04:31', NULL);

--
-- テーブルのデータのダンプ `bottles`
--

INSERT INTO `bottles` (`id`, `name`, `quantity`, `created_at`, `updated_at`, `deleted_at`, `account_id`) VALUES
(1, '赤ワイン1', 20, '2017-05-10 20:05:21', '2017-05-10 20:05:21', NULL, 1),
(2, '赤ワイン2', 10, '2017-05-10 20:05:21', '2017-05-10 20:05:21', NULL, 1),
(3, '白ワイン1', 20, '2017-05-10 20:06:02', '2017-05-10 20:06:02', NULL, 2),
(4, '白ワイン2', 10, '2017-05-10 20:06:02', '2017-05-10 20:06:02', NULL, 2);


--
-- テーブルのデータのダンプ `categories`
--

INSERT INTO `categories` (`id`, `name`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, '赤', '2017-05-10 20:06:26', '2017-05-10 20:06:26', NULL),
(2, '白', '2017-05-10 20:06:26', '2017-05-10 20:06:26', NULL),
(3, 'ワイン', '2017-05-10 20:06:38', '2017-05-10 20:06:38', NULL),
(4, 'フランス', '2017-05-10 20:06:38', '2017-05-10 20:06:38', NULL);


--
-- テーブルのデータのダンプ `bottle_categories`
--

INSERT INTO `bottle_categories` (`id`, `created_at`, `updated_at`, `deleted_at`, `bottle_id`, `category_id`) VALUES
(1, '2017-05-10 20:10:02', '2017-05-10 20:10:02', NULL, 1, 1),
(2, '2017-05-10 20:10:02', '2017-05-10 20:10:02', NULL, 1, 3),
(3, '2017-05-10 20:10:57', '2017-05-10 20:10:57', NULL, 2, 2),
(4, '2017-05-10 20:10:57', '2017-05-10 20:10:57', NULL, 2, 3);

-- +migrate Down

DROP TABLE IF EXISTS bottle_categories;
DROP TABLE IF EXISTS bottles;
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS categories;