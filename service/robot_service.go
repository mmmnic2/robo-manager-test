package service

import (
	"device-manager/dto"
	"device-manager/model"
	"device-manager/repository"
	"math"
	"net/http"
	"strings"
	"time"
)

type RobotService interface {
	GetRobots(request dto.RobotRequest) (dto.RobotResponse, error)
}

type robotService struct {
	robotRepository repository.RobotRepository
}

func (r *robotService) GetRobots(request dto.RobotRequest) (dto.RobotResponse, error) {
	if request.PageNumber == 0 {
		request.PageNumber = 1
	}
	if request.Limit == 0 {
		request.Limit = 10
	}
	query := r.robotRepository.DB().Model(&model.Robot{})
	var response dto.RobotResponse
	if len(request.Filters) > 0 {
		for key, values := range request.Filters {
			switch key {
			case "manufacturing_date":
				switch v := values.(type) {
				case []interface{}:
					if len(v) == 2 {
						startDate, errStart := time.Parse("2006-01-02", v[0].(string))
						endDate, errEnd := time.Parse("2006-01-02", v[1].(string))
						if errStart == nil && errEnd == nil {
							query = query.Where(key+" BETWEEN ? AND ?", startDate, endDate)
						}
					} else if len(v) == 1 {
						date, err := time.Parse("2006-01-02", v[0].(string))
						if err == nil {
							query = query.Where(key+"= ?", date)
						}
					} else if len(v) > 2 {
						return response, &dto.ErrorResponse{
							StatusCode: http.StatusBadRequest,
							Message:    "invalid data type for " + key}
					}
				case string:
					date, err := time.Parse("2006-01-02", v)
					if err == nil {
						query = query.Where(key+" = ?", date)
					}
				default:
					return response, &dto.ErrorResponse{
						StatusCode: http.StatusBadRequest,
						Message:    "invalid data type for " + key}
				}

			default:
				switch v := values.(type) {
				case []interface{}:
					var lowerValues []string
					for _, value := range v {
						lowerValues = append(lowerValues, strings.ToLower(value.(string)))
					}
					query = query.Where("LOWER("+key+") IN (?)", lowerValues)

				case string:
					query = query.Where("LOWER("+key+") LIKE ?", "%"+strings.ToLower(v)+"%")
				case int, float32, float64:
					query = query.Where(key+" = ?", v)
				default:
					return response, &dto.ErrorResponse{
						StatusCode: http.StatusBadRequest,
						Message:    "Unsupported filter type for key" + key}

				}
			}
		}
	}

	if request.Sorting.Column != "" {
		sortDirection := "asc"
		if request.Sorting.Order == "desc" {
			sortDirection = "desc"
		}
		query = query.Order(request.Sorting.Column + " " + sortDirection)
	}
	if err := query.Count(&response.TotalCount).Error; err != nil {
		return response, &dto.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to count robots."}
	}

	if int64(request.Limit) > response.TotalCount && request.PageNumber > 1 {
		request.PageNumber = 1
	}
	offset := (request.PageNumber - 1) * request.Limit
	if err := query.Offset(offset).Limit(request.Limit).Find(&response.Robots).Error; err != nil {
		return response, &dto.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to find robots"}
	}

	response.TotalPages = int(math.Ceil(float64(response.TotalCount) / float64(request.Limit)))
	response.CurrentPage = request.PageNumber
	return response, nil

}

func NewRobotService(robotRepository repository.RobotRepository) RobotService {
	return &robotService{robotRepository: robotRepository}
}
