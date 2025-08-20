package logger

import (
	"log/slog"
	"os"

	slogmulti "github.com/samber/slog-multi"
)

// NewLogger creates a new structured logger that writes to the console and a file.
// It now returns the file so its lifecycle can be managed.
func NewLogger() (*slog.Logger, *os.File) {
	// Open the log file.
	// In a real app, you might get the file path from a config.
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		// Fallback to a console-only logger if the file can't be opened.
		return slog.New(slog.NewJSONHandler(os.Stdout, nil)), nil
	}

	// Create a handler for console output (stdout) for development.
	consoleHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	// Create a handler for file output for production records.
	fileHandler := slog.NewJSONHandler(file, &slog.HandlerOptions{
		Level: slog.LevelInfo, // Only log Info level and higher to the file.
	})

	// Use slog-multi to combine the console and file handlers.
	handler := slogmulti.Fanout(consoleHandler, fileHandler)

	logger := slog.New(handler)

	return logger, file
}
