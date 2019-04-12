/**
 * 
 * @author dengshaojun
 * @create 2019-04-12 13:41
 */
package models

import "cmdb/pkg/database/mysql"

func InitDatabase() error {
	engine,err := mysql.GetEngine()
	if err != nil{
		return err
	}
	engine.Sync2(new(Label))
	return nil
}




