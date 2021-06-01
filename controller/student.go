package controller

import (
	"net/http"

	"github.com/Mersock/golang-echo-postgres-restful-api/model"
	"github.com/Mersock/golang-echo-postgres-restful-api/storage"
	"github.com/labstack/echo/v4"
)

func GetStudents(c echo.Context) error {
	students, _ := GetRepoStudents()
	return c.JSON(http.StatusOK, students)
}

func GetRepoStudents() ([]model.Student, error) {
	db := storage.GetDBInstance()
	students := []model.Student{}

	if err := db.Find(&students).Error; err != nil {
		return students, err
	}

	return students, nil
}
