/**
 * 
 * @author dengshaojun
 * @create 2019-04-12 13:41
 */
package models

import (
	"cmdb/pkg/database/mysql"
	"cmdb/pkg/utils"
	"github.com/go-xorm/xorm"
)

func InitDatabase() error {
	engine,err := mysql.GetEngine()
	if err != nil{
		return err
	}

	tables := []interface{}{new(Label)}
	for _,table := range tables{
		t, ok := table.(xorm.TableName)
		if ! ok {
			continue
		}

		exist, err := engine.IsTableExist(t.TableName())
		if err != nil{
			utils.Panic("")
		}
		if ! exist{
			engine.CreateTables(table)
		}
	}
	return nil
}





