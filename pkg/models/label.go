/**
 * 
 * @author dengshaojun
 * @create 2019-04-10 15:40
 */
package models

import "time"

type Label struct {
	Id          int64     `json:"id"`
	AppID       uint      `xorm:"varchar(25) notnull unique 'app_id'" json:"appId"`
	Key         string    `json:"label"`
	Value       string    `json:"value"`
	Rule        string    `json:"rule"`
	DefaultNode string    `json:"defaultNode"`
	CreatedAt   time.Time `xorm:"created" json:"createAt"`
	UpdatedAt   time.Time `xorm:"updated" json:"updatedAt"`
	DeletedAt   time.Time `xorm:"deleted" json:"deletedAt"`
}


func (Label) TableName() string {
	return "t_label"
}

