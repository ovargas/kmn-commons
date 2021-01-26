create table dummy
(
    dummy_id bigint auto_increment primary key,
    title    varchar(64) not null,
    constraint dummy_uk unique (dummy_id)
);
