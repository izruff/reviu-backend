BEGIN;

DROP TRIGGER update_post_vote_count_trigger ON post_votes;

DROP FUNCTION update_post_vote_count;

ALTER TABLE posts DROP COLUMN vote_count;

COMMIT;