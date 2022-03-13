package models

import (
    "time"
)

type Ability struct {
    ID          string    `json:"ability_id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    CreatedAt   time.Time `json:"-"`
    UpdatedAt   time.Time `json:"description"`
    DeletedAt   time.Time `json:"-"`
}
