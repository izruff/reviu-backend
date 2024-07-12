BEGIN;

ALTER TABLE comments ADD COLUMN vote_count integer default 0;

CREATE FUNCTION update_comment_vote_count() RETURNS TRIGGER AS $$
  BEGIN
    IF (TG_OP = 'INSERT') THEN
      UPDATE comments
      SET vote_count = vote_count + NEW.up::int
      WHERE id = NEW.comment_id;
    ELSIF (TG_OP = 'DELETE') THEN
      UPDATE comments
      SET vote_count = vote_count - OLD.up::int
      WHERE id = OLD.comment_id;
    ELSIF (TG_OP = 'UPDATE') THEN
      UPDATE comments
      SET vote_count = vote_count + NEW.up::int - OLD.up::int
      WHERE id = NEW.comment_id;
    END IF;
    RETURN NULL;
  END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_comment_vote_count_trigger
AFTER INSERT ON comment_votes
FOR EACH ROW
EXECUTE FUNCTION update_comment_vote_count();

COMMIT;