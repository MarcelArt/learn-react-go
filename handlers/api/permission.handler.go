
package api_handlers

import (
	"github.com/MarcelArt/learn-react-go/models"
	"github.com/MarcelArt/learn-react-go/repositories"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type PermissionHandler struct {
	BaseCrudHandler[models.Permission, models.PermissionDTO, models.PermissionPage]
	repo repositories.IPermissionRepo
}

func NewPermissionHandler(repo repositories.IPermissionRepo) *PermissionHandler {
	return &PermissionHandler{
		BaseCrudHandler: BaseCrudHandler[models.Permission, models.PermissionDTO, models.PermissionPage]{
			repo: repo,
			validator: validator.New(validator.WithRequiredStructEnabled()),
		},
		repo: repo,
	}
}

// Create creates a new permission
// @Summary Create a new permission
// @Description Create a new permission
// @Tags Permission
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Permission body models.PermissionDTO true "Permission data"
// @Success 201 {object} models.PermissionDTO
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /permission [post]
func (h *PermissionHandler) Create(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Create(c)
}

// Read retrieves a list of permissions
// @Summary Get a list of permissions
// @Description Get a list of permissions
// @Tags Permission
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "Page"
// @Param size query int false "Size"
// @Param sort query string false "Sort"
// @Param filters query string false "Filter"
// @Success 200 {array} models.PermissionPage
// @Router /permission [get]
func (h *PermissionHandler) Read(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Read(c)
}

// Update updates an existing permission
// @Summary Update an existing permission
// @Description Update an existing permission
// @Tags Permission
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Permission ID"
// @Param Permission body models.PermissionDTO true "Permission data"
// @Success 200 {object} models.PermissionDTO
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /permission/{id} [put]
func (h *PermissionHandler) Update(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Update(c)
}

// Delete deletes an existing permission
// @Summary Delete an existing permission
// @Description Delete an existing permission
// @Tags Permission
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Permission ID"
// @Success 200 {object} models.Permission
// @Failure 500 {object} string
// @Router /permission/{id} [delete]
func (h *PermissionHandler) Delete(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Delete(c)
}

// GetByID retrieves a permission by ID
// @Summary Get a permission by ID
// @Description Get a permission by ID
// @Tags Permission
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Permission ID"
// @Success 200 {object} models.Permission
// @Failure 500 {object} string
// @Router /permission/{id} [get]
func (h *PermissionHandler) GetByID(c *fiber.Ctx) error {
	return h.BaseCrudHandler.GetByID(c)
}
