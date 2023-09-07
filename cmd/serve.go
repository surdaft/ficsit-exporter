package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/surdaft/ficsit-exporter/prober"
)

var serve = &cobra.Command{
	Use:   "serve",
	Short: "start webserver",
	Run: func(cmd *cobra.Command, args []string) {
		prober := prober.New(cmd)

		r := gin.Default()
		r.GET("/probe", prober.Handle)
		r.Run()
	},
}
