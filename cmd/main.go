package main

import (
	"github.com/lozovoya/gohomework10_1/pkg/cbrf"
	"github.com/lozovoya/gohomework10_1/pkg/export"
	"github.com/lozovoya/gohomework10_1/pkg/locsave"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	const timeout = 10
	const JSONFile = "export.json"
	url, ok := os.LookupEnv("CBRF_URL")
	if !ok {
		log.Println("no URL available")
		os.Exit(1)
	}

	if err := execute(url, timeout, JSONFile); err != nil {
		log.Println(err)
		os.Exit(1)
	}

}

func execute(url string, timeout int, filename string) (err error) {

	svc := cbrf.NewService(url, time.Duration(timeout), &http.Client{})

	extract, err := svc.Extract()
	if err != nil {
		log.Println(err)
		return err
	}

	storage, err := locsave.SaveData(extract)
	if err != nil {
		log.Println(err)
		return err
	}

	err = export.ExportJSON(storage, filename)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
