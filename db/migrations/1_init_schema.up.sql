CREATE TABLE tweets (
    id int not null AUTO_INCREMENT,
    createdAt datetime default current_timestamp,
    username varchar(24) not null,
    text varchar(280) not null,
    primary key (id)
);

CREATE TABLE users (
    username varchar(24) not null,
    password varchar(24) not null,
    createdAt datetime default current_timestamp,
    email varchar(50) not null,
    firstName varchar(50) not null,
    lastName varchar(50) not null,
    birth date default null,
    primary key (username)
);

CREATE TABLE followers (
    follower int not null,
    publisher int not null,
    primary key (follower, publisher),
    constraint fk_follower foreign key (follower) references users(username) on delete cascade,
    constraint fk_publisher foreign key (publisher) references users(username) on delete cascade
);