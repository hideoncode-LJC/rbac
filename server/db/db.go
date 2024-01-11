package db

import (
	"server/config"
	"server/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"fmt"
)


// checkAndCreateTable 检查表是否存在，不存在则创建表
func checkAndCreateTable(db *gorm.DB, model interface{}) {
	// 检查表是否存在
	if !db.Migrator().HasTable(model) {
		// 如果表不存在，则创建表
		if err := db.AutoMigrate(model); err != nil {
			fmt.Println(err)
		}
	}
}

func GetDNS(mysqlConfig *config.MysqlConfig) string {
	return mysqlConfig.Username +
		":" +
		mysqlConfig.Password +
		"@(" +
		mysqlConfig.Host +
		":" +
		mysqlConfig.Port +
		")/" +
		mysqlConfig.Database +
		"?charset=utf8mb4&parseTime=True&loc=Local"
}

func GetDatabaseConnection() *gorm.DB {
	dsn := GetDNS(config.NewMysqlConfig())
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	checkAndCreateTable(db, &model.Role{})
	checkAndCreateTable(db, &model.User{})
	checkAndCreateTable(db, &model.Access{})
	checkAndCreateTable(db, &model.RoleAccess{})
	checkAndCreateTable(db, &model.RoleInheritance{})
	
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
}