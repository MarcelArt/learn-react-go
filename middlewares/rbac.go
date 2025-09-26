package middlewares

import (
	"fmt"
	"github.com/MarcelArt/learn-react-go/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type RBACMiddleware struct {
	db *gorm.DB
}

func NewRBACMiddleware(db *gorm.DB) *RBACMiddleware {
	return &RBACMiddleware{db: db}
}

// RequireAuth ensures the user is authenticated
func (rm *RBACMiddleware) RequireAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Check if user is already authenticated by previous middleware
		if c.Locals("user") == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authentication required",
			})
		}
		return c.Next()
	}
}

// RequireRole ensures the user has at least one of the specified roles
func (rm *RBACMiddleware) RequireRole(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := rm.getCurrentUser(c)
		if user == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "User not authenticated",
			})
		}

		// Check if user has any of the required roles
		userRoles, err := rm.getUserRoles(user.ID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to check user roles",
			})
		}

		// Check if user has any of the required roles
		for _, requiredRole := range roles {
			for _, userRole := range userRoles {
				if userRole.Name == requiredRole {
					// User has required role, proceed
					c.Locals("userRoles", userRoles)
					return c.Next()
				}
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": fmt.Sprintf("Access denied. Required roles: %v", roles),
		})
	}
}

// RequirePermission ensures the user has the specified permission
func (rm *RBACMiddleware) RequirePermission(resource string, action string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := rm.getCurrentUser(c)
		if user == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "User not authenticated",
			})
		}

		// Check if user has the required permission
		hasPermission, err := rm.checkUserPermission(user.ID, resource, action)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to check permissions",
			})
		}

		if !hasPermission {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": fmt.Sprintf("Access denied. Required permission: %s:%s", resource, action),
			})
		}

		return c.Next()
	}
}

// RequireAnyPermission ensures the user has at least one of the specified permissions
func (rm *RBACMiddleware) RequireAnyPermission(permissions ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := rm.getCurrentUser(c)
		if user == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "User not authenticated",
			})
		}

		// Check if user has any of the required permissions
		for _, permission := range permissions {
			hasPermission, err := rm.checkUserPermission(user.ID, permission, "")
			if err != nil {
				continue
			}
			if hasPermission {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": fmt.Sprintf("Access denied. Required any of permissions: %v", permissions),
		})
	}
}

// ResourceOwnership ensures user can only access their own resources or has admin role
func (rm *RBACMiddleware) ResourceOwnership(resourceType string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := rm.getCurrentUser(c)
		if user == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "User not authenticated",
			})
		}

		// Check if user is admin (can access all resources)
		userRoles, err := rm.getUserRoles(user.ID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to check user roles",
			})
		}

		// Check if user has admin role
		for _, role := range userRoles {
			if role.Name == "admin" || role.Name == "super_admin" {
				return c.Next()
			}
		}

		// Get resource ID from params
		resourceID := c.Params("id")
		if resourceID == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Resource ID not provided",
			})
		}

		// Check if user owns the resource
		// This will be implemented based on resource type
		ownsResource, err := rm.checkResourceOwnership(user.ID, resourceType, resourceID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to check resource ownership",
			})
		}

		if !ownsResource {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Access denied. You don't own this resource",
			})
		}

		return c.Next()
	}
}

// Helper methods
func (rm *RBACMiddleware) getCurrentUser(c *fiber.Ctx) *models.User {
	claims := c.Locals("user").(*jwt.MapClaims)
	userID := uint((*claims)["userId"].(float64))

	var user models.User
	if err := rm.db.First(&user, userID).Error; err != nil {
		return nil
	}

	return &user
}

func (rm *RBACMiddleware) getUserRoles(userID uint) ([]models.Role, error) {
	var user models.User
	if err := rm.db.Preload("Roles").First(&user, userID).Error; err != nil {
		return nil, err
	}
	return user.Roles, nil
}

func (rm *RBACMiddleware) checkUserPermission(userID uint, resource string, action string) (bool, error) {
	var user models.User
	if err := rm.db.Preload("Roles.Permissions").First(&user, userID).Error; err != nil {
		return false, err
	}

	// Check all user roles for the required permission
	for _, role := range user.Roles {
		if !role.IsActive {
			continue
		}

		for _, permission := range role.Permissions {
			if !permission.IsActive {
				continue
			}

			// If action is empty, only check resource
			if action == "" {
				if permission.Resource == resource {
					return true, nil
				}
				continue
			}

			// Check both resource and action
			if permission.Resource == resource && permission.Action == action {
				return true, nil
			}
		}
	}

	return false, nil
}

func (rm *RBACMiddleware) checkResourceOwnership(userID uint, resourceType string, resourceID string) (bool, error) {
	// This will be implemented based on specific resource types
	// For now, return false (ownership check fails)
	return false, nil
}

// GetCurrentUserRoles returns the current user's roles from context
func GetCurrentUserRoles(c *fiber.Ctx) []models.Role {
	if roles, ok := c.Locals("userRoles").([]models.Role); ok {
		return roles
	}
	return []models.Role{}
}

// HasRole checks if current user has a specific role
func HasRole(c *fiber.Ctx, role string) bool {
	roles := GetCurrentUserRoles(c)
	for _, r := range roles {
		if r.Name == role {
			return true
		}
	}
	return false
}