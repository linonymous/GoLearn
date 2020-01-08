/*
   Created by LINONYMOUS
   at 9/1/2019
*/
package Lesson_4

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
)

type Config struct {
	
}

func readConfig(path string) (*Config, error)  {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "Can't open configuration file")
	}
	
	defer file.Close()
	
	cfg := &Config{}
	
	return cfg, nil
}

func setupLogging()  {
	out, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	log.SetOutput(out)
}

func main()  {
	setupLogging()
	cfg, err := readConfig("kuchh/bhi/path")
	if err != nil {
		log.Printf("ERROR: %+v\n", err)
		n, errr := fmt.Fprintf(os.Stdout, "Error: %s ", err)

		if errr != nil {
			os.Exit(n)
		}
		os.Exit(n)
	}

	fmt.Println(cfg)
}
