package observability

import (
	"context"
	"io"

	"github.com/sirupsen/logrus"
)

type logrusLoggerKeyType struct{}

var logrusLoggerKey = &logrusLoggerKeyType{}

var nullLogger = &logrus.Logger{
	Out:       io.Discard,
	Formatter: new(logrus.TextFormatter),
	Hooks:     make(logrus.LevelHooks),
	Level:     logrus.PanicLevel,
}

func LoggerFromContext(ctx context.Context) *logrus.Entry {
	logger, ok := ctx.Value(logrusLoggerKey).(*logrus.Entry)
	if !ok || logger == nil {
		return logrus.NewEntry(nullLogger)
	}
	return logger
}

func LoggerToContext(ctx context.Context, logger *logrus.Entry) context.Context {
	return context.WithValue(ctx, logrusLoggerKey, logger)
}
