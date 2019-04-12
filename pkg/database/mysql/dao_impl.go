/**
 * Data Access Object，数据库访问对象
 * @author dengshaojun
 * @create 2019-04-12 14:23
 */
package mysql

import (
	"cmdb/pkg/utils"
	"github.com/go-xorm/xorm"
)

type PageData struct {
	Page       int         `json:"page"`
	PageSize   int         `json:"pageSize"`
	TotalCount int         `json:"totalCount"`
	Data       interface{} `json:"data"`
}

type MySqlDaoImpl struct {
	engine    *xorm.Engine
	tableName string
	NewSlice  func() interface{}
	bean      interface{}
}

func (m *MySqlDaoImpl)Init(tableName string, bean interface{}) {
	engine, err := GetEngine()
	if err != nil{
		utils.Panic("mysql get engine fail,as:", err.Error())
	}
	m.engine = engine
	m.bean = bean
	m.tableName = tableName
}

func (m *MySqlDaoImpl)Create(beans ...interface{}) (int64, error) {
	return m.engine.Table(m.tableName).Insert(beans)
}

func (m *MySqlDaoImpl)Count(bean ...interface{}) (int64, error) {
	return m.engine.Table(m.tableName).Count(bean)
}

func (m *MySqlDaoImpl)DeleteBy(query interface{}, args ...interface{}) (int64, error) {
	return m.engine.Table(m.tableName).Where(query, args).Delete(m.bean)
}

func (m *MySqlDaoImpl)DeleteById(id interface{}) (int64, error) {
	return m.engine.Table(m.tableName).ID(id).Delete(m.bean)
}

func (m *MySqlDaoImpl)Update(id interface{}, bean interface{}) (int64, error) {
	return m.engine.Table(m.tableName).ID(id).Update(bean)
}

func (m *MySqlDaoImpl)FindById(id interface{}, bean interface{}) (bool, error) {
	return m.engine.Table(m.tableName).ID(id).Get(bean)
}

func (m *MySqlDaoImpl)Find(bean interface{}, query interface{}, args ...interface{}) (bool, error) {
	return m.engine.Table(m.tableName).Where(query, args).Get(bean)
}

func (m *MySqlDaoImpl)ListAll(beans interface{}) (error) {
	return m.engine.Table(m.tableName).Find(beans)
}

func (m *MySqlDaoImpl)ListBy(beans interface{}, query interface{}, args ...interface{}) (error) {
	return m.engine.Table(m.tableName).Where(query, args).Find(beans)
}

func (m *MySqlDaoImpl)ListCol(cols interface{}, columns string, query interface{}, args ...interface{}) (error) {
	return m.engine.Table(m.tableName).Cols(columns).Where(query, args).Find(cols)
}

func (m *MySqlDaoImpl)ListPage(page, pageSize int, orderBy string, query interface{}, args ...interface{}) (*PageData, error) {
	total, err := m.engine.Table(m.tableName).Where(query, args).Count(m.bean)
	if err != nil{
		return nil, err
	}
	beans := m.NewSlice()
	err = m.engine.Table(m.tableName).Where(query, args).OrderBy(orderBy).Limit(pageSize, (page-1) * pageSize).Find(beans)
	if err != nil{
		return nil, err
	}
	return &PageData{
		Page:page,
		PageSize:pageSize,
		TotalCount:int(total),
		Data:beans,
	}, nil
}



