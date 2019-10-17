package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

type Users struct {
	gorm.Model
	Name string `json:"name"`
}
type Roles struct {
	gorm.Model
	RoleName string `json:"roleName"`
}
