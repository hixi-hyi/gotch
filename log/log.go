package log

import (
	"fmt"
	"runtime"

	"golang.org/x/net/context"
	aelog "google.golang.org/appengine/log"
)

func Debugf(ctx context.Context, format string, args ...interface{}) {
	aelog.Debugf(ctx, format, args...)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	aelog.Infof(ctx, format, args...)
}

func Warningf(ctx context.Context, format string, args ...interface{}) {
	aelog.Warningf(ctx, fmt.Sprintf("%s\t%s", format, getStackTrace(2)), args...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	aelog.Errorf(ctx, fmt.Sprintf("%s\t%s", format, getStackTrace(2)), args...)
}

func Criticalf(ctx context.Context, format string, args ...interface{}) {
	aelog.Criticalf(ctx, fmt.Sprintf("%s\t%s", format, getStackTrace(2)), args...)
}

func getStackTrace(skip int) string {
	ret := ""
	for i := skip; ; i++ {
		pc, _, _, ok := runtime.Caller(i)
		if !ok {
			break
		}

		fn := runtime.FuncForPC(pc)
		fileName, fileLine := fn.FileLine(pc)
		ret = ret + fmt.Sprintf("%s (%s line %d),", fn.Name(), fileName, fileLine)
	}
	return ret
}
