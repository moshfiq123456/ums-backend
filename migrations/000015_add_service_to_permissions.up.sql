ALTER TABLE permissions
    ADD COLUMN service VARCHAR(50) NOT NULL DEFAULT 'ums';

-- Tag existing permissions with 'ums'
UPDATE permissions SET service = 'ums';

-- Seed ecommerce & admin permission codes
INSERT INTO permissions (code, name, description, service) VALUES
    -- Ecommerce
    ('order:read',        'Read Orders',           'View customer orders',             'ecommerce'),
    ('order:update',      'Update Orders',         'Update order status',              'ecommerce'),
    ('order:delete',      'Delete Orders',         'Cancel/delete orders',             'ecommerce'),
    ('product:create',    'Create Products',       'Add new products',                 'ecommerce'),
    ('product:read',      'Read Products',         'View products',                    'ecommerce'),
    ('product:update',    'Update Products',       'Edit existing products',           'ecommerce'),
    ('product:delete',    'Delete Products',       'Remove products',                  'ecommerce'),
    ('category:manage',   'Manage Categories',     'Create/edit/delete categories',    'ecommerce'),
    ('inventory:manage',  'Manage Inventory',      'Update stock levels',              'ecommerce'),
    ('payment:read',      'Read Payments',         'View payment transactions',        'ecommerce'),
    ('discount:manage',   'Manage Discounts',      'Create/edit discount codes',       'ecommerce'),
    -- Admin panel
    ('customer:read',     'Read Customers',        'View customer profiles',           'admin'),
    ('customer:update',   'Update Customers',      'Edit customer information',        'admin'),
    ('customer:delete',   'Delete Customers',      'Remove customer accounts',         'admin'),
    ('analytics:read',    'Read Analytics',        'View reports and dashboards',      'admin'),
    ('settings:manage',   'Manage Settings',       'Configure store settings',         'admin')
ON CONFLICT (code) DO NOTHING;

CREATE INDEX idx_permissions_service ON permissions(service);
