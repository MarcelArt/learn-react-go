
package api_handlers

import (
	"github.com/MarcelArt/learn-react-go/models"
	"github.com/MarcelArt/learn-react-go/repositories"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type SchoolHandler struct {
	BaseCrudHandler[models.School, models.SchoolDTO, models.SchoolPage]
	repo repositories.ISchoolRepo
}

func NewSchoolHandler(repo repositories.ISchoolRepo) *SchoolHandler {
	return &SchoolHandler{
		BaseCrudHandler: BaseCrudHandler[models.School, models.SchoolDTO, models.SchoolPage]{
			repo: repo,
			validator: validator.New(validator.WithRequiredStructEnabled()),
		},
		repo: repo,
	}
}

// Create creates a new school
// @Summary Create a new school
// @Description Create a new school
// @Tags School
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param School body models.SchoolDTO true "School data"
// @Success 201 {object} models.SchoolDTO
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /school [post]
func (h *SchoolHandler) Create(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Create(c)
}

// Read retrieves a list of schools
// @Summary Get a list of schools
// @Description Get a list of schools
// @Tags School
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "Page"
// @Param size query int false "Size"
// @Param sort query string false "Sort"
// @Param filters query string false "Filter"
// @Success 200 {array} models.SchoolPage
// @Router /school [get]
func (h *SchoolHandler) Read(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Read(c)
}

// Update updates an existing school
// @Summary Update an existing school
// @Description Update an existing school
// @Tags School
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "School ID"
// @Param School body models.SchoolDTO true "School data"
// @Success 200 {object} models.SchoolDTO
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /school/{id} [put]
func (h *SchoolHandler) Update(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Update(c)
}

// Delete deletes an existing school
// @Summary Delete an existing school
// @Description Delete an existing school
// @Tags School
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "School ID"
// @Success 200 {object} models.School
// @Failure 500 {object} string
// @Router /school/{id} [delete]
func (h *SchoolHandler) Delete(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Delete(c)
}

// GetByID retrieves a school by ID
// @Summary Get a school by ID
// @Description Get a school by ID
// @Tags School
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "School ID"
// @Success 200 {object} models.School
// @Failure 500 {object} string
// @Router /school/{id} [get]
func (h *SchoolHandler) GetByID(c *fiber.Ctx) error {
	return h.BaseCrudHandler.GetByID(c)
}
