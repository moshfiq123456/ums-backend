-- =========================
-- USER HIERARCHY
-- =========================
CREATE TABLE user_hierarchy (
    id BIGSERIAL PRIMARY KEY,
    parent_user_id UUID NOT NULL,
    child_user_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_user_hierarchy_parent FOREIGN KEY (parent_user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_user_hierarchy_child FOREIGN KEY (child_user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT uq_user_hierarchy UNIQUE (parent_user_id, child_user_id)
);

-- USER HIERARCHY INDEXES
CREATE INDEX idx_user_hierarchy_parent ON user_hierarchy(parent_user_id);
CREATE INDEX idx_user_hierarchy_child ON user_hierarchy(child_user_id);

