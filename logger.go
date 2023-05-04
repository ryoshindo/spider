package spider

import (
	"log"
	"os"

	"github.com/fatih/color"
	"golang.org/x/exp/slog"
)

// NOTE: In the future, it will be moved to the slog package.

func newLogger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags)
}

func (s *App) Log(level slog.Level, msg string) {
	switch level {
	case slog.LevelDebug:
		s.logDebug(msg)
	case slog.LevelInfo:
		s.logInfo(msg)
	case slog.LevelWarn:
		s.logWarn(msg)
	case slog.LevelError:
		s.logError(msg)
	}
}

func (s *App) logDebug(msg string) {
	s.logger.Printf(msg)
}

func (s *App) logInfo(msg string) {
	s.logger.Printf(msg)
}

func (s *App) logWarn(msg string) {
	s.logger.Printf(color.YellowString(msg))
}

func (s *App) logError(msg string) {
	s.logger.Printf(color.RedString(msg))
}
