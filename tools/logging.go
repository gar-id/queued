package tools

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/gar-id/queued/internal/general/config/caches"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func createLog(loglocation string) *os.File {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// Create folder if doesn't exist
	if _, err := os.Stat(loglocation); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(loglocation, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	loglocation = fmt.Sprintf("%v/queued.log", loglocation)
	logfile, err := os.OpenFile(loglocation, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	return logfile
}

func ZapLogger(storeOption string) *zap.Logger {
	logger := coreZap(storeOption)
	defer logger.Sync()

	return logger
}

func coreZap(storeOption string) *zap.Logger {
	// Setup logging
	var stdout, file zapcore.WriteSyncer
	var logLocation, logLevel string
	logLevel = DefaultString(caches.MainConfig.QueueD.Log.Level, "debug")
	logLocation = DefaultString(caches.MainConfig.QueueD.Log.Location, path.Join("/", "var", "log", "queued"))

	if storeOption == "file" {
		logfile := createLog(logLocation)
		file = zapcore.AddSync(logfile)
	} else if storeOption == "console" {
		stdout = zapcore.AddSync(os.Stdout)
	} else if storeOption == "both" {
		logfile := createLog(logLocation)
		file = zapcore.AddSync(logfile)
		stdout = zapcore.AddSync(os.Stdout)
	} else {
		log.Fatal("storeOption is undefined or using unknown string.")
	}

	// Log level
	var level zap.AtomicLevel
	switch logLevel {
	case "debug":
		level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warning":
		level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "panic":
		level = zap.NewAtomicLevelAt(zap.PanicLevel)
	case "fatal":
		level = zap.NewAtomicLevelAt(zap.FatalLevel)
	default:
		level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	productionCfg := zap.NewProductionEncoderConfig()
	productionCfg.TimeKey = "timestamp"
	productionCfg.FunctionKey = "func"
	productionCfg.EncodeDuration = zapcore.MillisDurationEncoder
	productionCfg.EncodeName = zapcore.FullNameEncoder
	productionCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	productionCfg.EncodeCaller = zapcore.FullCallerEncoder

	developmentCfg := zap.NewDevelopmentEncoderConfig()
	developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	developmentCfg.CallerKey = ""
	developmentCfg.EncodeCaller = nil

	consoleEncoder := zapcore.NewConsoleEncoder(developmentCfg)
	fileEncoder := zapcore.NewJSONEncoder(productionCfg)

	var core zapcore.Core
	if storeOption == "file" {
		core = zapcore.NewTee(
			zapcore.NewCore(fileEncoder, file, level),
		)
	} else if storeOption == "console" {
		core = zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, stdout, level),
		)
	} else if storeOption == "both" {
		core = zapcore.NewTee(
			zapcore.NewCore(fileEncoder, file, level),
			zapcore.NewCore(consoleEncoder, stdout, level),
		)
	}

	return zap.New(core, zap.AddCaller(), zap.WithCaller(true))
}
