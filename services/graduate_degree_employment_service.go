package services

import (
	"api-template/managers"
	"api-template/models"
)

var GraduateDegreeEmploymentService = &graduateDegreeEmploymentService{}

type graduateDegreeEmploymentService struct{}

// QueryByUniversityAndYear 查询指定大学和年份的就业数据（业务逻辑）
func (s *graduateDegreeEmploymentService) QueryByUniversityAndDegree(university string, degree string) ([]models.GraduateDegreeEmploymentResponse, error) {
	if university == "" || degree == "" {
		return []models.GraduateDegreeEmploymentResponse{}, nil
	}

	return managers.GraduateEmploymentManager.QueryByUniversityAndDegree(university, degree)
}
