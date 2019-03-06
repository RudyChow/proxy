package cmd

import (
	"github.com/RudyChow/proxy/app/http"
	"github.com/RudyChow/proxy/app/schedules"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "proxy",
	Short: "proxypool",
	Long:  `blahblahblah`,
	Run: func(cmd *cobra.Command, args []string) {
		go schedules.Run()

		//开启http服务
		http.StartHttpServer()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
