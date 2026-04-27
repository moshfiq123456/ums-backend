CREATE TABLE organizations (
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name        VARCHAR(200) NOT NULL,
    slug        VARCHAR(100) UNIQUE NOT NULL,
    domain      VARCHAR(200),
    logo_url    VARCHAR(500),
    is_active   BOOLEAN NOT NULL DEFAULT true,
    plan        VARCHAR(50) NOT NULL DEFAULT 'free',
    settings    JSONB DEFAULT '{}',
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMP
);

CREATE INDEX idx_organizations_slug     ON organizations(slug);
CREATE INDEX idx_organizations_domain   ON organizations(domain);
CREATE INDEX idx_organizations_deleted  ON organizations(deleted_at);

-- Insert a default organization for existing data
INSERT INTO organizations (id, name, slug, plan)
VALUES ('00000000-0000-0000-0000-000000000001', 'Default Organization', 'default', 'free');
