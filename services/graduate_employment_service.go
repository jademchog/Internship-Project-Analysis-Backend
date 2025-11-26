package services

import (
	"api-template/managers"
	"api-template/models"
)

var GraduateEmploymentService = &graduateEmploymentService{}

type graduateEmploymentService struct{}


// QueryByUniversityAndYear 查询指定大学和年份的就业数据（业务逻辑）
func (s *graduateEmploymentService) QueryByUniversityAndYear(university string, year int) ([]models.GraduateEmploymentResponse, error) {
	if university == "" || year <= 0 {
		return []models.GraduateEmploymentResponse{}, nil
	}

	return managers.GraduateEmploymentManager.QueryByUniversityAndYear(university, year)
}

