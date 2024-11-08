package log_util

import (
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"runtime"
	"time"
)

func Init() {
	// Trace > Debug > Info > Warn > Error > Fatal > Panic
	logrus.SetLevel(logrus.InfoLevel)

	// 在输出日志中添加文件名和方法信息
	// logrus.SetReportCaller(true)

	// 指定日志格式
	logrus.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		TimestampFormat: time.DateTime,
		CallerFirst:     true,
		NoColors:        true,
		ShowFullLevel:   true,
		CustomCallerFormatter: func(frame *runtime.Frame) string {
			return fmt.Sprintf(" %s:%d", frame.File, frame.Line)
		},
	})

	// 指定输出流
	logFileName := "logs/info.log"
	logWriter, err := rotatelogs.New(
		logFileName+".%Y%m%d.log",                 // 分割后的文件名称
		rotatelogs.WithLinkName(logFileName),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 设置最大保存时间(7天)
		rotatelogs.WithRotationTime(24*time.Hour), // 设置日志切割时间间隔(1天)
	)
	if err != nil {
		fmt.Println("Failed to setup log rotation:", err)
		os.Exit(1)
	}
	logrus.SetOutput(io.MultiWriter(os.Stdout, logWriter))
}
