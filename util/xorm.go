package util

import (
	"goDnParse/model"
	"xorm.io/xorm"
)

// 全局注册
var Orm *xorm.Engine

func GetOrmEngine() (*xorm.Engine, error) {
	//获取配置文件
	config := Cfg.Mysql

	//注册数据库
	engine, err := xorm.NewEngine(
		"mysql",
		config.User+":"+config.Password+"@("+config.Ip+":"+config.Port+")/"+config.DatabaseName+"?charset="+config.Charset)
	if err != nil {
		return nil, err
	}
	engine.ShowSQL(true)
	engine.Sync2(new(model.User))
	//engine.Sync2(new(model.UserFollow))
	//engine.Sync2(new(model.Article))
	//engine.Sync2(new(model.ArticleComment))
	//engine.Sync2(new(model.Belong))
	//engine.Sync2(new(model.BelongFollower))
	Orm = engine
	return engine, nil
}
