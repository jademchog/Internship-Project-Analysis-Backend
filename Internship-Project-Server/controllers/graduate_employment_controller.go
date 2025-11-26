package controllers

import (
	"net/http"
	"strconv"

	"api-template/services"

	"github.com/gin-gonic/gin"
)

// QueryGraduateEmployment 根据大学和年份查询就业数据
// @Summary 查询就业数据
// @Tags graduate-employment
// @Produce json
// @Param university query string true "大学名称"
// @Param year query int true "年份"
// @Success 200 {array} models.GraduateEmploymentResponse
// @Router /api/v1/graduate-employment [get]
func QueryGraduateEmployment(ctx *gin.Context) {
	university := ctx.Query("university")
	yearStr := ctx.Query("year")

	if university == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "university 参数不能为空"})
		return
	}

	if yearStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "year 参数不能为空"})
		return
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "year 必须是数字"})
		return
	}

	results, err := services.GraduateEmploymentService.QueryByUniversityAndYear(university, year)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"university": university,
		"year":       year,
		"data":       results,
	})
}
