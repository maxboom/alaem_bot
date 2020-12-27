create table users
(
    id            int auto_increment
        primary key,
    username      varchar(255)         not null,
    is_authorized tinyint(1) default 0 not null
);

create table alarms
(
    id      int auto_increment
        primary key,
    user_id int          not null,
    time    time         not null,
    text    varchar(255) null,
    constraint alarms_users_id_fk
        foreign key (user_id) references users (id)
            on update cascade on delete cascade
)
    collate = utf8mb4_unicode_ci;

