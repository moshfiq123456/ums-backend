ALTER TABLE roles DROP COLUMN IF EXISTS organization_id;
ALTER TABLE roles ADD CONSTRAINT roles_name_key UNIQUE (name);
ALTER TABLE roles ADD CONSTRAINT roles_code_key UNIQUE (code);
