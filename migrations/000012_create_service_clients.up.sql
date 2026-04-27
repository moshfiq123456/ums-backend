CREATE TABLE service_clients (
    id                  UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    organization_id     UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    name                VARCHAR(100) NOT NULL,
    client_id           VARCHAR(100) UNIQUE NOT NULL,
    client_secret_hash  TEXT NOT NULL,
    allowed_scopes      TEXT[] DEFAULT '{}',
    is_active           BOOLEAN NOT NULL DEFAULT true,
    created_at          TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_service_clients_org_id    ON service_clients(organization_id);
CREATE INDEX idx_service_clients_client_id ON service_clients(client_id);
