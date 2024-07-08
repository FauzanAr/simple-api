package usermodel

import userentity "simple-api.com/m/src/modules/users/entities"

type UserAllDetailResponse struct {
	Users []userentity.UserAll `json:"users"`
}
