create database chat;

use chat;

create table if not exists accounts (
    id BIGINT UNSIGNED AUTO_INCREMENT,
    account_name text NOT NULL,
    email varchar(256) UNIQUE,
    password text NOT NULL,
    primary key (id),
    UNIQUE (email),
    CHECK (LENGTH(password) > 0 AND LENGTH(account_name) > 0)
);

create index idx_accounts on accounts (id, email);

create table if not exists messages (
    message_id BIGINT UNSIGNED AUTO_INCREMENT,
    message_from BIGINT UNSIGNED,
    message_to BIGINT UNSIGNED,
    content text NOT NULL,
    created_at DATETIME(3) NOT NULL,
    update_at DATETIME(3) NOT NULL,

    primary key (message_id),
    constraint foreign key (message_from) references accounts(id) ON UPDATE CASCADE ON DELETE SET NULL,
    constraint foreign key (message_to) references accounts(id),
    CHECK (LENGTH(content) > 0)
);

create index idx_messages on messages (message_id);

create table if not exists user_friend
(
    id     BIGINT UNSIGNED AUTO_INCREMENT,
    uid1   BIGINT UNSIGNED,
    uid2   BIGINT UNSIGNED,
    STATUS ENUM ('REQ_UID1', 'REQ_UID2', 'FRIEND') NOT NULL,

    primary key (id),
    CHECK ( uid1 < uid2 )
);