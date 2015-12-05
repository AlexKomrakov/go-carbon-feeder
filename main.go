package main

import (
	"fmt"
	"time"
)



func main() {
	config, err := Config{}.FromYml("config.yml")
	if err != nil {
		panic(err)
	}


	ticker := time.NewTicker(config.Interval * time.Second)
	quit := make(chan struct{})
	func() {
		for {
			select {
			case <- ticker.C:
				for _, metric := range config.Metric {
					out, err := metric.Value.Run()
					if err != nil {
						panic(err)
					}

					fmt.Printf("%s %s", metric.Key, out)
				}

			case <- quit:
				ticker.Stop()
				return
			}
		}
	}()



}