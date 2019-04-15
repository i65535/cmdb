/**
 * 
 * @author dengshaojun
 * @create 2019-04-12 14:39
 */
package dao

import (
	"cmdb/pkg/database/mysql"
	"cmdb/pkg/models"
	"cmdb/pkg/utils"
	"fmt"
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

func (d *LabelDaoImpl) CreateLabel(label *models.Label) (int64, error) {
	affected, err := d.Create(*label)
	if err != nil{
		utils.Errorf("LabelDaoImpl CreateLabel [%v] fail,as:%s", label, err.Error())
		return 0, err
	}else{
		return affected, nil
	}
}


func (d *LabelDaoImpl) CreateLabelByMap(label *map[string]interface{}) (int64, error) {
	affected, err := d.CreateWithMap(label)
	if err != nil{
		utils.Errorf("LabelDaoImpl CreateLabel [%v] fail,as:%s", label, err.Error())
		return 0, err
	}else{
		return affected, nil
	}
}

func (d *LabelDaoImpl) DeleteLabel(labelId int64) (int64, error) {
	affected, err := d.DeleteById(labelId)
	if err != nil{
		utils.Errorf("LabelDaoImpl DeleteLabel [%v] fail,as:%s", labelId, err.Error())
		return 0, err
	}else{
		return affected, nil
	}
}

func (d *LabelDaoImpl) UpdateLabel(label *models.Label) (int64, error) {
	if label == nil{
		return 0, fmt.Errorf("invalid paramter, label=nil")
	}

	affected, err := d.Update(label.Id, label)
	if err != nil{
		utils.Errorf("LabelDaoImpl UpdateLabel [%v] fail,as:%s", label.Id, err.Error())
		return 0, err
	}else{
		return affected, nil
	}
}

func (d *LabelDaoImpl) FindAll() (*[]models.Label, error) {
	var labels []models.Label
	err := d.ListAll(&labels)
	if err != nil{
		utils.Error("LabelDaoImpl FindAll fail,as:", err.Error())
		return nil, err
	}
	return &labels, nil
}

func (d *LabelDaoImpl) ReadLabel(labelId int64) (*models.Label, error) {
	var label models.Label
	exist, err := d.FindById(labelId, &label)
	if err != nil{
		utils.Errorf("LabelDaoImpl Read [%d] fail,as:%s", labelId, err.Error())
		return nil, err
	}else if exist{
		return  &label, nil
	}else{
		utils.Debugf("LabelDaoImpl the label id=[%d] not exist", labelId)
		return nil, fmt.Errorf("the label [%d] not exist", labelId)
	}
}

func (d *LabelDaoImpl) ReadLabelPage(page, pageSize int, orderBy string, query interface{}, args ...interface{}) (*mysql.PageData, error) {
	rlt, err := d.ListPage(page,pageSize, orderBy, query, args...)
	if err != nil{
		utils.Errorf("LabelDaoImpl ReadLabelPage fail,as:%s", err.Error())
		return nil, err
	}else{
		return rlt, nil
	}
}