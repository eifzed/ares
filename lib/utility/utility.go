package utility

import (
	"net/mail"

	// "github.com/eifzed/antre-app/internal/entity/repo/antre"
	"github.com/eifzed/ares/lib/utility/jwt"
)

func StringExistInSlice(item string, itemSlice []string) bool {
	for _, i := range itemSlice {
		if i == item {
			return true
		}
	}
	return false
}

// func MstRoleExistInSlice(item antre.MstRole, itemSlice []antre.MstRole) bool {
// 	for _, i := range itemSlice {
// 		if i.Name == item.Name {
// 			return true
// 		}
// 	}
// 	return false
// }

func RoleExistInSlice(item jwt.Role, itemSlice []jwt.Role) bool {
	for _, i := range itemSlice {
		if i.ID == item.ID {
			return true
		}
	}
	return false
}

func IntExistInSlice(item int, itemSlice []int) bool {
	for _, i := range itemSlice {
		if i == item {
			return true
		}
	}
	return false
}

func Int32ExistInSlice(item int32, itemSlice []int32) bool {
	for _, i := range itemSlice {
		if i == item {
			return true
		}
	}
	return false
}

func Int64ExistInSlice(item int64, itemSlice []int64) bool {
	for _, i := range itemSlice {
		if i == item {
			return true
		}
	}
	return false
}

func IsValidEmailFormat(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
