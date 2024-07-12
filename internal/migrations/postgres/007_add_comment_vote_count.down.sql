BEGIN;

DROP TRIGGER update_comment_vote_count_trigger ON comment_votes;

DROP FUNCTION update_comment_vote_count;

ALTER TABLE comments DROP COLUMN vote_count;

COMMIT;