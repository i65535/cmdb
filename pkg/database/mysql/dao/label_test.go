/**
 * 
 * @author dengshaojun
 * @create 2019-04-15 15:20
 */
package dao

import (
	"cmdb/pkg/models"
	"encoding/json"
	"testing"
)

func TestLabelDaoImpl_CreateLabel(t *testing.T) {
	label := models.Label{
		AppID:1,
		Key:"mysql",
		Value:"enable",
		Rule:"-",
		DefaultNode:"host001",
	}
	dao := NewLabelDao()
	id, err := dao.CreateLabel(&label)
	if err != nil{
		t.Errorf("LabelDaoImpl CreateLabel fail,as[%s]", err.Error())
	}else{
		t.Log("LabelDaoImpl CreateLabel success, id=", id)
	}
}

func TestLabelDaoImpl_CreateLabel2(t *testing.T) {
	data := map[string]interface{}{
		"app_id":1,
		"key":"mongodb",
		"value":"enable",
		"rule":"-",
		"default_node":"host001",
	}
	dao := NewLabelDao()
	id, err := dao.CreateLabelByMap(&data)
	if err != nil{
		t.Errorf("LabelDaoImpl CreateLabel fail,as[%s]", err.Error())
	}else{
		t.Log("LabelDaoImpl CreateLabel success, id=", id)
	}
}

func TestLabelDaoImpl_DeleteLabel(t *testing.T) {
	dao := NewLabelDao()
	affected, err := dao.DeleteLabel(3)
	if err != nil{
		t.Errorf("LabelDaoImpl DeleteLabel fail,as[%s]", err.Error())
	}else{
		t.Log("LabelDaoImpl DeleteLabel success, affected=", affected)
	}
}

func TestLabelDaoImpl_UpdateLabel(t *testing.T) {
	label := models.Label{
		Id:1,
		AppID:1,
		Key:"mysql",
		Value:"enable",
		Rule:"-",
		DefaultNode:"host001",
	}
	dao := NewLabelDao()
	affected, err := dao.UpdateLabel(&label)
	if err != nil{
		t.Errorf("LabelDaoImpl UpdateLabel fail,as[%s]", err.Error())
	}else{
		t.Log("LabelDaoImpl UpdateLabel success, affected=", affected)
	}
}

func TestLabelDaoImpl_ReadLabelPage(t *testing.T) {
	dao := NewLabelDao()
	labels, err := dao.ReadLabelPage(1, 10, "id ASC", "id > ?", "0")
	if err != nil{
		t.Errorf("LabelDaoImpl DeleteLabel fail,as[%s]", err.Error())
	}else{
		bytes, _ := json.Marshal(labels)
		t.Log("LabelDaoImpl DeleteLabel success, data=", string(bytes))
	}
}

func TestLabelDaoImpl_FindAll(t *testing.T) {
	dao := NewLabelDao()
	labels, err := dao.FindAll()
	if err != nil{
		t.Errorf("LabelDaoImpl DeleteLabel fail,as[%s]", err.Error())
	}else{
		bytes, _ := json.Marshal(labels)
		t.Log("LabelDaoImpl DeleteLabel success, data=", string(bytes))
	}
}

func TestLabelDaoImpl_ReadLabel(t *testing.T) {
	dao := NewLabelDao()
	label, err := dao.ReadLabel(1)
	if err != nil{
		t.Errorf("LabelDaoImpl ReadLabel fail,as[%s]", err.Error())
	}else{
		bytes, _ := json.Marshal(label)
		t.Log("LabelDaoImpl ReadLabel success, data=", string(bytes))
	}
}