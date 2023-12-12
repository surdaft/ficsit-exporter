/*
Copyright © 2023 AP-Hunt

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"text/template"

	"github.com/AP-Hunt/ficsit-exporter/pkg/exporter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	frmHostname string = "localhost"
	frmPort     string = "8080"

	outputMetrics bool = false
)

// exporterCmd represents the exporter command
var exporterCmd = &cobra.Command{
	Use: "exporter",
	Run: func(cmd *cobra.Command, args []string) {
		if viper.GetBool("outputMetrics") {
			outputRawMetrics()
			return
		}

		promExporter := exporter.NewPrometheusExporter("http://" + viper.GetString("hostname") + ":" + viper.GetString("port"))
		promExporter.Start()

		fmt.Printf(`
		Ficsit Remote Monitoring Companion (v%s)

		To access the realtime map visit:
		http://localhost:8000/?frmport=8080

			If you have configured Ficsit Remote Monitoring
			to use a port other than 8080 for its web server,
			change the "frmport" query string parameter to
			match the port you chose and refresh the page.

		To access metrics in Prometheus visit:
		http://localhost:9090/

		Press Ctrl + C to exit.
		`, cmd.Version)

		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		<-sigChan

		err := promExporter.Stop()
		if err != nil {
			fmt.Printf("error stopping prometheus exporter: %s", err)
		}
	},
}

func outputRawMetrics() {
	tpl := template.New("metrics_table")
	tpl.Funcs(template.FuncMap{
		"List": strings.Join,
	})

	tpl, err := tpl.Parse(`
<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Description</th>
            <th>Labels</th>
        </tr>
    </thead>
    <tbody>
		{{range .Metrics}}
        <tr>
            <td>{{.Name}}</td>
            <td>{{.Help}}</td>
            <td>{{List .Labels ", "}}</td>
        </tr>
		{{ end -}}
	</tbody>
</table>
`)
	if err != nil {
		fmt.Printf("Error generating metrics table: %s", err)
		os.Exit(1)
	}

	tpl.Execute(
		os.Stdout,
		struct {
			Metrics []exporter.MetricVectorDetails
		}{
			Metrics: exporter.RegisteredMetricVectors,
		},
	)
}

func init() {
	rootCmd.AddCommand(exporterCmd)

	exporterCmd.Flags().StringVar(&frmHostname, "hostname", frmHostname, "hostname of Ficsit Remote Monitoring webserver")
	exporterCmd.Flags().StringVar(&frmPort, "port", frmPort, "port of Ficsit Remote Monitoring webserver")
	exporterCmd.Flags().BoolVar(&outputMetrics, "outputMetrics", outputMetrics, "show metrics and exit")

	viper.BindPFlag("hostname", exporterCmd.Flags().Lookup("hostname"))
	viper.BindPFlag("port", exporterCmd.Flags().Lookup("port"))
	viper.BindPFlag("outputMetrics", exporterCmd.Flags().Lookup("outputMetrics"))
}
