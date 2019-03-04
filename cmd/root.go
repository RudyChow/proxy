package cmd

import (
	"fmt"
	"github.com/RudyChow/proxy/api"
	"github.com/RudyChow/proxy/spiders"
	"github.com/RudyChow/proxy/utils/filters"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "proxy",
	Short: "proxypool",
	Long:  `blahblahblah`,
	Run: func(cmd *cobra.Command, args []string) {
		//更新可用的代理ip池
		go filters.UpdateUsefulProxy()

		//开启http服务
		go api.StartHttpServer()

		//开始采集数据
		spiders.StartCrawl()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
