/**
 * 
 * @author dengshaojun
 * @create 2019-04-12 14:11
 */
package models

import "testing"

func TestInitDatabase(t *testing.T) {
	err := InitDatabase()
	if err != nil{
		t.Error("InitDatabase fail,as", err.Error())
	}else{
		t.Log("InitDatabase success")
	}
}