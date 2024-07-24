BEGIN;

ALTER TABLE topics DROP CONSTRAINT (
    SELECT constraint_name
    FROM information_schema.table_constraints
    WHERE table_name = 'topics' AND constraint_type = 'UNIQUE';
);
ALTER TABLE topics DROP COLUMN IF EXISTS hub;
ALTER TABLE topics ADD CONSTRAINT topics_topic_key UNIQUE topic;

ALTER TABLE tags DROP CONSTRAINT (
    SELECT constraint_name
    FROM information_schema.table_constraints
    WHERE table_name = 'tags' AND constraint_type = 'UNIQUE';
);
ALTER TABLE tags DROP COLUMN IF EXISTS hub;
ALTER TABLE tags ADD CONSTRAINT tags_tag_key UNIQUE tag;

DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'hub_enum') THEN
        DROP TYPE hub_enum;
    END IF;
END$$;

COMMIT;