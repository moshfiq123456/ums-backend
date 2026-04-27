DELETE FROM permissions WHERE service IN ('ecommerce', 'admin');
ALTER TABLE permissions DROP COLUMN IF EXISTS service;
