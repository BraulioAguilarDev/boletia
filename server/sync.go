package server

import (
	"boletia/config"
	"boletia/currency"
	"boletia/log"
	"boletia/monitor"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/golang/glog"
)

// saveCurrencies uses out usecase interface to do the insert
func saveCurrencies(usecase currency.Usecase, res *http.Response) error {
	glog.Infoln("Saving currencies in db...")

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

// saveLog keep log request
func saveLog(usecase log.Usecase, total time.Duration, code int, url string, tm time.Time) error {
	glog.Infof("Saving the log...")

	if err := usecase.Create(total, code, url, tm); err != nil {
		return err
	}
	return nil
}

// Sync function calls external function to get information
func (app *App) Sync() {

	var code int

	currencyMonitor := monitor.NewHandler()
	for {
		<-time.After(time.Duration(config.Config.Period) * time.Second)

		// fix: timer
		start := time.Now()

		glog.Infof("URL: %s", config.Config.ApiURL)
		result, end, err := currencyMonitor.GetCurrencies()
		if result.StatusCode == http.StatusOK {
			// Keeping info
			if err := saveCurrencies(app.CurrencyUsecase, result); err != nil {
				glog.Errorf("Error saving currencies: %s", err.Error())
			}
		}

		if err != nil {
			code = http.StatusInternalServerError
		} else {
			code = result.StatusCode
		}

		total := end.Sub(start)
		if err := saveLog(app.LogUsecase, total, code, config.Config.ApiURL, time.Now()); err != nil {
			glog.Errorf("Error saving logs: %s", err.Error())
		}

		glog.Infoln("...................................")
	}
}
