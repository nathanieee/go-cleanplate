package consttypes

type (
	UserRole uint
)

const (
	UR_USER  UserRole = 0
	UR_ADMIN UserRole = 1
)

func (enum UserRole) Uint() uint {
	return uint(enum)
}
