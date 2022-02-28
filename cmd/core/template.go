package core

import (
	"github.com/Yolo-zb/gin-console/helper"
	"github.com/Yolo-zb/gin-console/src/gorm"
	"github.com/Yolo-zb/gin-console/src/template"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command {
	Use: "create-model",
	Short: "version subcommand show git version info.",
	Run: func(cmd *cobra.Command, args []string) {
		camel := helper.TranCamel(args[0])
		model := template.Model{
			TableName : args[0],
			Camel : helper.LowerFirst(camel[:1]) + camel[1:],
			BigCamel : camel,
			PathTemplate : map[string]string{
				"dao" : template.DaoTemplate,
				"service" : template.ServiceTemplate,
				"model" : template.ModelTemplate,
			},
			Gorm: gorm.A{},
		}
		model.Execute()
	},
}