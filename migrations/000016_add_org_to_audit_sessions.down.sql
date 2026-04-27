ALTER TABLE audit_logs DROP COLUMN IF EXISTS organization_id, DROP COLUMN IF EXISTS service, DROP COLUMN IF EXISTS ip_address;
ALTER TABLE login_sessions DROP COLUMN IF EXISTS organization_id, DROP COLUMN IF EXISTS device_info;
