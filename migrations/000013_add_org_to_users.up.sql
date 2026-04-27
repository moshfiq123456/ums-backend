-- Add organization_id (default to the seeded org for existing rows)
ALTER TABLE users
    ADD COLUMN organization_id UUID NOT NULL
        DEFAULT '00000000-0000-0000-0000-000000000001'
        REFERENCES organizations(id) ON DELETE CASCADE,
    ADD COLUMN user_type VARCHAR(20) NOT NULL DEFAULT 'member',
    ADD COLUMN metadata  JSONB DEFAULT '{}';

-- Remove the global email unique constraint
ALTER TABLE users DROP CONSTRAINT IF EXISTS users_email_key;

-- Add per-org email uniqueness
ALTER TABLE users ADD CONSTRAINT uq_users_email_org UNIQUE (email, organization_id);

-- Drop the default after migration (org_id must be explicit going forward)
ALTER TABLE users ALTER COLUMN organization_id DROP DEFAULT;

CREATE INDEX idx_users_org_id    ON users(organization_id);
CREATE INDEX idx_users_user_type ON users(user_type);
