package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/moshfiq123456/ums-backend/internal/config"
	"github.com/moshfiq123456/ums-backend/internal/constants"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var defaultOrgID = constants.DefaultOrgID

func hash(pw string) string {
	h, err := bcrypt.GenerateFromPassword([]byte(pw), 10)
	if err != nil {
		log.Fatal("bcrypt:", err)
	}
	return string(h)
}

func ago(d time.Duration) time.Time { return time.Now().Add(-d) }

func main() {
	_ = godotenv.Load(".env")
	cfg := config.LoadConfig()

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port,
		cfg.Database.User, cfg.Database.Password, cfg.Database.DBName,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("connect:", err)
	}
	log.Println("✅ connected")

	// ══════════════════════════════════════════════════════════════
	// DEFAULT ORGANIZATION
	// ══════════════════════════════════════════════════════════════
	db.Exec(`INSERT INTO organizations (id,name,slug,is_active,plan,settings,created_at,updated_at)
		VALUES (?,?,?,true,'free','{}',NOW(),NOW()) ON CONFLICT (id) DO NOTHING`,
		defaultOrgID, "Default Organization", "default")
	log.Println("  organization: default org seeded")

	// ══════════════════════════════════════════════════════════════
	// ROLES  (scoped to default org)
	// ══════════════════════════════════════════════════════════════
	type role struct {
		name, code, desc string
		active           bool
	}
	roles := []role{
		{"Super Admin", "super_admin", "Full unrestricted system access", true},
		{"Manager", "manager", "Manages users, roles and assignments", true},
		{"HR Manager", "hr_manager", "Handles user onboarding and status changes", true},
		{"Support Agent", "support_agent", "Read access to users and tickets", true},
		{"Viewer", "viewer", "Read-only access across the system", true},
		{"Auditor", "auditor", "Access to audit logs and reports", true},
		{"Guest", "guest", "Minimal temporary access", false},
	}
	for _, r := range roles {
		db.Exec(`INSERT INTO roles (organization_id,name,code,description,is_active,created_at,updated_at)
			VALUES (?,?,?,?,?,NOW(),NOW()) ON CONFLICT DO NOTHING`,
			defaultOrgID, r.name, r.code, r.desc, r.active)
	}
	log.Printf("  roles: %d seeded", len(roles))

	// ══════════════════════════════════════════════════════════════
	// PERMISSIONS
	// ══════════════════════════════════════════════════════════════
	perms := [][]string{
		// User management
		{"user:create", "Create User", "Create a new user account"},
		{"user:read", "Read User", "View user profile and details"},
		{"user:update", "Update User", "Edit user information"},
		{"user:delete", "Delete User", "Permanently delete a user"},
		{"user:status", "Update User Status", "Change status: active / inactive / blocked"},
		// Role management
		{"role:create", "Create Role", "Create a new role"},
		{"role:read", "Read Role", "View role details"},
		{"role:update", "Update Role", "Edit role name and description"},
		{"role:delete", "Delete Role", "Delete a role"},
		{"role:status", "Update Role Status", "Activate or deactivate a role"},
		// Permission management
		{"permission:create", "Create Permission", "Create a new permission"},
		{"permission:read", "Read Permission", "View permission details"},
		{"permission:update", "Update Permission", "Edit permission information"},
		{"permission:delete", "Delete Permission", "Delete a permission"},
		// User-role assignments
		{"user_role:assign", "Assign Role to User", "Assign one or more roles to a user"},
		{"user_role:remove", "Remove Role from User", "Remove roles from a user"},
		{"user_role:read", "Read User Roles", "View roles assigned to a user"},
		// User-permission overrides
		{"user_permission:assign", "Assign Permission to User", "Grant a permission directly to a user"},
		{"user_permission:remove", "Remove Permission from User", "Revoke a direct permission from a user"},
		{"user_permission:read", "Read User Permissions", "View direct permissions of a user"},
		// Hierarchy
		{"hierarchy:manage", "Manage Hierarchy", "Assign/remove parent-child user relationships"},
		{"hierarchy:read", "Read Hierarchy", "View the user org hierarchy"},
		// Reports & Audit
		{"report:generate", "Generate Reports", "Create and export system reports"},
		{"audit:read", "Read Audit Logs", "View the system audit trail"},
		{"audit:export", "Export Audit Logs", "Download audit log exports"},
	}
	for _, p := range perms {
		db.Exec(`INSERT INTO permissions (organization_id,code,name,description,created_at,updated_at)
			VALUES (?,?,?,?,NOW(),NOW()) ON CONFLICT DO NOTHING`, defaultOrgID, p[0], p[1], p[2])
	}

	log.Printf("  permissions: %d seeded", len(perms))

	// ══════════════════════════════════════════════════════════════
	// ROLE → PERMISSION MAPPING
	// ══════════════════════════════════════════════════════════════
	assign := func(roleCode string, codes []string) {
		for _, pc := range codes {
			db.Exec(`INSERT INTO role_permissions (role_id,permission_id)
				SELECT r.id,p.id FROM roles r, permissions p
				WHERE r.organization_id=? AND r.code=? AND p.code=? ON CONFLICT DO NOTHING`,
				defaultOrgID, roleCode, pc)
		}
	}

	var allCodes []string
	for _, p := range perms {
		allCodes = append(allCodes, p[0])
	}

	assign("super_admin", allCodes)

	assign("manager", []string{
		"user:create", "user:read", "user:update", "user:status",
		"role:read", "permission:read",
		"user_role:assign", "user_role:remove", "user_role:read",
		"user_permission:assign", "user_permission:remove", "user_permission:read",
		"hierarchy:manage", "hierarchy:read",
		"report:generate",
	})

	assign("hr_manager", []string{
		"user:create", "user:read", "user:update", "user:status",
		"role:read", "user_role:assign", "user_role:read",
		"hierarchy:read",
	})

	assign("support_agent", []string{
		"user:read", "role:read", "permission:read",
		"user_role:read", "hierarchy:read",
	})

	assign("viewer", []string{
		"user:read", "role:read", "permission:read",
		"user_role:read", "user_permission:read", "hierarchy:read",
	})

	assign("auditor", []string{
		"user:read", "role:read", "permission:read",
		"audit:read", "audit:export", "report:generate",
	})

	assign("guest", []string{
		"user:read",
	})

	log.Println("  role_permissions seeded")

	// ══════════════════════════════════════════════════════════════
	// USERS  (all under default org)
	// ══════════════════════════════════════════════════════════════
	type seedUser struct {
		name, email, password, phone, status string
		roles                                []string
		createdAgo                           time.Duration
	}

	users := []seedUser{
		// Admins
		{"Super Admin", "admin@ums.com", "Admin@123456", "+8801700000001", "active", []string{"super_admin"}, 90 * 24 * time.Hour},
		// Managers
		{"Alice Rahman", "alice@ums.com", "Manager@123456", "+8801700000002", "active", []string{"manager"}, 80 * 24 * time.Hour},
		{"David Kim", "david@ums.com", "Manager@123456", "+8801700000010", "active", []string{"manager", "hr_manager"}, 75 * 24 * time.Hour},
		// HR
		{"Priya Sharma", "priya@ums.com", "HrPass@123456", "+8801700000005", "active", []string{"hr_manager"}, 60 * 24 * time.Hour},
		{"Omar Hassan", "omar@ums.com", "HrPass@123456", "+8801700000006", "active", []string{"hr_manager"}, 55 * 24 * time.Hour},
		// Support
		{"Sophie Turner", "sophie@ums.com", "Support@123456", "+8801700000007", "active", []string{"support_agent"}, 45 * 24 * time.Hour},
		{"James Lee", "james@ums.com", "Support@123456", "+8801700000008", "active", []string{"support_agent"}, 40 * 24 * time.Hour},
		// Auditor
		{"Nina Patel", "nina@ums.com", "Audit@123456", "+8801700000009", "active", []string{"auditor"}, 30 * 24 * time.Hour},
		// Viewers
		{"Bob Chen", "bob@ums.com", "Viewer@123456", "+8801700000003", "active", []string{"viewer"}, 70 * 24 * time.Hour},
		{"Charlie Davis", "charlie@ums.com", "User@123456", "+8801700000004", "active", []string{"viewer"}, 65 * 24 * time.Hour},
		// Inactive / Blocked
		{"Eve Wilson", "eve@ums.com", "User@123456", "+8801700000011", "inactive", []string{}, 20 * 24 * time.Hour},
		{"Frank Bruno", "frank@ums.com", "User@123456", "+8801700000012", "blocked", []string{"guest"}, 10 * 24 * time.Hour},
		// No role
		{"Grace Liu", "grace@ums.com", "User@123456", "+8801700000013", "active", []string{}, 5 * 24 * time.Hour},
		{"Henry Park", "henry@ums.com", "User@123456", "+8801700000014", "active", []string{}, 2 * 24 * time.Hour},
	}

	for _, u := range users {
		createdAt := ago(u.createdAgo)
		db.Exec(`INSERT INTO users (organization_id,name,email,password_hash,phone,status,created_at,updated_at)
			VALUES (?,?,?,?,?,?,?,?) ON CONFLICT (organization_id, email) DO NOTHING`,
			defaultOrgID, u.name, u.email, hash(u.password), u.phone, u.status, createdAt, createdAt)

		for _, rc := range u.roles {
			db.Exec(`INSERT INTO user_roles (user_id,role_id,assigned_at)
				SELECT u.id,r.id,? FROM users u, roles r
				WHERE u.email=? AND r.organization_id=? AND r.code=? ON CONFLICT DO NOTHING`,
				createdAt.Add(time.Hour), u.email, defaultOrgID, rc)
		}
	}
	log.Printf("  users: %d seeded", len(users))

	// ══════════════════════════════════════════════════════════════
	// USER PERMISSIONS  (direct overrides)
	// ══════════════════════════════════════════════════════════════
	type userPerm struct{ email, permCode string }
	userPerms := []userPerm{
		{"bob@ums.com", "audit:read"},
		{"charlie@ums.com", "report:generate"},
		{"sophie@ums.com", "user:update"},
		{"grace@ums.com", "user:read"},
		{"grace@ums.com", "role:read"},
	}
	for _, up := range userPerms {
		db.Exec(`INSERT INTO user_permissions (user_id,permission_id,allow)
			SELECT u.id,p.id,true FROM users u, permissions p
			WHERE u.email=? AND p.code=? ON CONFLICT DO NOTHING`, up.email, up.permCode)
	}
	log.Printf("  user_permissions: %d seeded", len(userPerms))

	// ══════════════════════════════════════════════════════════════
	// USER HIERARCHY
	// ══════════════════════════════════════════════════════════════
	//
	//  admin
	//  ├── alice (manager)
	//  │   ├── priya (hr)
	//  │   │   └── omar (hr)
	//  │   ├── sophie (support)
	//  │   │   └── james (support)
	//  │   └── bob (viewer)
	//  │       └── charlie (viewer)
	//  └── david (manager)
	//      ├── nina (auditor)
	//      ├── eve
	//      ├── grace
	//      └── henry
	//
	hierarchies := [][2]string{
		{"admin@ums.com", "alice@ums.com"},
		{"admin@ums.com", "david@ums.com"},
		{"alice@ums.com", "priya@ums.com"},
		{"alice@ums.com", "sophie@ums.com"},
		{"alice@ums.com", "bob@ums.com"},
		{"priya@ums.com", "omar@ums.com"},
		{"sophie@ums.com", "james@ums.com"},
		{"bob@ums.com", "charlie@ums.com"},
		{"david@ums.com", "nina@ums.com"},
		{"david@ums.com", "eve@ums.com"},
		{"david@ums.com", "grace@ums.com"},
		{"david@ums.com", "henry@ums.com"},
	}
	for _, h := range hierarchies {
		db.Exec(`INSERT INTO user_hierarchy (parent_user_id,child_user_id,created_at)
			SELECT p.id,c.id,NOW() FROM users p, users c
			WHERE p.email=? AND c.email=? ON CONFLICT DO NOTHING`, h[0], h[1])
	}
	log.Printf("  user_hierarchy: %d entries seeded", len(hierarchies))

	// ══════════════════════════════════════════════════════════════
	// AUDIT LOGS
	// ══════════════════════════════════════════════════════════════
	type auditEntry struct {
		actorEmail, action, entity, entityEmail string
		daysAgo                                 int
	}
	audits := []auditEntry{
		{"admin@ums.com", "user_created", "user", "alice@ums.com", 80},
		{"admin@ums.com", "role_assigned", "user", "alice@ums.com", 80},
		{"admin@ums.com", "user_created", "user", "david@ums.com", 75},
		{"admin@ums.com", "role_assigned", "user", "david@ums.com", 75},
		{"alice@ums.com", "user_created", "user", "priya@ums.com", 60},
		{"alice@ums.com", "role_assigned", "user", "priya@ums.com", 60},
		{"alice@ums.com", "user_created", "user", "sophie@ums.com", 45},
		{"alice@ums.com", "user_created", "user", "bob@ums.com", 70},
		{"david@ums.com", "user_created", "user", "nina@ums.com", 30},
		{"david@ums.com", "user_status_changed", "user", "eve@ums.com", 5},
		{"admin@ums.com", "user_status_changed", "user", "frank@ums.com", 2},
		{"admin@ums.com", "role_created", "role", "auditor", 35},
		{"admin@ums.com", "role_created", "role", "hr_manager", 62},
		{"admin@ums.com", "permission_created", "permission", "audit:read", 35},
		{"admin@ums.com", "permission_created", "permission", "report:generate", 35},
		{"alice@ums.com", "user_permission_assigned", "user", "bob@ums.com", 25},
		{"david@ums.com", "hierarchy_assigned", "user", "nina@ums.com", 29},
		{"admin@ums.com", "hierarchy_assigned", "user", "alice@ums.com", 79},
	}
	for _, a := range audits {
		createdAt := ago(time.Duration(a.daysAgo) * 24 * time.Hour)
		db.Exec(`INSERT INTO audit_logs (user_id,action,entity,entity_id,created_at)
			SELECT u.id,?,?,?,? FROM users u WHERE u.email=?`,
			a.action, a.entity, a.entityEmail, createdAt, a.actorEmail)
	}
	log.Printf("  audit_logs: %d entries seeded", len(audits))

	// ══════════════════════════════════════════════════════════════
	// LOGIN SESSIONS  (historical — already logged out)
	// ══════════════════════════════════════════════════════════════
	type session struct {
		email, ip, ua string
		daysAgo        int
	}
	sessions := []session{
		{"admin@ums.com", "192.168.0.101", "Mozilla/5.0 Chrome/123 Linux", 10},
		{"admin@ums.com", "192.168.0.101", "Mozilla/5.0 Chrome/123 Linux", 5},
		{"alice@ums.com", "192.168.0.102", "Mozilla/5.0 Firefox/124 Linux", 8},
		{"alice@ums.com", "192.168.0.102", "Mozilla/5.0 Firefox/124 Linux", 3},
		{"david@ums.com", "192.168.0.105", "Mozilla/5.0 Chrome/123 Windows", 7},
		{"bob@ums.com", "192.168.0.103", "Mozilla/5.0 Safari/17 macOS", 6},
		{"priya@ums.com", "192.168.0.106", "Mozilla/5.0 Chrome/123 Linux", 4},
		{"sophie@ums.com", "192.168.0.107", "Mozilla/5.0 Edge/123 Windows", 2},
		{"nina@ums.com", "192.168.0.109", "Mozilla/5.0 Chrome/123 Linux", 1},
		{"frank@ums.com", "10.0.0.55", "curl/7.88.0", 15},
	}
	for _, s := range sessions {
		loggedIn := ago(time.Duration(s.daysAgo)*24*time.Hour + 8*time.Hour)
		loggedOut := loggedIn.Add(2 * time.Hour)
		sessionID := uuid.New().String()
		db.Exec(`INSERT INTO login_sessions
				(id,user_id,refresh_token_hash,refresh_expires_at,ip_address,user_agent,logged_in_at,logged_out_at)
			SELECT ?,u.id,'historical-session-hash',?,?,?,?,? FROM users u WHERE u.email=?`,
			sessionID,
			loggedIn.Add(7*24*time.Hour),
			s.ip, s.ua, loggedIn, loggedOut,
			s.email,
		)
	}
	log.Printf("  login_sessions: %d seeded", len(sessions))

	// ══════════════════════════════════════════════════════════════
	// SUMMARY
	// ══════════════════════════════════════════════════════════════
	log.Println("")
	log.Println("✅ Seed complete!")
	log.Println("")
	log.Println("  Default Org ID:", defaultOrgID)
	log.Println("")
	log.Println("  Credentials:")
	log.Println("  admin@ums.com    / Admin@123456   → super_admin")
	log.Println("  alice@ums.com    / Manager@123456 → manager")
	log.Println("  david@ums.com    / Manager@123456 → manager + hr_manager")
	log.Println("  priya@ums.com    / HrPass@123456  → hr_manager")
	log.Println("  omar@ums.com     / HrPass@123456  → hr_manager")
	log.Println("  sophie@ums.com   / Support@123456 → support_agent")
	log.Println("  james@ums.com    / Support@123456 → support_agent")
	log.Println("  nina@ums.com     / Audit@123456   → auditor")
	log.Println("  bob@ums.com      / Viewer@123456  → viewer")
	log.Println("  charlie@ums.com  / User@123456    → viewer")
	log.Println("  eve@ums.com      / User@123456    → (inactive)")
	log.Println("  frank@ums.com    / User@123456    → guest (blocked)")
	log.Println("  grace@ums.com    / User@123456    → (no role)")
	log.Println("  henry@ums.com    / User@123456    → (no role)")
}
