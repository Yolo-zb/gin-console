package gorm

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	dbDSNFormat = "%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local"
)

type Conf struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
	Params   string
	MaxConn  int
	MaxOpen  int
}

type gormPool map[string]*gorm.DB

var GormPool gormPool

func InitGormPool(connectConf map[string]Conf)  {
	for name, conf := range connectConf{
		//fmt.Println(conf)
		dsn := fmt.Sprintf(dbDSNFormat, conf.User, conf.Password, conf.Host, conf.Port, conf.Database)
		if conf.Params != "" {
			dsn = fmt.Sprintf("%s&%s", dsn, conf.Params)
		}
		DB, err := gorm.Open("mysql", dsn)
		if err != nil {
			fmt.Println("Gorm 异常：", err)
		}
		//根据*grom.DB对象获得*sql.DB的通用数据库接口
		sqlDb := DB.DB()
		sqlDb.SetMaxIdleConns(conf.MaxConn) //设置最大连接数
		sqlDb.SetMaxOpenConns(conf.MaxOpen) //设置最大的空闲连接数
		data, _ := json.Marshal(sqlDb.Stats()) //获得当前的SQL配置情况
		fmt.Println(string(data))
		GormPool = make(gormPool)
		GormPool[name] = DB
	}
}

type A struct {
}

func (a *A) GetGorm(name string) *gorm.DB {
	return GormPool[name]
}

func GetGorm(name string) *gorm.DB {
	return GormPool[name]
}

func Close(name string) {
	GormPool[name].Close()
}