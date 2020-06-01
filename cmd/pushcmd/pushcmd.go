package pushcmd

import (
	"fmt"
	"log"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/nguyenvanduocit/system-care/internal/gosmc"
	"github.com/nguyenvanduocit/system-care/internal/smc"
)

var metrics = map[string]string{
	"TC0P": "cpu_1_proximity",
	"TC1C": "cpu_core_1",
	"TC2C": "cpu_core_2",
	"TC3C": "cpu_core_3",
	"TC4C": "cpu_core_4",
	"TC5C": "cpu_core_5",
	"TC6C": "cpu_core_6",
	"TaLC": "airflow_left",
	"TaRC": "airflow_right",
}

var RootCmd = &cobra.Command{
	Use: "push",
	RunE: handler,
}

func handler (cmd *cobra.Command, args []string) error {
	server := viper.GetString("server")
	token := viper.GetString("token")
	org := viper.GetString("org")
	bucket := viper.GetString("bucket")

	statChan := make(chan map[string]interface{})
	errChan := make(chan error)

	go getStat(statChan, errChan)
	go pushStat(statChan, token, bucket, org, server)
	log.Println(<-errChan)
	return nil
}


func getStat(resultChan chan map[string]interface{}, errChan chan error) () {
	c, res := gosmc.SMCOpen("AppleSMC")
	if res != gosmc.IOReturnSuccess {
		errChan <- fmt.Errorf("unable to open Apple SMC; return code %v", res)
		return
	}
	defer gosmc.SMCClose(c)

	stats := make(map[string]interface{})
	for {
		for key, name := range metrics {
			f, _, err := smc.GetKeyFloat32(c,key)
			if err != nil {
				errChan <- err
				return
			}
			stats[name] = f
		}
		resultChan <- stats
		time.Sleep(1 * time.Minute)
	}
}

func pushStat(statChan chan map[string]interface{}, token, bucket, org, server string) {
	client := influxdb2.NewClientWithOptions(server, token, influxdb2.DefaultOptions().SetUseGZip(true))
	defer client.Close()

	writeApi := client.WriteApi(org, bucket)
	for {
		select {
			case stat := <-statChan:
				p := influxdb2.NewPoint(
					"temperature",
					nil,
					stat,
					time.Now())
				writeApi.WritePoint(p)
				writeApi.Flush()
		}
	}
}
