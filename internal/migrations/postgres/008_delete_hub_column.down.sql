BEGIN;

CREATE TYPE hub_enum AS ENUM (
	'books',
	'movies',
	'series',
	'anime',
	'gaming'
);

ALTER TABLE topics DROP CONSTRAINT IF EXISTS topics_topic_key;
ALTER TABLE topics ADD COLUMN hub hub_enum NOT NULL;
ALTER TABLE topics ADD CONSTRAINT topics_topic_hub_key UNIQUE (topic, hub);

ALTER TABLE tags DROP CONSTRAINT IF EXISTS tags_tag_key;
ALTER TABLE tags ADD COLUMN hub hub_enum NOT NULL;
ALTER TABLE tags ADD CONSTRAINT tags_tag_hub_key UNIQUE (tag, hub);

COMMIT;