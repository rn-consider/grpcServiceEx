package dao

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DB *gorm.DB // 导出 DB 变量以在其他包中使用
)

func InitMySQL() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("conf")
	err = viper.ReadInConfig()
	if err != nil {
		panic("Reading config fatal!")
	}
	dbhost := viper.GetString("database.mysql.host")
	dbusername := viper.GetString("database.mysql.user")
	dbpassword := viper.GetString("database.mysql.password")
	dbport := viper.GetString("database.mysql.port")
	dbname := viper.GetString("database.mysql.dbname")
	maxIdleConns := viper.GetInt("database.mysql.maxIdleConns") // 注意这里使用GetInt
	maxOpenConns := viper.GetInt("database.mysql.maxOpenConns") // 注意这里使用GetInt
	dsn := dbusername + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		return
	}

	// 设置连接池参数
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetMaxIdleConns(maxIdleConns)
	// 自动迁移数据库表（根据您的数据模型定义）

	return sqlDB.Ping()
}

func Close() {
	if DB != nil {
		sqlDB, _ := DB.DB()
		sqlDB.Close()
	}
}
