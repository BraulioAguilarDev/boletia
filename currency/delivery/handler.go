package handler

import (
	"boletia/currency"
	"boletia/pkg/request"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/glog"
)

const ALL = "all"

type Handler struct {
	Usecase currency.Usecase
}

type Dates struct {
	Start time.Time
	End   time.Time
}

func NewHandler(usecase currency.Usecase) *Handler {
	return &Handler{
		Usecase: usecase,
	}
}

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

// GetCurrencies is a handler an is used in /api/v1/currencies
func (h *Handler) GetCurrencies(ctx *fiber.Ctx) error {
	// Get Params and query
	currency := ctx.Params("currency")
	finit := ctx.Query("finit")
	fend := ctx.Query("fend")

	// Check format dates & return error if are not valid
	dates, err := checkDates(finit, fend)
	if err != nil {
		glog.Errorf("Error dates parsing: %s", err.Error())

		ctx.Status(http.StatusBadRequest)
		ctx.JSON(request.Failure(err.Error()))
		return nil
	}

	if currency == ALL {

	} else {
		// Pass values to usecase -- currencies
		currencies, err := h.Usecase.GetCurrenciesByCode(currency, dates.Start, dates.End)
		if err != nil {
			glog.Errorf("Error getting currencies by code: %s", err.Error())

			ctx.Status(http.StatusInternalServerError)
			ctx.JSON(request.Failure(err.Error()))
			return nil
		}

		ctx.JSON(request.Response{
			Success: true,
			Data:    currencies,
		})
	}

	return nil
}
