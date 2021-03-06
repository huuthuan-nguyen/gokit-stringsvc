package main

import (
	"github.com/go-kit/kit/log"
	"time"
)

func loggingMiddleware(logger log.Logger) ServiceMiddleware {
	return func(next StringService) StringService {
		return logmw{logger, next}
	}
}

type logmw struct {
	logger log.Logger
	next   StringService
}

func (mw logmw) Uppercase(s string) (output string, err error) {
	defer func(begin time.Time) {
		err := mw.logger.Log(
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
		if err != nil {
			return
		}
	}(time.Now())

	output, err = mw.next.Uppercase(s)
	return
}

func (mw logmw) Count(s string) (n int) {
	defer func(begin time.Time) {
		err := mw.logger.Log(
			"method", "count",
			"input", s,
			"n", n,
			"took", time.Since(begin),
		)
		if err != nil {
			return
		}
	}(time.Now())

	n = mw.next.Count(s)
	return
}
