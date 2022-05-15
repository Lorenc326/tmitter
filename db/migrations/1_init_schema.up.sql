BEGIN;

CREATE TABLE tweets (
    id int not null AUTO_INCREMENT,
    createdAt timestamp default current_timestamp,
    username varchar(24) not null,
    text varchar(280) not null,
    primary key (id)
);

CREATE TABLE users (
    username varchar(24) not null,
    password char(60) not null,
    createdAt timestamp default current_timestamp,
    email varchar(50) not null,
    firstName varchar(50) not null,
    lastName varchar(50) not null,
    birth date default null,
    primary key (username)
);

CREATE TABLE followers (
    follower varchar(24) not null,
    publisher varchar(24) not null,
    primary key (follower, publisher),
    constraint fk_follower foreign key (follower) references users(username) on delete cascade,
    constraint fk_publisher foreign key (publisher) references users(username) on delete cascade
);

COMMIT;