BEGIN;

CREATE TABLE IF NOT EXISTS post_views (
  post_id integer references posts(id) not null,
  user_id integer references users(id) not null,
  created_at timestamptz default current_timestamp,
  primary key (post_id, user_id)
);

COMMIT;