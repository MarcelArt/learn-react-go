package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type TenancyMiddleware struct {
	db *gorm.DB
}

func NewTenancyMiddleware(db *gorm.DB) *TenancyMiddleware {
	return &TenancyMiddleware{db: db}
}

func (tm *TenancyMiddleware) Tenancy() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Skip tenancy check for public routes
		if tm.isPublicRoute(c) {
			return c.Next()
		}

		// Get school ID from JWT token
		schoolID := tm.getSchoolIDFromToken(c)
		if schoolID == 0 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "School ID not found in token",
			})
		}

		// Set school ID in context for use in handlers
		c.Locals("schoolID", schoolID)

		// Apply school scope to database queries
		tm.applySchoolScope(c)

		return c.Next()
	}
}

func (tm *TenancyMiddleware) isPublicRoute(c *fiber.Ctx) bool {
	publicRoutes := map[string]bool{
		"/api/auth/login":        true,
		"/api/auth/register":     true,
		"/api/auth/refresh":      true,
		"/api/auth/register-school": true,
		"/api/schools/public":     true,
		"/":                      true,
		"/swagger":               true,
	}

	path := c.Path()
	method := c.Method()

	// Allow OPTIONS requests (CORS preflight)
	if method == "OPTIONS" {
		return true
	}

	// Check if route is in public routes
	return publicRoutes[path]
}

func (tm *TenancyMiddleware) getSchoolIDFromToken(c *fiber.Ctx) uint {
	user := c.Locals("user")
	if user == nil {
		return 0
	}

	claims := user.(*jwt.MapClaims)
	schoolID, ok := (*claims)["schoolId"].(float64)
	if !ok {
		return 0
	}

	return uint(schoolID)
}

func (tm *TenancyMiddleware) applySchoolScope(c *fiber.Ctx) {
	schoolID := c.Locals("schoolID").(uint)

	// This will be used by repositories to scope queries
	c.Locals("schoolScope", func(db *gorm.DB) *gorm.DB {
		return db.Where("school_id = ?", schoolID)
	})
}

// SuperAdminMiddleware bypasses tenancy checks for super admins
func (tm *TenancyMiddleware) SuperAdminMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user")
		if user == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "User not authenticated",
			})
		}

		claims := user.(*jwt.MapClaims)
		_, ok := (*claims)["userId"].(float64)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid user ID in token",
			})
		}

		// Check if user is super admin (school_id is null for super admins)
		schoolID, hasSchoolID := (*claims)["schoolId"]
		if !hasSchoolID || schoolID == nil {
			// User is super admin, allow access
			c.Locals("isSuperAdmin", true)
			return c.Next()
		}

		// User is not super admin
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Super admin access required",
		})
	}
}

// GetSchoolScope returns the school scope from context
func GetSchoolScope(c *fiber.Ctx) func(*gorm.DB) *gorm.DB {
	if scope, ok := c.Locals("schoolScope").(func(*gorm.DB) *gorm.DB); ok {
		return scope
	}
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}

// IsSuperAdmin checks if the current user is a super admin
func IsSuperAdmin(c *fiber.Ctx) bool {
	return c.Locals("isSuperAdmin") == true
}