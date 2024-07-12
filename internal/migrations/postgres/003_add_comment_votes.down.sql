BEGIN;

DROP TABLE IF EXISTS comment_votes;

ALTER TABLE post_votes RENAME TO votes;

COMMIT;