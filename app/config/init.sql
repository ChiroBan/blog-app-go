create table posts (
  id serial not null unique,
  title varchar(64),
  content text,
  primary key(id)
);


insert into posts(title, content)
values
    ('First Post', 'Once upon a time on earth'),
    ('Second Post', 'Computer science is an interesting field');