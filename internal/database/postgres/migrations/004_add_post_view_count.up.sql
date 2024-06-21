BEGIN;

ALTER TABLE posts ADD COLUMN view_count integer default 0;

CREATE FUNCTION incr_post_view_count() RETURNS TRIGGER AS $$
  BEGIN
    UPDATE posts
    SET view_count = view_count + 1
    WHERE id = NEW.post_id;
    RETURN NEW;
  END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER incr_post_view_count_trigger
AFTER INSERT ON post_views
FOR EACH ROW
EXECUTE FUNCTION incr_post_view_count();

COMMIT;