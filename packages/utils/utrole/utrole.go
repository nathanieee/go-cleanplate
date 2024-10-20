package utrole

import (
	"project-skbackend/internal/controllers/responses"
	"project-skbackend/internal/models"
	"project-skbackend/packages/consttypes"

	"github.com/google/uuid"
)

func RoleTranslate(role responses.BaseRole) (uuid.UUID, consttypes.UserRole, bool) {
	switch role.Role {
	// * add another role here before processing it
	case consttypes.UR_ADMIN:
		res, ok := role.Data.(*models.Admin)
		if !ok {
			return uuid.UUID{}, consttypes.UserRole(0), false
		}
		return res.ID, res.User.Role, true
	default:
		return uuid.UUID{}, consttypes.UserRole(0), false
	}
}
