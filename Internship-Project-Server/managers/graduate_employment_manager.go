package managers

import (
	"api-template/models"

	"gorm.io/gorm"
)

type graduateEmploymentManager struct {
	db *gorm.DB
}

func (m *graduateEmploymentManager) QueryByUniversityAndYear(university string, year int) ([]models.GraduateEmploymentResponse, error) {
	var results []models.GraduateEmploymentResponse

	err := m.db.Session(&gorm.Session{AllowGlobalUpdate: true}).
		Table("survey").
		Where("university = ? AND year = ?", university, year).
		Select("school", "employment_rate_overall").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}
	return results, nil
}
