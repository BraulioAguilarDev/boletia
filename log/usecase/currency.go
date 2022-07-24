package usecase

import (
	"boletia/log"
	"boletia/model"
	"time"
)

type LogUsecase struct {
	LogRepo log.Repository
}

func NewLogUsecase(repo log.Repository) *LogUsecase {
	return &LogUsecase{
		LogRepo: repo,
	}
}

func (l LogUsecase) Create(total time.Duration, code int, url string, tm time.Time) error {

	data := model.CreateLogParams{
		Duration:  int32(total),
		Code:      int32(code),
		Request:   url,
		CreatedAt: tm,
	}

	if err := l.LogRepo.Create(data); err != nil {
		return err
	}

	return nil
}
