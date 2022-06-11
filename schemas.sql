CREATE TABLE login (
    id serial PRIMARY KEY,
    hash varchar(100) not null,
    email text UNIQUE NOT NULL
);

CREATE TABLE users (
    id serial PRIMARY KEY,
    name varchar(100),
    email text UNIQUE NOT NULL,
    entries BIGINT default 0,
    joined TIMESTAMP NOT NULL,
    age int default 0,
    pet varchar(100) default 'none'
);

insert into users (name, email, entries, joined, age, pet) values
('hello', 'hello@gmail.com', 5, '2021-12-26', 23, 'cat'),
('Badrri', 'badrri1995@gmail.com', 0, '2022-06-11', 27, 'dog'),
('Lysander Stark', 'colonelstark@gmail.com', 2, '2022-1-6', 63, 'tiger'),
('Masey Ferguson', 'mferguson@gmail.com', 15, '2011-2-12', 44, 'lizard'),
('John Watson', 'drwatson@gmail.com', 51, '2001-11-3', 33, 'mouse'),
('Sherlock Holmes', 'holmes@gmail.com', 22, '2021-10-21', 34, 'bunny');

-- encoded password = password
insert into login (hash, email) values('$2a$04$lvK6HpXxNy5kkHthZpyvp.XDdw6iZVuevbJTmi4cHCbK5KbvF1fuy', 'test@gmail.com');
insert into users (name, email, entries, joined, age, pet) values('testtest', 'test@gmail.com', 4, '2021-12-26', 21, 'monkey');