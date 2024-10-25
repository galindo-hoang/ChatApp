create table if not exists accounts (
    id BIGINT UNSIGNED AUTO_INCREMENT,
    account_name text,
    email varchar(256),
    hash_password text,
    primary key (id),
    UNIQUE (email)
);

create table if not exists messages (
    message_id BIGINT UNSIGNED AUTO_INCREMENT,
    message_from BIGINT UNSIGNED,
    message_to BIGINT UNSIGNED,
    content text,
    created_at timestamp,
    primary key (message_id),
    constraint foreign key (message_from) references accounts(id),
    constraint foreign key (message_to) references accounts(id)
);

-- MIGRATE
DROP TABLE IF EXISTS messages;

DROP TABLE IF EXISTS accounts;