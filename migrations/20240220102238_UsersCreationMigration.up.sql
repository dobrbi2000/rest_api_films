CREATE TABLE users (
    user_id bigserial not null primary key,
    login varchar not null unique,
    password varchar not null   
);

CREATE TABLE films (
    film_id bigserial not null primary key,
    title varchar(150) not null unique,
    description varchar(1000) not null,
    year smallint not null,
    rating smallint not null CHECK (rating between 1 and 10)
);

CREATE TABLE actors (
    actor_id bigserial not null primary key,
    name varchar not null unique,
    gender varchar not null,
    birth_date varchar not null
);

CREATE TABLE filmactor (
    film_id bigint not null,
    actor_id bigint not null,
    primary key (film_id, actor_id),
    foreign key (film_id) references films(film_id),
    foreign key (actor_id) references actors(actor_id)
);



    
 



