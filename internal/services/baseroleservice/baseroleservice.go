package baseroleservice

import (
	"project-skbackend/internal/controllers/responses"
	"project-skbackend/internal/models"
	"project-skbackend/internal/repositories/adminrepo"
	"project-skbackend/packages/consttypes"
	"project-skbackend/packages/utils/utrole"
)

type (
	BaseRoleService struct {
		radmn adminrepo.IAdminRepository
	}

	IBaseRoleService interface {
		GetAdminByBaseRole(roleres responses.BaseRole) (*models.Admin, error)
	}
)

func NewBaseRoleService(
	radmn adminrepo.IAdminRepository,
) *BaseRoleService {
	return &BaseRoleService{
		radmn: radmn,
	}
}

func (s *BaseRoleService) GetAdminByBaseRole(roleres responses.BaseRole) (*models.Admin, error) {
	var (
		a   *models.Admin
		err error
	)

	rid, rtype, ok := utrole.RoleTranslate(roleres)
	if !ok {
		return nil, consttypes.ErrUserInvalidRole
	}

	switch rtype {
	case consttypes.UR_ADMIN:
		a, err = s.radmn.GetByID(rid)
		if err != nil {
			return nil, consttypes.ErrAdminNotFound
		}
	default:
		return nil, consttypes.ErrUserInvalidRole
	}

	return a, nil
}
