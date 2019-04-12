/**
 * 
 * @author dengshaojun
 * @create 2019-04-10 15:40
 */
package models

type Label struct {
	ID          uint   `json:"id"`
	AppID       uint   `json:"appId"`
	Key         string `json:"label"`
	Value       string `json:"value"`
	Rule        string `json:"rule"`
	DefaultNode string `json:"defaultNode"`
}


func (Label) TableName() string {
	return "t_label"
}

