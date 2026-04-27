-- Audit logs
ALTER TABLE audit_logs
    ADD COLUMN organization_id UUID REFERENCES organizations(id) ON DELETE SET NULL,
    ADD COLUMN service         VARCHAR(50) DEFAULT 'ums',
    ADD COLUMN ip_address      VARCHAR(45);

UPDATE audit_logs SET organization_id = '00000000-0000-0000-0000-000000000001';

CREATE INDEX idx_audit_logs_org_id  ON audit_logs(organization_id);
CREATE INDEX idx_audit_logs_service ON audit_logs(service);

-- Login sessions
ALTER TABLE login_sessions
    ADD COLUMN organization_id UUID REFERENCES organizations(id) ON DELETE CASCADE,
    ADD COLUMN device_info     JSONB DEFAULT '{}';

UPDATE login_sessions SET organization_id = '00000000-0000-0000-0000-000000000001';

CREATE INDEX idx_login_sessions_org_id ON login_sessions(organization_id);
