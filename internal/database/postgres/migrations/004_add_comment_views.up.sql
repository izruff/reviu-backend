BEGIN;

CREATE TABLE IF NOT EXISTS comment_views (
  comment_id integer references comments(id) not null,
  user_id integer references users(id) not null,
  created_at timestamptz default current_timestamp,
  primary key (comment_id, user_id)
);

COMMIT;