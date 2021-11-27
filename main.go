package main

import (
	"aScan/Module"
	"flag"
	"github.com/fatih/color"
	"os"
	"strings"
)

var mode string = "top100"
var portRange string = ""
var filename string = ""
var address string = "127.0.0.1"
var runTimes int = 600
var webRunTimes int = 100
var diswebscan int = 0
var disurvivalscan int = 0
var taskList = make(chan string, 65536*255)
var resList = make(chan string, 65536*255)
var checkList = make(chan string, 256)
var resCheck = make(chan string, 256)
var ipList = make([]string, 0, 256)
var aliveList = make([]string, 0, 256)

func main() {
	Menu()
	Module.RememberTime()
	// 存活探测
	if disurvivalscan == 0 {
		if filename == "" {
			ipList = Module.AnalyzeTarget(address, ipList)
		} else {
			ipList = Module.ReadTarget(ipList, filename)
		}
		Module.ConstructCheck(checkList, ipList)
		Module.LogRate(1)
		Module.RunCheck(checkList, resCheck, ipList)
		aliveList = Module.SaveAlive(resCheck, aliveList)
		Module.OutAlive(aliveList)
	} else {
		if filename == "" {
			//color.Red("[-] Must match '-f ips.txt'")
			//os.Exit(1)
			aliveList = Module.AnalyzeTarget(address, aliveList)
		} else {
			aliveList = Module.ReadTarget(aliveList, filename)
		}
		Module.OutAlive(aliveList)
	}
	Module.LogRate(2)

	// 端口探测
	go func() {
		Module.ReceiveResult(resList)
	}()
	Module.ScanTask(runTimes, taskList, resList)
	Module.AddTarget(taskList, resList, aliveList, mode, portRange)
	Module.OutResult()

	// web探测
	Module.LogRate(3)
	if diswebscan != 1 {
		go func() {
			Module.ReceiveTitle()
		}()
		Module.ControlScan(webRunTimes)
		Module.ImportAll()
		Module.OutTitle()
	}

	Module.CountTime()

}

func Menu() {
	flag.StringVar(&address, "t", "127.0.0.1", "Scanned IP address(192.168.1.1 || 192.168.1.1-255 || 192.168.1.1/24), 192.168.1.1/12 is not allowed!")
	flag.StringVar(&portRange, "p", "", "Which ports range to scan(-p 22,23,8080-8081)")
	flag.StringVar(&mode, "m", "top100", "Which ports dict to scan(lite, top100, top1000, all, custom)")
	flag.IntVar(&runTimes, "r", 600, "The number of open coroutines for Port detection")
	flag.IntVar(&webRunTimes, "wr", 100, "The number of open coroutines for Web detection")
	flag.StringVar(&filename, "f", "", "Scan the ip address in the file(-f ips.txt)")
	flag.IntVar(&diswebscan, "dw", 0, "Disable Web Scanning(-dw 1)")
	flag.IntVar(&disurvivalscan, "ds", 0, "Disable Survival Detection(-ds 1)")

	if len(os.Args) == 1 {
		color.Red("Please Run './aScan -h' To See The Usage")
		os.Exit(1)
	}
	flag.Parse()
	if strings.Contains(address, "/") && !strings.Contains(address, "/24") {
		color.Red("[-] 192.168.1.1/12 is not allowed!")
		os.Exit(1)
	}
	Module.StartOut(address, runTimes, mode, webRunTimes)
}
