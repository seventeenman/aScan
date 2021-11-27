package Module

import (
	"github.com/fatih/color"
	"strings"
)

var builder strings.Builder

func LogRate(cho int) {
	if cho == 1 {
		builder.Reset()
		builder.WriteString("[*] ")
		builder.WriteString("Survival detection")
		resOut := builder.String()
		builder.Reset()
		color.Cyan(resOut)
	} else if cho == 2 {
		builder.Reset()
		builder.WriteString("[*] ")
		builder.WriteString("Port detection")
		resOut := builder.String()
		builder.Reset()
		color.Cyan(resOut)
	} else if cho == 3 {
		builder.Reset()
		builder.WriteString("[*] ")
		builder.WriteString("Web detection")
		resOut := builder.String()
		builder.Reset()
		color.Cyan(resOut)
	}
}

func StartOut(address string, runTimes int, mode string, webRunTimes int) {
	color.Magenta(` ▄▄▄        ██████  ▄████▄   ▄▄▄       ███▄    █ 
▒████▄    ▒██    ▒ ▒██▀ ▀█  ▒████▄     ██ ▀█   █ 
▒██  ▀█▄  ░ ▓██▄   ▒▓█    ▄ ▒██  ▀█▄  ▓██  ▀█ ██▒
░██▄▄▄▄██   ▒   ██▒▒▓▓▄ ▄██▒░██▄▄▄▄██ ▓██▒  ▐▌██▒
 ▓█   ▓██▒▒██████▒▒▒ ▓███▀ ░ ▓█   ▓██▒▒██░   ▓██░
 ▒▒   ▓▒█░▒ ▒▓▒ ▒ ░░ ░▒ ▒  ░ ▒▒   ▓▒█░░ ▒░   ▒ ▒ 
  ▒   ▒▒ ░░ ░▒  ░ ░  ░  ▒     ▒   ▒▒ ░░ ░░   ░ ▒░
  ░   ▒   ░  ░  ░  ░          ░   ▒      ░   ░ ░ 
      ░  ░      ░  ░ ░            ░  ░         ░ 
                   ░                             @seventeen`)
	color.Magenta("target: %s  runtimes: %d  mode: %s  webRunTimes: %d", address, runTimes, mode, webRunTimes)
}
