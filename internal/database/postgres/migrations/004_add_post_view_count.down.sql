BEGIN;

DROP TRIGGER incr_post_view_count_trigger ON post_views;

DROP FUNCTION incr_post_view_count;

ALTER TABLE posts DROP COLUMN view_count;

COMMIT;