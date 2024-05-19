package log

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
	"github.com/mattn/go-isatty"
)

func InitializeLogger() {
	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stderr, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.TimeOnly,
			NoColor:    !isatty.IsTerminal(os.Stderr.Fd()),
		}),
	))
}

func Fatal(v ...any) {
	Error(v...)
	os.Exit(1)
}

func Fatalf(format string, v ...any) {
	Errorf(format, v...)
	os.Exit(1)
}

func Error(v ...any) {
	slog.Error(fmt.Sprint(v...))
}

func Errorf(format string, v ...any) {
	slog.Error(fmt.Sprintf(format, v...))
}

func Println(v ...any) {
	slog.Info(fmt.Sprint(v...))
}

func Printf(format string, v ...any) {
	slog.Info(fmt.Sprintf(format, v...))
}
