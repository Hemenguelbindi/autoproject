package hitdir

import (
	"fmt"

	"github.com/spf13/cobra"
)


var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "0.0.1",
	Long:  "This is programm in create project in you lang programming",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Auto project generate procject and happy v0.0.1")
	},
}
