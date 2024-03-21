package utils

import (
	"fmt"
	"strings"
)

func ErrorValider(field, tag string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("Поле %s обязательно для заполнения", field)
	case "email":
		return fmt.Sprintf("Поле '%s' должно содержать адрес почты", field)
	case "min":
		if strings.EqualFold(field, "Username") {
			return fmt.Sprintf("Ник не должен быть меньше 3-х и больше 16-ти")
		}
		return fmt.Sprintf("Поле %s имеет не подходящий размер", field)
	case "max":
		if strings.EqualFold(field, "Username") {
			return fmt.Sprintf("Ник не должен быть меньше 3-х и больше 16-ти")
		}
		return fmt.Sprintf("Поле %s имеет не подходящий размер", field)
	default:
		return fmt.Sprintf("Поле '%s' имеет ошибку", field)
	}
}
