package handlers

import (
	"goscraper/backend/src/helpers"
	"goscraper/backend/src/types"
)

func GetAttendance(token string) (*types.AttendanceResponse, error) {
	scraper := helpers.NewAcademicsFetch(token)
	attendance, err := scraper.GetAttendance()

	return attendance, err

}
