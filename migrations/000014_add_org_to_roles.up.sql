ALTER TABLE roles
    ADD COLUMN organization_id UUID NOT NULL
        DEFAULT '00000000-0000-0000-0000-000000000001'
        REFERENCES organizations(id) ON DELETE CASCADE;

-- Drop global unique constraints
ALTER TABLE roles DROP CONSTRAINT IF EXISTS roles_name_key;
ALTER TABLE roles DROP CONSTRAINT IF EXISTS roles_code_key;

-- Add per-org unique constraints
ALTER TABLE roles ADD CONSTRAINT uq_roles_name_org UNIQUE (name, organization_id);
ALTER TABLE roles ADD CONSTRAINT uq_roles_code_org UNIQUE (code, organization_id);

-- Drop the default after migration
ALTER TABLE roles ALTER COLUMN organization_id DROP DEFAULT;

CREATE INDEX idx_roles_org_id ON roles(organization_id);
