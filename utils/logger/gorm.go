package logger

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"
	_ "gorm.io/gorm/utils"
)

// CustomLogger is a custom logger that writes logs to a file
type CustomLogger struct {
	Logger   *log.Logger
	LogLevel logger.LogLevel
}

// NewCustomLogger creates a new CustomLogger
func NewCustomLogger(logLevel logger.LogLevel) logger.Interface {
	file, err := os.OpenFile("logs/gorm.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("failed to open log file")
	}
	return &CustomLogger{
		Logger:   log.New(file, "\r\n", log.LstdFlags),
		LogLevel: logLevel,
	}
}

func (c *CustomLogger) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := *c
	newlogger.LogLevel = level
	return &newlogger
}

func (c *CustomLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if c.LogLevel >= logger.Info {
		c.Logger.Printf("[info] "+msg, data...)
	}
}

func (c *CustomLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if c.LogLevel >= logger.Warn {
		c.Logger.Printf("[warn] "+msg, data...)
	}
}

func (c *CustomLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if c.LogLevel >= logger.Error {
		c.Logger.Printf("[error] "+msg, data...)
	}
}

func (c *CustomLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if c.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && c.LogLevel >= logger.Error:
		sql, rows := fc()
		if rows == -1 {
			c.Logger.Printf("[error] %s [%.3fms] %s", err, float64(elapsed.Nanoseconds())/1e6, sql)
		} else {
			c.Logger.Printf("[error] %s [%.3fms] %d rows affected or returned %s", err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > 200*time.Millisecond && c.LogLevel >= logger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", 200*time.Millisecond)
		if rows == -1 {
			c.Logger.Printf("[warn] %s [%.3fms] %s", slowLog, float64(elapsed.Nanoseconds())/1e6, sql)
		} else {
			c.Logger.Printf("[warn] %s [%.3fms] %d rows affected or returned %s", slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case c.LogLevel >= logger.Info:
		sql, rows := fc()
		if rows == -1 {
			c.Logger.Printf("[info] [%.3fms] %s", float64(elapsed.Nanoseconds())/1e6, sql)
		} else {
			c.Logger.Printf("[info] [%.3fms] %d rows affected or returned %s", float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
