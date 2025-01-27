package nlog

//
// type loggerHandler struct {
// 	slog.Handler
// 	Logger *log.Logger
// }
//
// func (h *loggerHandler) Enabled(ctx context.Context, level slog.Level) bool {
// 	return h.Handler.Enabled(ctx, level)
// }
//
// func (h *loggerHandler) Handle(ctx context.Context, record slog.Record) error {
// 	levelStr := record.Level.String() + ":"
// 	switch record.Level {
// 	case slog.LevelDebug:
// 		levelStr = color.Magenta("DEBU")
// 	case slog.LevelInfo:
// 		levelStr = color.Blue("INFO")
// 	case slog.LevelWarn:
// 		levelStr = color.Yellow("WARN")
// 	case slog.LevelError:
// 		levelStr = color.Red("ERRO")
// 	}
// 	msgStr := color.Cyan(record.Message)
// 	timeStr := record.Time.Format("2006-01-02 15:04:05")
// 	// sourceStr := fmt.Sprintf("%s:%d", f, l)
// 	// h.Logger.Println(timeStr, levelStr, msgStr, sourceStr)
// 	h.Logger.Println(timeStr, levelStr, msgStr)
// 	return nil
// }
//
// func newLoggerHandler(level Level) *loggerHandler {
// 	return &loggerHandler{
// 		Handler: slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
// 			Level: level,
// 		}),
// 		Logger: log.New(os.Stdout, "", 0),
// 	}
// }
