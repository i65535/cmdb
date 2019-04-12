/**
 * 
 * @author dengshaojun
 * @create 2019-04-10 15:19
 */
package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	//github.com/tidwall/gjson
)



func GetEngine() (*xorm.Engine, error){
	engine, err:=xorm.NewEngine("mysql", "root:@tcp(127.0.0.1:3306)/cmp?charset=utf8")
	if err != nil{
		return nil, err
	}

	err = engine.Ping()
	if err != nil{
		return nil, err
	}

	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	engine.SetMaxIdleConns(2)
	engine.SetMaxOpenConns(20)
	engine.SetTableMapper(core.SameMapper{})
	return engine, nil
}

