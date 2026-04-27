ALTER TABLE users
    DROP COLUMN IF EXISTS organization_id,
    DROP COLUMN IF EXISTS user_type,
    DROP COLUMN IF EXISTS metadata;

ALTER TABLE users ADD CONSTRAINT users_email_key UNIQUE (email);
