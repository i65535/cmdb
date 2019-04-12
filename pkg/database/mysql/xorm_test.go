/**
 * 
 * @author dengshaojun
 * @create 2019-04-12 14:01
 */
package mysql

import (
	"cmdb/pkg/models"
	"testing"
)

func TestGetEngine(t *testing.T) {
	engine,err := GetEngine()
	if err != nil{
		engine.Sync2(new(models.Label))
	}
}

