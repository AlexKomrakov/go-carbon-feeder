package main

import (
	"fmt"
	"time"
"text/template"
	"bytes"
)

var (
	data = Data{make(map[string]string)}
)

type Data struct {
	Variables map[string]string
}

func main() {
	config, err := Config{}.FromYml("config.yml")
	if err != nil {
		panic(err)
	}


	ticker := time.NewTicker(config.Interval * time.Second)

	for _, metric := range config.Variables {
		out, err := metric.Value.Run()
		if err != nil {
			panic(err)
		}

		data.Variables[metric.Key] = string(out)
	}

	func() {
		for {
			select {
			case <-ticker.C:
				for _, metric := range config.Metric {
					out, err := metric.Value.Run()
					if err != nil {
						panic(err)
					}

					tmpl, err := template.New("value").Parse(metric.Key)
					if err != nil {
						panic(err)
					}
					value := new(bytes.Buffer)
					err = tmpl.Execute(value, data)

					fmt.Printf("%s %s", value, out)
				}
			}
		}
	}()



}