// SPDX-License-Identifier: Apache-2.0
// Copyright 2021 Authors of KubeArmor

// Package log contains log util wrappers for enhanced pretty logging
package log

import (
	"encoding/json"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ============ //
// == Logger == //
// ============ //

// ZapLogger Handler
var Logger *zap.SugaredLogger

// init Function
func init() {
	initLogger()
}

// customTimeEncoder Function
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000000"))
}

// initLogger Function
func initLogger() {
	defaultConfig := []byte(`{
		"level": "info",
		"encoding": "console",
		"outputPaths": ["stdout"]
	}`)

	config := zap.Config{}
	if err := json.Unmarshal(defaultConfig, &config); err != nil {
		panic(err)
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder   // 修改时间戳的格式
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 日志级别使用大写显示
	config.EncoderConfig = encoderConfig

	// this is not read from config/viper as logger is initialized before config
	if val, ok := os.LookupEnv("DEBUG"); ok && val == "true" {
		config.Level.SetLevel(zap.DebugLevel) // set to enable debug logging
	}

	logger, err := config.Build(zap.AddCaller())
	if err != nil {
		panic(err)
	}

	Logger = logger.Sugar()
}
