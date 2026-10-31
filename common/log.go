package common

import (
	"fmt"
	"log"
	"os"
)

var SysLogger = log.New(os.Stdout, "\n[SYS] ", log.LstdFlags)

func SysLog(v ...any) {
	SysLogger.Print(v...)
}
func SysLogf(format string, v ...any) {
	SysLogger.Printf(format, v...)
}

var ErrLogger = log.New(os.Stdout, "\n[ERR] ", log.LstdFlags|log.Llongfile)

func ErrorLog(v ...any) {
	_ = ErrLogger.Output(2, fmt.Sprint(v...))
}
func ErrorLogf(format string, a ...any) {
	_ = ErrLogger.Output(2, fmt.Sprintf(format, a...))
}

var DebugLogger = log.New(os.Stdout, "\n[DEBUG] ", log.LstdFlags|log.Lshortfile)

func DebugLog(v ...any) {
	if DebugMode {
		_ = DebugLogger.Output(2, fmt.Sprint(v...))
	}
}
func DebugLogf(format string, a ...any) {
	if DebugMode {
		_ = DebugLogger.Output(2, fmt.Sprintf(format, a...))
	}
}

var FatalLogger = log.New(os.Stdout, "\n[FATAL] ", log.LstdFlags|log.Llongfile)

func FatalLog(v ...any) {
	_ = FatalLogger.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}
