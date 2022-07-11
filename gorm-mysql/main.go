package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

//docker run -itd -e MYSQL_ROOT_PASSWORD='bgbiao.top' --name go-orm-mysql  -p 13306:3306 mysql:5.6
//docker exec -it go-orm-mysql mysql -u root -p
//bgbiao.top
type User struct {
	Id       uint   `gorm:"AUTO_INCREMENT"`
	Name     string `gorm:"size:50"`
	Age      int    `gorm:"size:3"`
	Birthday *time.Time
	Email    string `gorm:"type:varchar(50);unique_index"`
	PassWord string `gorm:"type:varchar(25)"`
}

var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open("mysql", "root:mky@(localhost:3306)/test_api?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("fail to create connection to mysql:%v", err)
	}
	defer db.Close()
	// 自动迁移数据结构(table schema)
	// 注意:在gorm中，默认的表名都是结构体名称的复数形式，比如User结构体默认创建的表为users
	// db.SingularTable(true) 可以取消表名的复数形式，使得表名和结构体名称一致
	db.AutoMigrate(&User{})

	// 添加唯一索引
	db.Model(&User{}).AddUniqueIndex("name_email", "id", "name", "email")

	db.Create(&User{Name: "mky", Age: 19, Email: "111111@qq.com"})
	db.Create(&User{Name: "mky2", Age: 3, Email: "222222@qq.com"})

	var user User
	var users []User

	//查看插入后全部元素
	db.Find(&users)
	fmt.Println(users)
	//查询一条记录
	db.First(&user, "name = ?", "mky")
	fmt.Println("查看查询记录：", user)
	db.Model(&user).Update("name")
}
