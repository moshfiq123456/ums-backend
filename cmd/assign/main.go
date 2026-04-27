package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/moshfiq123456/ums-backend/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func exec(db *gorm.DB, sql string, args ...interface{}) {
	if err := db.Exec(sql, args...).Error; err != nil {
		log.Printf("  ⚠  %v", err)
	}
}

func assignRole(db *gorm.DB, email, roleCode string) {
	exec(db, `INSERT INTO user_roles (user_id, role_id, assigned_at)
		SELECT u.id, r.id, NOW() FROM users u, roles r
		WHERE u.email = ? AND r.code = ? ON CONFLICT DO NOTHING`, email, roleCode)
	log.Printf("  role  %-40s → %s", email, roleCode)
}

func assignPerm(db *gorm.DB, email, permCode string) {
	exec(db, `INSERT INTO user_permissions (user_id, permission_id, allow)
		SELECT u.id, p.id, true FROM users u, permissions p
		WHERE u.email = ? AND p.code = ? ON CONFLICT DO NOTHING`, email, permCode)
	log.Printf("  perm  %-40s → %s", email, permCode)
}

func addHierarchy(db *gorm.DB, parentEmail, childEmail string) {
	exec(db, `INSERT INTO user_hierarchy (parent_user_id, child_user_id, created_at)
		SELECT p.id, c.id, NOW() FROM users p, users c
		WHERE p.email = ? AND c.email = ? ON CONFLICT DO NOTHING`, parentEmail, childEmail)
	log.Printf("  hier  %s  →  %s", parentEmail, childEmail)
}

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

	// ── List current users ────────────────────────────────────────
	type row struct {
		Email  string
		Name   string
		Status string
	}
	var users []row
	db.Raw(`SELECT email, name, status FROM users ORDER BY created_at`).Scan(&users)
	log.Println("\n── Current users ──────────────────────────────")
	for _, u := range users {
		log.Printf("  %-35s %-20s %s", u.Email, u.Name, u.Status)
	}

	// ══════════════════════════════════════════════════════════════
	// ROLE ASSIGNMENTS
	// Give several users multiple roles for richer test data
	// ══════════════════════════════════════════════════════════════
	log.Println("\n── Assigning roles ────────────────────────────")

	// Bob gets an extra support_agent role
	assignRole(db, "bob@ums.com", "support_agent")

	// Charlie gets hr_manager on top of viewer
	assignRole(db, "charlie@ums.com", "hr_manager")

	// Nina gets viewer in addition to auditor
	assignRole(db, "nina@ums.com", "viewer")

	// Grace gets viewer role (she had none)
	assignRole(db, "grace@ums.com", "viewer")

	// Henry gets support_agent (he had none)
	assignRole(db, "henry@ums.com", "support_agent")

	// Sophie gets viewer on top of support_agent
	assignRole(db, "sophie@ums.com", "viewer")

	// Priya gets viewer as well
	assignRole(db, "priya@ums.com", "viewer")

	// If the user Saif exists from earlier testing, give them a role too
	assignRole(db, "nasty.munna@example.com", "viewer")

	// ══════════════════════════════════════════════════════════════
	// DIRECT USER PERMISSION OVERRIDES
	// ══════════════════════════════════════════════════════════════
	log.Println("\n── Assigning direct permissions ───────────────")

	// Henry: can read users and roles directly (no need for a full role)
	assignPerm(db, "henry@ums.com", "user:read")
	assignPerm(db, "henry@ums.com", "role:read")
	assignPerm(db, "henry@ums.com", "permission:read")

	// Grace: can generate reports
	assignPerm(db, "grace@ums.com", "report:generate")
	assignPerm(db, "grace@ums.com", "audit:read")

	// Bob: gets user:update override
	assignPerm(db, "bob@ums.com", "user:update")
	assignPerm(db, "bob@ums.com", "user:status")

	// Charlie: gets role management perms directly
	assignPerm(db, "charlie@ums.com", "role:create")
	assignPerm(db, "charlie@ums.com", "role:update")

	// James: gets audit read
	assignPerm(db, "james@ums.com", "audit:read")
	assignPerm(db, "james@ums.com", "report:generate")

	// Nina: gets hierarchy manage on top of auditor
	assignPerm(db, "nina@ums.com", "hierarchy:manage")

	// Sophie: gets user:create override
	assignPerm(db, "sophie@ums.com", "user:create")

	// Saif if exists
	assignPerm(db, "nasty.munna@example.com", "user:read")
	assignPerm(db, "nasty.munna@example.com", "role:read")

	// ══════════════════════════════════════════════════════════════
	// HIERARCHY — extend the org tree
	// ══════════════════════════════════════════════════════════════
	log.Println("\n── Adding hierarchy ───────────────────────────")

	// james now also reports to bob (cross-team dotted line)
	addHierarchy(db, "bob@ums.com", "james@ums.com")

	// omar now also under david (dotted line)
	addHierarchy(db, "david@ums.com", "omar@ums.com")

	// grace and henry under alice too (dual reporting)
	addHierarchy(db, "alice@ums.com", "grace@ums.com")
	addHierarchy(db, "alice@ums.com", "henry@ums.com")

	// charlie and frank under nina
	addHierarchy(db, "nina@ums.com", "charlie@ums.com")
	addHierarchy(db, "nina@ums.com", "frank@ums.com")

	// Saif if exists
	addHierarchy(db, "alice@ums.com", "nasty.munna@example.com")

	// ── Final counts ──────────────────────────────────────────────
	var counts struct {
		Roles       int64
		Perms       int64
		UserRoles   int64
		UserPerms   int64
		Hierarchies int64
	}
	db.Raw(`SELECT COUNT(*) FROM roles`).Scan(&counts.Roles)
	db.Raw(`SELECT COUNT(*) FROM permissions`).Scan(&counts.Perms)
	db.Raw(`SELECT COUNT(*) FROM user_roles`).Scan(&counts.UserRoles)
	db.Raw(`SELECT COUNT(*) FROM user_permissions`).Scan(&counts.UserPerms)
	db.Raw(`SELECT COUNT(*) FROM user_hierarchy`).Scan(&counts.Hierarchies)

	log.Println("\n── Summary ────────────────────────────────────")
	log.Printf("  roles            : %d", counts.Roles)
	log.Printf("  permissions      : %d", counts.Perms)
	log.Printf("  user_roles       : %d", counts.UserRoles)
	log.Printf("  user_permissions : %d", counts.UserPerms)
	log.Printf("  user_hierarchy   : %d", counts.Hierarchies)
	log.Println("\n✅ Done!")
}
