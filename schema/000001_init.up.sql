CREATE TABLE users
(
  id serial not null unique,
  name varchar(255) not null unique,
  email varchar(255) not null unique,
  about varchar default '',
  password varchar(255) not null
);

CREATE TABLE posts
(
  id serial not null unique,
  title varchar(255) not null,
  description varchar not null,
  image_url varchar(255) not null,
  crypto_currency varchar(255) not null,
  created_at timestamp not null,
  owner_id int references users (id) on delete cascade not null
);

CREATE TABLE comments
(
  id serial not null unique,
  post_id int references posts (id) on delete cascade not null,
  comment varchar(255) not null,
  owner_id int references users (id) on delete cascade not null
);

CREATE TABLE replies
(
  id serial not null unique,
  comment_id int references comments (id) on delete cascade not null,
  reply varchar(255) not null,
  owner_id int references users (id) on delete cascade not null
);

CREATE TABLE followers
(
  user_id int references users (id) on delete cascade not null,
  follower_id int not null
);

CREATE TABLE likes
(
  post_id int references posts (id) on delete cascade not null,
  user_id int references users (id) on delete cascade not null
);

CREATE TABLE watchlist
(
  watchlist_id serial not null unique,
  user_id int references users (id) on delete cascade not null,
  watchlist varchar(255) not null,
  crypto_currency varchar(255) not null
);