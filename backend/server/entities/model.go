package entities

import "github.com/google/uuid"

type Model struct {
	ID        uuid.UUID                  `json:"ID"        example:"b0351a3b-aef0-4a7b-8a6b-5303be0b9040"`
	CreatedAt DateYearMonthDayHourMinute `json:"createdAt" example:"2023-01-02 14:31"`
}
