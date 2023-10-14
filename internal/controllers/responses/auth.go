package responses

import (
	"project-skbackend/packages/consttypes"
	"time"

	"github.com/google/uuid"
)

type (
	AuthResponse struct {
		ID                 uuid.UUID           `json:"id"`
		Email              string              `json:"email" example:"email@email.com"`
		Role               consttypes.UserRole `json:"role" gorm:"not null" example:"f7fbfa0d-5f95-42e0-839c-d43f0ca757a4"`
		ConfirmationSentAt time.Time           `json:"confirmationSentAt,omitempty" example:"2023-01-01T15:01:00+00:00"`
		ConfirmedAt        time.Time           `json:"confirmedAt,omitempty" example:"2023-01-01T15:01:00+00:00"`
		CreatedAt          time.Time           `json:"createdAt,omitempty" example:"2023-01-01T15:01:00+00:00"`
		UpdatedAt          time.Time           `json:"updatedAt,omitempty" example:"2023-02-11T15:01:00+00:00"`
		Token              string              `json:"token,omitempty"`
		Expires            time.Time           `json:"expires,omitempty"`
	}
)
