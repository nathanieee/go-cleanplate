package userservice

import (
	"encoding/json"
	"project-skbackend/internal/controllers/requests"
	"project-skbackend/internal/controllers/responses"
	"project-skbackend/internal/models"
	"project-skbackend/internal/models/base"
	"project-skbackend/internal/repositories/adminrepo"
	"project-skbackend/internal/repositories/userrepo"
	"project-skbackend/packages/consttypes"
	"project-skbackend/packages/utils/utpagination"
	"project-skbackend/packages/utils/utstring"

	"github.com/google/uuid"
)

type (
	UserService struct {
		ruser userrepo.IUserRepository
		radmn adminrepo.IAdminRepository
	}

	IUserService interface {
		Create(req requests.CreateUser) (*responses.User, error)
		GetByID(uid uuid.UUID) (*responses.User, error)
		FindAll(p utpagination.Pagination) (*utpagination.Pagination, error)
		Delete(uid uuid.UUID) error
		Update(req requests.UpdateUser, uid uuid.UUID) (*responses.User, error)
		GetUserName(uid uuid.UUID) (string, string, error)
		GetRoleDataByUserID(uid uuid.UUID) (*responses.BaseRole, error)

		UpdateOwnPassword(uid uuid.UUID, req requests.UpdatePassword) error
	}
)

func NewUserService(
	ruser userrepo.IUserRepository,
	radmn adminrepo.IAdminRepository,
) *UserService {
	return &UserService{
		ruser: ruser,
		radmn: radmn,
	}
}

func (s *UserService) Create(req requests.CreateUser) (*responses.User, error) {
	var (
		ures *responses.User
	)

	u := &models.User{
		Email:    req.Email,
		Password: req.Password,
	}

	u, err := s.ruser.Create(*u)
	if err != nil {
		return nil, err
	}

	umar, _ := json.Marshal(u)
	err = json.Unmarshal(umar, &ures)
	if err != nil {
		return nil, err
	}

	return ures, err
}

func (s *UserService) GetByID(uid uuid.UUID) (*responses.User, error) {
	u, err := s.ruser.GetByID(uid)
	if err != nil {
		return nil, err
	}

	ures, err := u.ToResponse()
	if err != nil {
		return nil, err
	}

	return ures, err
}

func (s *UserService) FindAll(p utpagination.Pagination) (*utpagination.Pagination, error) {
	users, err := s.ruser.FindAll(p)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) Delete(uid uuid.UUID) error {
	u := models.User{
		Model: base.Model{ID: uid},
	}

	err := s.ruser.Delete(u)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) Update(
	req requests.UpdateUser,
	uid uuid.UUID,
) (*responses.User, error) {
	u, err := s.ruser.GetByID(uid)
	if err != nil {
		return nil, err
	}

	u, err = req.ToModel(*u, consttypes.UR_USER)
	if err != nil {
		return nil, err
	}

	u, err = s.ruser.Update(*u)
	if err != nil {
		return nil, err
	}

	ures, err := u.ToResponse()
	if err != nil {
		return nil, err
	}

	return ures, err
}

func (s *UserService) GetUserName(uid uuid.UUID) (string, string, error) {
	var (
		firstname string = ""
		lastname  string = ""
	)

	user, err := s.ruser.GetByID(uid)
	if err != nil {
		return "", "", err
	}

	switch user.Role {
	case consttypes.UR_ADMIN:
		a, err := s.radmn.GetByUserID(uid)
		if err != nil {
			return "", "", err
		}

		firstname = a.FirstName
		lastname = a.LastName
	default:
		return "", "", consttypes.ErrUserInvalidRole
	}

	return firstname, lastname, nil
}

func (s *UserService) GetRoleDataByUserID(uid uuid.UUID) (*responses.BaseRole, error) {
	var (
		data any
	)

	user, err := s.ruser.GetByID(uid)
	if err != nil {
		return nil, err
	}

	switch user.Role {
	case consttypes.UR_ADMIN:
		a, err := s.radmn.GetByUserID(user.ID)
		if err != nil {
			return nil, err
		}

		data = a
	default:
		return nil, consttypes.ErrUserInvalidRole
	}

	return &responses.BaseRole{
		Data: data,
		Role: user.Role,
	}, nil
}

func (s *UserService) UpdateOwnPassword(uid uuid.UUID, req requests.UpdatePassword) error {
	user, err := s.ruser.GetByID(uid)
	if err != nil {
		return err
	}

	ok := utstring.CheckPasswordHash(req.OldPassword, user.Password)
	if !ok {
		return consttypes.ErrInvalidEmailOrPassword
	}

	_, err = s.ruser.UpdatePassword(uid, req.NewPassword)
	if err != nil {
		return err
	}

	return nil
}
