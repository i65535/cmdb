/**
 * 
 * @author dengshaojun
 * @create 2019-04-12 14:39
 */
package dao

import (
	"cmdb/pkg/database/mysql"
	"cmdb/pkg/models"
)

type LabelDaoImpl struct {
	mysql.MySqlDaoImpl
}

func NewLabelDao() *LabelDaoImpl {
	dao := LabelDaoImpl{}
	m := new(models.Label)
	dao.Init(m.TableName(), m)
	dao.NewSlice = func() interface{} {
		var slice []models.Label
		return &slice
	}

	return &dao
}