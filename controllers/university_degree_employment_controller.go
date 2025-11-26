package controllers

import (
	"net/http"

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
func QueryUniDegreeEmployment(ctx *gin.Context) {
	university := ctx.Query("university")
	degree := ctx.Query("degree")

	if university == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "university 参数不能为空"})
		return
	}

	if degree == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "degree 参数不能为空"})
		return
	}

	results, err := services.GraduateDegreeEmploymentService.QueryByUniversityAndDegree(university, degree)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"university": university,
		"degree":     degree,
		"data":       results,
	})
}
