package handlers

import (
	"goscraper/backend/src/helpers"
	"goscraper/backend/src/types"
)

func GetCourses(token string) (*types.CourseResponse, error) {
	scraper := helpers.NewCoursePage(token)
	course, err := scraper.GetCourses()

	return course, err
}
