package handler

import (
	"errors"
	"strconv"
	"time"
)

type Dates struct {
	Start time.Time
	End   time.Time
}

// validateParams code params
func validateParams(code string) error {

	if len(code) != 3 {
		return errors.New("len currency not valid")
	}

	_, err := strconv.Atoi(code)
	if err == nil {
		return errors.New("type currency code not valid")
	}

	return nil
}

// checkDates validates start and end queries
func checkDates(finit, fend string) (*Dates, error) {
	var (
		start, end time.Time
		err        error
	)

	if len(finit) > 0 {
		start, err = time.Parse("2006-01-02T15:04:05", finit)
		if err != nil {
			return nil, err
		}
	}

	if len(fend) > 0 {
		end, err = time.Parse("2006-01-02T15:04:05", fend)
		if err != nil {
			return nil, err
		}
	}

	dates := &Dates{
		Start: start,
		End:   end,
	}

	return dates, nil
}
