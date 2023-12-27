package model

import (
	"gorm.io/gorm"
)

// User 表
type User struct {
	gorm.Model
	Username    string  `json:"username"`// 用户名
	Password string  `json:"password"`// 密码
	Email string  `json:"email"`// 邮箱
	IsAdmin bool  `json:"isadmin"`// 是否是管理员
	RoleID uint  `json:"roleid"`
	Role Role 	`json:"role"`
}

// Role 表
type Role struct {
	gorm.Model
	Name   string `json:"rolename"`
	Users  []User 
	Accesses []Access `gorm:"many2many:role_accesses;foreignKey:ID;associationForeignKey:ID"`
}


// Access 表
type Access struct {
	gorm.Model
	Title string `json:"title"`// 界面的名字 
	Content string `json:"content"`// 界面的内容
	Roles []Role `gorm:"many2many:role_accesses;foreignKey:ID;associationForeignKey:ID"`
}

// RoleAccess 表
type RoleAccess struct {
	RoleID  uint // 角色ID
	AccessID uint // 界面ID
	/*
	权限等级
	1： 读权限
	2： 读、写权限
	3： 读、写、删除权限
	*/
	Level  int 
}


func NewUser(name, password, email string, isAdmin, status bool, roleID uint) *User {
	return &User{
		Username:    name,
		Password: password,
		Email: email,
		IsAdmin: isAdmin,
		RoleID: roleID,
	}
}

func NewRole(name string) *Role {
	return &Role {
		Name: name,
	}
}

func NewAccess(title, content string) *Access {
	return &Access {
		Title: title,
		Content: content,
	}
}

type AccessWithLevel struct {
	Access Access
	Level int
}

type AddRoleAccess struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Level []string `json:"level"`
}

type ResRole struct {
	ID uint
	Name   string 
	Accesses []AccessWithLevel
}

type AddRoleModel struct {
	RoleName string `json:"roleName"`
	Accesses []AddRoleAccess `json:"accesses"`
}

type AddAccessModel struct {
	Title string `json:"title"`
	Content string `json:"content"`
}