/**
 * 
 * @author dengshaojun
 * @create 2019-04-10 15:10
 */
package command

import (
	"cmdb/pkg/models"
	"github.com/spf13/cobra"
)

var xormCmd = &cobra.Command{
	Use:   "xorm",
	Short: "Test xorm interface",
	Run: func(cmd *cobra.Command, args []string) {
		err := models.InitDatabase()
		if err != nil {
			panic(err.Error())
		}
	},
}

