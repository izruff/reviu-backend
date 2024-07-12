BEGIN;

ALTER TABLE posts ADD COLUMN vote_count integer default 0;

CREATE FUNCTION update_post_vote_count() RETURNS TRIGGER AS $$
  BEGIN
    IF (TG_OP = 'INSERT') THEN
      UPDATE posts
      SET vote_count = vote_count + NEW.up::int
      WHERE id = NEW.post_id;
    ELSIF (TG_OP = 'DELETE') THEN
      UPDATE posts
      SET vote_count = vote_count - OLD.up::int
      WHERE id = OLD.post_id;
    ELSIF (TG_OP = 'UPDATE') THEN
      UPDATE posts
      SET vote_count = vote_count + NEW.up::int - OLD.up::int
      WHERE id = NEW.post_id;
    END IF;
    RETURN NULL;
  END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_post_vote_count_trigger
AFTER INSERT ON post_votes
FOR EACH ROW
EXECUTE FUNCTION update_post_vote_count();

COMMIT;