
package repositories

import (
	"github.com/MarcelArt/learn-react-go/models"
	"gorm.io/gorm"
)

const schoolPageQuery = "-- Write your query here --"

type ISchoolRepo interface {
	IBaseCrudRepo[models.School, models.SchoolDTO, models.SchoolPage]
}

type SchoolRepo struct {
	BaseCrudRepo[models.School, models.SchoolDTO, models.SchoolPage]
}

func NewSchoolRepo(db *gorm.DB) *SchoolRepo {
	return &SchoolRepo{
		BaseCrudRepo: BaseCrudRepo[models.School, models.SchoolDTO, models.SchoolPage]{
			db:        db,
			pageQuery: schoolPageQuery,
		},
	}
}
