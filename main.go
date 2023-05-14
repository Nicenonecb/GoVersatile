package main

import (
	"github.com/nicenonecb/GoVersatile/internal/app/routers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// 数据库端口号、用户名和密码更新 DSN
	dsn := "user=postgres password=postgre host=localhost port=5452 dbname=your_database_name sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	//启动路由
	r := routers.SetupRouter(db)
	r.Run(":9090")
}
