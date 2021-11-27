package Module

import (
	"github.com/fatih/color"
	"strconv"
	"time"
)

var startTime time.Time

// 记录开始时间
func RememberTime() {
	startTime = time.Now()
}

// 记录结束时间并计算出所用时间
func CountTime() {
	color.Cyan("[*] Scan end")
	allTime := int(time.Since(startTime) / time.Second)
	builder.Reset()
	builder.WriteString(strconv.Itoa(allTime))
	builder.WriteString(" seconds total")
	timeOut := builder.String()
	builder.Reset()
	color.Blue(timeOut)
}
