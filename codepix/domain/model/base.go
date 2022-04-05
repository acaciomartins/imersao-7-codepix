package model

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Base struct {
	ID       string    `json:"id" valid:"uuid"`
	CreateAt time.Time `json:"created_at" valid:"-"`
	UpdateAt time.Time `json:"update_at" valid:"-"`
}
