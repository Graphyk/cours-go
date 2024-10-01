package main

import (
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

func argsOrEnvBool(cmd *cobra.Command, args, env string) bool {
	if cmd.Flags().Changed(args) {
		v, _ := cmd.Flags().GetBool(args)
		return v
	}

	return os.Getenv(env) != ""
}

func argsOrEnvString(cmd *cobra.Command, args, env string) string {
	if cmd.Flags().Changed(args) {
		v, _ := cmd.Flags().GetString(args)
		return v
	}

	return os.Getenv(env)
}

var DebugLog bool

var rootCmd = &cobra.Command{
	Use:   "google-shopping-connector",
	Short: "Google Shopping Connector",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logLevel := new(slog.LevelVar)
		switch argsOrEnvString(cmd, "log-level", "LOG_LEVEL") {
		case "debug":
			DebugLog = true
			logLevel.Set(slog.LevelDebug)
		case "info":
			logLevel.Set(slog.LevelInfo)
		case "warning":
			logLevel.Set(slog.LevelWarn)
		case "error":
			logLevel.Set(slog.LevelError)
		}

		if argsOrEnvBool(cmd, "verbose", "VERBOSE") {
			DebugLog = true
			logLevel.Set(slog.LevelDebug)
		}

		logHandlerOptions := &slog.HandlerOptions{Level: logLevel}
		if argsOrEnvBool(cmd, "no-log-time", "NO_LOG_TIME") {
			logHandlerOptions.ReplaceAttr = func(groups []string, a slog.Attr) slog.Attr {
				if a.Key == slog.TimeKey {
					return slog.Attr{}
				}

				return a
			}
		}

		var logHandler slog.Handler
		switch argsOrEnvString(cmd, "log-format", "LOG_FORMAT") {
		case "json":
			logHandler = slog.NewJSONHandler(os.Stdout, logHandlerOptions)
		default:
			logHandler = slog.NewTextHandler(os.Stdout, logHandlerOptions)
		}

		slog.SetDefault(slog.New(logHandler))
	},
}

func main() {
	rootCmd.PersistentFlags().Bool("verbose", false, "Verbose output. Deprecated, use --log-level=debug")
	rootCmd.PersistentFlags().String("log-format", "info", "Log format to use (text or json)")
	rootCmd.PersistentFlags().String("log-level", "text", "Log level to use (debug, info, warning, error)")
	rootCmd.PersistentFlags().Bool("no-log-time", false, "Don't print time in log")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
