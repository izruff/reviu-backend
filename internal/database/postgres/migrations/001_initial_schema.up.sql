BEGIN;

-- define hub names
CREATE TYPE hub_enum AS ENUM (
  'books',
  'movies',
  'series',
  'anime',
  'gaming'
);

-- initialize users table
CREATE TABLE IF NOT EXISTS users (
  id integer primary key generated always as identity,
  email varchar unique not null,
  password_hash varchar not null,
  mod_role boolean not null,
  username varchar unique not null,
  nickname varchar,
  about varchar,
  created_at timestamptz default current_timestamp
);

-- initialize topics table
CREATE TABLE IF NOT EXISTS topics (
  id integer primary key generated always as identity,
  topic varchar not null,
  hub hub_enum not null,
  description varchar,
  created_at timestamptz default current_timestamp,
  unique (topic, hub)
);

-- initialize tags table
CREATE TABLE IF NOT EXISTS tags (
  id integer primary key generated always as identity,
  tag varchar not null,
  hub hub_enum not null,
  created_at timestamptz default current_timestamp,
  unique (tag, hub)
);

-- initialize posts table
CREATE TABLE IF NOT EXISTS posts (
  id integer primary key generated always as identity,
  title varchar,
  content varchar,
  author_id integer references users(id) not null,
  topic_id integer references topics(id) not null,
  created_at timestamptz default current_timestamp,
  updated_at timestamptz,
  -- moderator only; null if not deleted
  deleted_at timestamptz,
  reason_for_deletion varchar,
  moderator_id integer references users(id)
);

-- initialize comments table
CREATE TABLE IF NOT EXISTS comments (
  id integer primary key generated always as identity,
  content varchar,
  author_id integer references users(id) not null,
  post_id integer references posts(id) not null,
  parent_comment_id integer references topics(id),
  created_at timestamptz default current_timestamp,
  updated_at timestamptz,
  -- moderator only; null if not deleted
  deleted_at timestamptz,
  reason_for_deletion varchar,
  moderator_id integer references users(id)
);

-- initialize votes table
CREATE TABLE IF NOT EXISTS votes (
  up boolean not null,
  post_id integer references posts(id) not null,
  user_id integer references users(id) not null,
  created_at timestamptz default current_timestamp,
  primary key (post_id, user_id)
);

-- initialize tagged posts table
CREATE TABLE IF NOT EXISTS tagged_posts (
  post_id integer references posts(id) not null,
  tag_id integer references tags(id) not null,
  primary key (post_id, tag_id)
);

-- initialize bookmarks table
CREATE TABLE IF NOT EXISTS bookmarks (
  post_id integer references posts(id) not null,
  user_id integer references users(id) not null,
  created_at timestamptz default current_timestamp,
  primary key (post_id, user_id)
);

-- initialize relations table
CREATE TABLE IF NOT EXISTS relations (
  follower_id integer references users(id) not null,
  following_id integer references users(id) not null,
  created_at timestamptz default current_timestamp,
  primary key (follower_id, following_id),
  check (follower_id <> following_id)
);

-- initialize subscriptions table
CREATE TABLE IF NOT EXISTS subscriptions (
  topic_id integer references topics(id) not null,
  user_id integer references users(id) not null,
  created_at timestamptz default current_timestamp,
  primary key (topic_id, user_id)
);

-- initialize ban history table
CREATE TABLE IF NOT EXISTS ban_history (
  start_time timestamptz default current_timestamp,
  end_time timestamptz not null,
  reason varchar,
  user_id integer references users(id) not null,
  moderator_id integer references users(id) not null,
  primary key (start_time, user_id),
  check (start_time < end_time)
);

COMMIT;