package server

import (
	"boletia/config"
	"boletia/currency"
	"boletia/monitor"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/golang/glog"
)

// saveIndb uses out usecase interface to do the insert
func saveIndb(res *http.Response, usecase currency.Usecase) error {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var data monitor.Response
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	if err := usecase.Create(data); err != nil {
		return err
	}

	return nil
}

// Sync function calls external function to get information
func (app *App) Sync() {
	currencyMonitor := monitor.NewHandler()
	for {
		<-time.After(time.Duration(config.Config.Period) * time.Second)

		// TODO timer

		result, err := currencyMonitor.GetCurrencies()
		if err != nil || result.StatusCode != http.StatusOK {

			// Save the current error
			if err != nil {
				fmt.Println(err)
			}

			glog.Errorf("Error sync...")
			// glog
		} else {
			// Keeping info
			if err := saveIndb(result, app.CurrencyUsecase); err != nil {
				glog.Errorf("Error saving data from monitor: %s", err.Error())
			}
		}
	}
}
