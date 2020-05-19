package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
	"zeus/pkg/api/dto"
	"zeus/pkg/api/model"
)

type Role struct {
}

//Get - get single roel info
func (Role) Get(id int, preload bool) model.Role {
	var role model.Role
	db := GetDb()
	if preload {
		db = db.Preload("Domain")
	}
	db.Where("id = ?", id).First(&role)
	return role
}

// GetRolesByIds
func (Role) GetRolesByIds(ids string) []model.Role {
	var roles []model.Role
	db := GetDb()
	db.Where("id in (?)", strings.Split(ids, ",")).Find(&roles)
	return roles
}

// GetRolesByNames
func (Role) GetRolesByNames(names []string) []model.Role {
	var roles []model.Role
	db := GetDb()
	db = db.Preload("Domain").Preload("DataPerm")
	db.Where("role_name in (?)", names).Find(&roles)
	return roles
}

//Get - get single roel infoD
func (u Role) GetByName(name string) model.Role {
	var role model.Role
	db.Where("role_name = ?", name).Preload("Domain").First(&role)
	return role
}

// List - users list
func (u Role) List(listDto dto.GeneralListDto) ([]model.Role, int64) {
	var roles []model.Role
	var total int64
	db := GetDb()
	for sk, sv := range dto.TransformSearch(listDto.Q, dto.RoleListSearchMapping) {
		db = db.Where(fmt.Sprintf("%s = ?", sk), sv)
	}
	db = db.Preload("DataPerm")
	db.Preload("Domain").Offset(listDto.Skip).Limit(listDto.Limit).Find(&roles)
	db.Model(&model.Role{}).Count(&total)
	return roles, total
}

// Create - new role
func (u Role) Create(role *model.Role) *gorm.DB {
	var row model.Role
	db := GetDb()
	db.Where("name = ? or role_name = ?", role.Name, role.RoleName).First(&row)
	if row.Id > 0 {
		return nil
	}
	return db.Create(role)
}

// Update - update role
func (u Role) Update(role *model.Role, ups map[string]interface{}) *gorm.DB {
	db := GetDb()
	return db.Model(role).Update(ups)
}

// Delete - delete role
func (u Role) Delete(role *model.Role) *gorm.DB {
	db := GetDb()
	return db.Delete(role)
}
