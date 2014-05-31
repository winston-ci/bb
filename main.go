package main

import (
	"flag"
	"fmt"
	"os"

	"code.google.com/p/gographviz"
	"github.com/fraenkel/candiedyaml"

	"github.com/winston-ci/winston/config"
)

func filename() string {
	flag.Parse()
	return flag.Arg(0)
}

func fatal(err error) {
	println(err.Error())
	os.Exit(1)
}

func loadConfig() config.Config {
	configFile, err := os.Open(filename())
	if err != nil {
		fatal(err)
	}

	var conf config.Config
	err = candiedyaml.NewDecoder(configFile).Decode(&conf)
	if err != nil {
		fatal(err)
	}
	configFile.Close()

	return conf
}

func main() {
	conf := loadConfig()

	graph := gographviz.NewEscape()
	graph.SetName("winston")
	graph.SetDir(true)

	for _, resource := range conf.Resources {
		label := fmt.Sprintf("Resource: %s", resource.Name)
		graph.AddNode("", label, nil)
	}

	for _, job := range conf.Jobs {
		label := fmt.Sprintf("Job: %s", job.Name)
		graph.AddNode("", label, nil)

		for _, input := range job.Inputs {
			if len(input.Passed) > 0 {
				for _, passed := range input.Passed {
					srcLabel := fmt.Sprintf("Job: %s", passed)
					graph.AddEdge(srcLabel, label, true, nil)
				}
			} else {
				srcLabel := fmt.Sprintf("Resource: %s", input.Resource)
				graph.AddEdge(srcLabel, label, true, nil)
			}
		}

		for _, output := range job.Outputs {
			dstLabel := fmt.Sprintf("Resource: %s", output.Resource)
			attrs := map[string]string{
				"label": fmt.Sprintf("branch: %s", output.Params["branch"]),
			}
			graph.AddEdge(label, dstLabel, true, attrs)
		}
	}

	fmt.Printf("%s", graph)
}
