package models

import (
	"project-skbackend/internal/models/helper"
	"project-skbackend/packages/consttypes"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type (
	Caregiver struct {
		helper.Model
		UserID      uuid.UUID         `json:"userID" gorm:"not null" binding:"required" example:"f7fbfa0d-5f95-42e0-839c-d43f0ca757a4"`
		User        User              `json:"user"`
		Gender      consttypes.Gender `json:"gender" gorm:"not null" binding:"required" example:"Male"`
		FirstName   string            `json:"firstName" gorm:"not null" binding:"required" example:"Jonathan"`
		LastName    string            `json:"lastName" gorm:"not null" binding:"required" example:"Vince"`
		DateOfBirth datatypes.Date    `json:"date" gorm:"not null" binding:"required"` // TODO - add an example on the dateofbirth based on what format is the date
	}
)
