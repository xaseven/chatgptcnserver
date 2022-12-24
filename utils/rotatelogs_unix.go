// +build !windows

package utils

import (
  "chatgptcnserver/global"
  zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
  "go.uber.org/zap/zapcore"
  "os"
  "path"
  "time"
)

// GetWriteSyncer zap logger中加入file-rotatelogs
func GetWriteSyncer() (zapcore.WriteSyncer, error) {
  fileWriter, err := zaprotatelogs.New(
    path.Join(global.Chatgptcn_CONFIG.Zap.Director, "%Y-%m-%d.log"),
    zaprotatelogs.WithLinkName(global.Chatgptcn_CONFIG.Zap.LinkName),
    zaprotatelogs.WithMaxAge(365*24*time.Hour),
    zaprotatelogs.WithRotationTime(24*time.Hour),
  )
  if global.Chatgptcn_CONFIG.Zap.LogInConsole {
    return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
  }
  return zapcore.AddSync(fileWriter), err
}
