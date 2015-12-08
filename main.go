package main

import (
	"fmt"
	"time"
"text/template"
	"bytes"
	"regexp"
	"encoding/json"
)

var (
	data = Data{make(map[string]interface{})}
)

type Data struct {
	Variables map[string]interface{}
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

		if (metric.Regex != "") {
			re1 := regexp.MustCompile(metric.Regex)
			result:= re1.FindAllStringSubmatch(string(out), -1)
			data.Variables[metric.Key] = result
		} else if (metric.Json) {
			var f interface{}
			err := json.Unmarshal(out, &f)
			if err != nil {
				panic(err)
			}
			data.Variables[metric.Key] = f
		} else {
			data.Variables[metric.Key] = string(out)
		}
	}

	func() {
		for {
			select {
			case <-ticker.C:
				tmpl, err := template.New("value").Parse(config.Metric)
				if err != nil {
					panic(err)
				}
				value := new(bytes.Buffer)
				err = tmpl.Execute(value, data)

				fmt.Printf("%s", value)
			}
		}
	}()



}