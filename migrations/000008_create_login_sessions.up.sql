
-- =========================
-- LOGIN SESSIONS
-- =========================
CREATE TABLE login_sessions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    ip_address VARCHAR(45),
    user_agent TEXT,
    logged_in_at TIMESTAMP NOT NULL DEFAULT NOW(),
    logged_out_at TIMESTAMP,
    CONSTRAINT fk_login_sessions_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- LOGIN SESSIONS
CREATE INDEX idx_login_sessions_user_id ON login_sessions(user_id);
CREATE INDEX idx_login_sessions_logged_in_at ON login_sessions(logged_in_at);