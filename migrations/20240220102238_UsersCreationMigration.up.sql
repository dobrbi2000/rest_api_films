CREATE TABLE users (
    id bigserial not null primary key,
    login varchar not null unique,
    password varchar not null   
);

CREATE TABLE films (
    id bigserial not null primary key,
    title varchar(150) not null unique,
    description varchar(1000) not null,
    year smallint not null,
    rating smallint not null CHECK (rating between 1 and 10),
    actors jsonb
);

CREATE TABLE actors (
    id bigserial not null primary key,
    name varchar not null unique,
    gender varchar not null,
    birth_date varchar not null,
    films_ids integer[]
);


    
 



