package db

import (
	"context"
	application "github.com/debugger84/modulus-application"
	"gorm.io/gorm/logger"
	"runtime/debug"
	"strings"
	"time"
)

type GormLogger struct {
	application.Logger
	cfg *ModuleConfig
}

func NewGormLogger(cfg *ModuleConfig, logger application.Logger) *GormLogger {
	return &GormLogger{Logger: logger, cfg: cfg}
}

func (g GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return g
}

func (g GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	if err != nil && err.Error() != "invalid value" {
		// it is inside the "if" statement only to avoid unnecessary calculations
		//of rows and execution time unlogged queries
		sql, rows := fc()
		if err.Error() == "context canceled" {
			g.Logger.Warn(
				ctx,
				err.Error(),
				"elapsedTime", elapsed,
				"sql", sql,
				"rows", rows,
				"trace", g.getTrace(),
			)
		} else {
			g.Logger.Error(
				ctx,
				err.Error(),
				"elapsedTime", elapsed,
				"sql", sql,
				"rows", rows,
				"trace", g.getTrace(),
			)
		}
	} else if elapsed > time.Duration(g.cfg.slowQueryLimit)*time.Millisecond ||
		(err != nil && err.Error() == "context canceled") {
		sql, rows := fc()
		g.Logger.Warn(
			ctx,
			"Too long execution",
			"elapsedTime", elapsed,
			"sql", sql,
			"rows", rows,
			"trace", g.getTrace(),
		)
	} else if *g.cfg.loggingEnabled {
		sql, rows := fc()
		g.Logger.Debug(
			ctx,
			"SQL execution",
			"elapsedTime", elapsed,
			"sql", sql,
			"rows", rows,
			"trace", g.getTrace(),
		)
	}
}

func (g GormLogger) getTrace() string {
	stack := debug.Stack()
	s := strings.Split(string(stack), "\n")
	l := len(s)
	if l > 16 {
		l = 16
	}
	return strings.Join(s[10:l], "\n")
}
