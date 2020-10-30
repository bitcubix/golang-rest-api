package logger

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/viper"
	"log"
	"os"
)

var (
	logger *log.Logger
)

func init() {
	logger = log.New(os.Stdout, "LOG ", log.Ldate|log.Ltime)
	if viper.GetBool("logger.debug") {
		logger.SetFlags(log.Lshortfile)
	}
}

func Debug(args ...interface{}) {
	if viper.GetBool("logger.debug") {
		col := color.New(color.FgHiBlack, color.BgBlue, color.Bold).SprintfFunc()
		logger.SetPrefix(col("DEBUG\t"))
		logger.Println(fmt.Sprint(args...))
	}
}

func Info(args ...interface{}) {
	col := color.New(color.FgHiBlack, color.BgGreen, color.Bold).SprintfFunc()
	logger.SetPrefix(col("INFO\t"))
	logger.Println(fmt.Sprint(args...))
}

func Warning(args ...interface{}) {
	col := color.New(color.FgHiBlack, color.BgYellow, color.Bold).SprintfFunc()
	logger.SetPrefix(col("WARN\t"))
	logger.Println(fmt.Sprint(args...))
}

func Error(args ...interface{}) {
	col := color.New(color.FgHiBlack, color.BgHiRed, color.Bold).SprintfFunc()
	logger.SetPrefix(col("ERROR\t"))
	logger.Println(fmt.Sprint(args...))
}

func Fatal(args ...interface{}) {
	col := color.New(color.FgHiBlack, color.BgRed, color.Bold).SprintfFunc()
	logger.SetPrefix(col("FATAL\t"))
	logger.Println(fmt.Sprint(args...))
}

func Panic(args ...interface{}) {
	col := color.New(color.FgHiBlack, color.BgHiMagenta, color.Bold).SprintfFunc()
	logger.SetPrefix(col("PANIC\t"))
	logger.Println(fmt.Sprint(args...))
}
