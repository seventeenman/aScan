package Module

import (
	"bytes"
	"github.com/fatih/color"
	"io/ioutil"
	"net"
	"os/exec"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
	"unsafe"
)

var aliveTarget []string
var portSort []string
var targetAll []string
var CustomPorts = make([]string, 0, 65536)
var wg sync.WaitGroup
var LitePorts = []string{"21", "22", "80", "81", "135", "139", "443", "445", "1433", "3306", "5432", "6379", "7001", "8000", "8080", "8089", "9000", "9200", "11211", "27017"}
var Top100Ports = []string{"21", "22", "23", "25", "53", "69", "79", "80", "81", "82", "83", "84", "85", "86", "87", "88", "89", "110", "111", "135", "139", "143", "161", "322", "389", "443", "445", "465", "512", "513", "514", "515", "524", "587", "873", "888", "993", "995", "999", "1080", "1158", "1352", "1433", "1521", "1863", "2049", "2100", "2181", "2222", "2323", "3128", "3306", "3389", "4848", "4899", "5000", "5061", "5432", "5631", "5632", "5800", "5900", "6379", "7001", "7080", "7090", "8000", "8008", "8009", "8069", "8080", "8081", "8082", "8083", "8084", "8085", "8086", "8087", "8088", "8089", "8090", "8141", "8161", "8291", "8443", "8888", "8899", "8880", "9001", "9080", "9090", "9200", "9300", "9443", "9898", "9900", "17001", "17002", "17003", "11211", "20080", "27017"}
var Top1000Ports = []string{"1", "3", "4", "6", "7", "9", "13", "17", "19", "20", "21", "22", "23", "24", "25", "26", "30", "32", "33", "37", "42", "43", "49", "53", "70", "79", "80", "81", "82", "83", "84", "85", "88", "89", "90", "99", "100", "106", "109", "110", "111", "113", "119", "125", "135", "139", "143", "144", "146", "161", "163", "179", "199", "211", "212", "222", "254", "255", "256", "259", "264", "280", "301", "306", "311", "340", "366", "389", "406", "407", "416", "417", "425", "427", "443", "444", "445", "458", "464", "465", "481", "497", "500", "512", "513", "514", "515", "524", "541", "543", "544", "545", "548", "554", "555", "563", "587", "593", "616", "617", "625", "631", "636", "646", "648", "666", "667", "668", "683", "687", "691", "700", "705", "711", "714", "720", "722", "726", "749", "765", "777", "783", "787", "800", "801", "808", "843", "873", "880", "888", "898", "900", "901", "902", "903", "911", "912", "981", "987", "990", "992", "993", "995", "999", "1000", "1001", "1002", "1007", "1009", "1010", "1011", "1021", "1022", "1023", "1024", "1025", "1026", "1027", "1028", "1029", "1030", "1031", "1032", "1033", "1034", "1035", "1036", "1037", "1038", "1039", "1040", "1041", "1042", "1043", "1044", "1045", "1046", "1047", "1048", "1049", "1050", "1051", "1052", "1053", "1054", "1055", "1056", "1057", "1058", "1059", "1060", "1061", "1062", "1063", "1064", "1065", "1066", "1067", "1068", "1069", "1070", "1071", "1072", "1073", "1074", "1075", "1076", "1077", "1078", "1079", "1080", "1081", "1082", "1083", "1084", "1085", "1086", "1087", "1088", "1089", "1090", "1091", "1092", "1093", "1094", "1095", "1096", "1097", "1098", "1099", "1100", "1102", "1104", "1105", "1106", "1107", "1108", "1110", "1111", "1112", "1113", "1114", "1117", "1119", "1121", "1122", "1123", "1124", "1126", "1130", "1131", "1132", "1137", "1138", "1141", "1145", "1147", "1148", "1149", "1151", "1152", "1154", "1163", "1164", "1165", "1166", "1169", "1174", "1175", "1183", "1185", "1186", "1187", "1192", "1198", "1199", "1201", "1213", "1216", "1217", "1218", "1233", "1234", "1236", "1244", "1247", "1248", "1259", "1271", "1272", "1277", "1287", "1296", "1300", "1301", "1309", "1310", "1311", "1322", "1328", "1334", "1352", "1417", "1433", "1434", "1443", "1455", "1461", "1494", "1500", "1501", "1503", "1521", "1524", "1533", "1556", "1580", "1583", "1594", "1600", "1641", "1658", "1666", "1687", "1688", "1700", "1717", "1718", "1719", "1720", "1721", "1723", "1755", "1761", "1782", "1783", "1801", "1805", "1812", "1839", "1840", "1862", "1863", "1864", "1875", "1900", "1914", "1935", "1947", "1971", "1972", "1974", "1984", "1998", "1999", "2000", "2001", "2002", "2003", "2004", "2005", "2006", "2007", "2008", "2009", "2010", "2013", "2020", "2021", "2022", "2030", "2033", "2034", "2035", "2038", "2040", "2041", "2042", "2043", "2045", "2046", "2047", "2048", "2049", "2065", "2068", "2099", "2100", "2103", "2105", "2106", "2107", "2111", "2119", "2121", "2126", "2135", "2144", "2160", "2161", "2170", "2179", "2190", "2191", "2196", "2200", "2222", "2251", "2260", "2288", "2301", "2323", "2366", "2381", "2382", "2383", "2393", "2394", "2399", "2401", "2492", "2500", "2522", "2525", "2557", "2601", "2602", "2604", "2605", "2607", "2608", "2638", "2701", "2702", "2710", "2717", "2718", "2725", "2800", "2809", "2811", "2869", "2875", "2909", "2910", "2920", "2967", "2968", "2998", "3000", "3001", "3003", "3005", "3006", "3007", "3011", "3013", "3017", "3030", "3031", "3050", "3052", "3071", "3077", "3128", "3168", "3211", "3221", "3260", "3261", "3268", "3269", "3283", "3300", "3301", "3306", "3322", "3323", "3324", "3325", "3333", "3351", "3367", "3369", "3370", "3371", "3372", "3389", "3390", "3404", "3476", "3493", "3517", "3527", "3546", "3551", "3580", "3659", "3689", "3690", "3703", "3737", "3766", "3784", "3800", "3801", "3809", "3814", "3826", "3827", "3828", "3851", "3869", "3871", "3878", "3880", "3889", "3905", "3914", "3918", "3920", "3945", "3971", "3986", "3995", "3998", "4000", "4001", "4002", "4003", "4004", "4005", "4006", "4045", "4111", "4125", "4126", "4129", "4224", "4242", "4279", "4321", "4343", "4443", "4444", "4445", "4446", "4449", "4450", "4451", "4452", "4453", "4454", "4455", "4456", "4457", "4458", "4459", "4460", "4461", "4462", "4463", "4464", "4465", "4466", "4467", "4468", "4469", "4470", "4471", "4472", "4473", "4474", "4475", "4476", "4477", "4478", "4479", "4480", "4481", "4482", "4483", "4484", "4485", "4486", "4487", "4488", "4489", "4490", "4491", "4492", "4493", "4494", "4495", "4496", "4497", "4498", "4499", "4500", "4501", "4502", "4503", "4504", "4505", "4506", "4507", "4508", "4509", "4510", "4511", "4512", "4513", "4514", "4515", "4516", "4517", "4518", "4519", "4520", "4521", "4522", "4523", "4524", "4525", "4526", "4527", "4528", "4529", "4530", "4531", "4532", "4533", "4534", "4535", "4536", "4537", "4538", "4539", "4540", "4541", "4542", "4543", "4544", "4545", "4546", "4547", "4548", "4549", "4550", "4567", "4662", "4848", "4899", "4900", "4998", "5000", "5001", "5002", "5003", "5004", "5009", "5030", "5033", "5050", "5051", "5054", "5060", "5061", "5080", "5087", "5100", "5101", "5102", "5120", "5190", "5200", "5214", "5221", "5222", "5225", "5226", "5269", "5280", "5298", "5357", "5405", "5414", "5431", "5432", "5440", "5500", "5510", "5544", "5550", "5555", "5560", "5566", "5631", "5633", "5666", "5678", "5679", "5718", "5730", "5800", "5801", "5802", "5810", "5811", "5815", "5822", "5825", "5850", "5859", "5862", "5877", "5900", "5901", "5902", "5903", "5904", "5906", "5907", "5910", "5911", "5915", "5922", "5925", "5950", "5952", "5959", "5960", "5961", "5962", "5963", "5987", "5988", "5989", "5998", "5999", "6000", "6001", "6002", "6003", "6004", "6005", "6006", "6007", "6009", "6025", "6059", "6100", "6101", "6106", "6112", "6123", "6129", "6156", "6346", "6389", "6502", "6510", "6543", "6547", "6565", "6566", "6567", "6580", "6646", "6666", "6667", "6668", "6669", "6689", "6692", "6699", "6779", "6788", "6789", "6792", "6839", "6881", "6901", "6969", "7000", "7001", "7002", "7004", "7007", "7019", "7025", "7070", "7100", "7103", "7106", "7200", "7201", "7402", "7435", "7443", "7496", "7512", "7625", "7627", "7676", "7741", "7777", "7778", "7800", "7911", "7920", "7921", "7937", "7938", "7999", "8000", "8001", "8002", "8007", "8008", "8009", "8010", "8011", "8021", "8022", "8031", "8042", "8045", "8080", "8081", "8082", "8083", "8084", "8085", "8086", "8087", "8088", "8089", "8090", "8093", "8099", "8100", "8180", "8181", "8192", "8193", "8194", "8200", "8222", "8254", "8290", "8291", "8292", "8300", "8333", "8383", "8400", "8402", "8443", "8500", "8600", "8649", "8651", "8652", "8654", "8701", "8800", "8873", "8888", "8899", "8880", "8994", "9000", "9001", "9002", "9003", "9009", "9010", "9011", "9040", "9050", "9060", "9043", "9071", "9080", "9081", "9090", "9091", "9099", "9100", "9101", "9102", "9103", "9110", "9111", "9200", "9207", "9220", "9290", "9415", "9418", "9485", "9500", "9443", "9502", "9503", "9535", "9575", "9593", "9594", "9595", "9618", "9666", "9876", "9877", "9878", "9898", "9900", "9917", "9943", "9944", "9968", "9998", "9999", "10000", "10001", "10002", "10003", "10004", "10009", "10010", "10012", "10024", "10025", "10082", "10180", "10215", "10243", "10566", "10616", "10617", "10621", "10626", "10628", "10629", "10778", "11110", "11111", "11967", "12000", "12174", "12265", "12345", "13456", "13722", "13782", "13783", "14000", "14238", "14441", "14442", "15000", "15002", "15003", "15004", "15660", "15742", "16000", "16001", "16012", "16016", "16018", "16080", "16113", "16992", "16993", "17877", "17988", "18040", "18101", "18988", "19101", "19283", "19315", "19350", "19780", "19801", "19842", "20000", "20005", "20031", "20221", "20222", "20828", "21571", "22939", "23502", "24444", "24800", "25734", "25735", "26214", "27000", "27352", "27353", "27355", "27356", "27715", "28201", "30000", "30718", "30951", "31038", "31337", "32768", "32769", "32770", "32771", "32772", "32773", "32774", "32775", "32776", "32777", "32778", "32779", "32780", "32781", "32782", "32783", "32784", "32785", "33354", "33899", "34571", "34572", "34573", "35500", "38292", "40193", "40911", "41511", "42510", "44176", "44442", "44443", "44501", "45100", "48080", "49152", "49153", "49154", "49155", "49156", "49157", "49158", "49159", "49160", "49161", "49163", "49165", "49167", "49175", "49176", "49400", "49999", "50000", "50001", "50002", "50003", "50006", "50300", "50389", "50500", "50636", "50800", "51103", "51493", "52673", "52822", "52848", "52869", "54045", "54328", "55055", "55056", "55555", "55600", "56737", "56738", "57294", "57797", "58080", "60020", "60443", "61532", "61900", "62078", "63331", "64623", "64680", "65000", "65129", "65389"}

func PortScan(target string, resList chan string) {
	conn, err := net.DialTimeout("tcp", target, time.Second*3)
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()
	if err == nil || conn != nil {
		wg.Add(1)
		resList <- target
	}
}

func CmdPing(checkList chan string, resCheck chan string) {
	defer func() {
		wg.Done()
	}()
	sysType := runtime.GOOS

	for {
		host, ok := <-checkList
		if !ok {
			break
		}
		if sysType == "linux" {
			cmd := exec.Command("/bin/sh", "-c", "ping -c 1 -w 1 "+host+">/dev/null && echo true || echo false")
			var out bytes.Buffer
			cmd.Stdout = &out
			cmd.Run()
			if strings.Contains(out.String(), "true") {
				resCheck <- host
			}
		} else if sysType == "windows" {
			cmd := exec.Command("cmd", "/c", "ping -n 1 -w 1 "+host+" && echo true || echo false")
			var out bytes.Buffer
			cmd.Stdout = &out
			cmd.Run()
			if strings.Contains(out.String(), "true") {
				resCheck <- host
			}
		} else if sysType == "darwin" {
			cmd := exec.Command("/bin/bash", "-c", "ping -c 1 -W 1 "+host+" >/dev/null && echo true || echo false")
			var out bytes.Buffer
			cmd.Stdout = &out
			cmd.Run()
			if strings.Contains(out.String(), "true") {
				resCheck <- host
			}
		} else {
			color.Red("[-] Unsupported operating system ")
		}
	}
}

func RunCheck(checkList chan string, resCheck chan string, ipList []string) {
	for pool := 0; pool < len(ipList); pool++ {
		wg.Add(1)
		go CmdPing(checkList, resCheck)
	}
	wg.Wait()
}

// 关闭结果的通道，迭代保存
func SaveAlive(resCheck chan string, aliveList []string) []string {
	close(resCheck)
	for {
		res, ok := <-resCheck
		if !ok {
			break
		}
		aliveList = append(aliveList, res)
	}
	return aliveList
}

func OutAlive(aliveList []string) {
	var ipRes = make([]string, 0, 256)
	ipRes = ResSort(aliveList)
	for _, res := range ipRes {
		builder.Reset()
		builder.WriteString(" [+] ")
		builder.WriteString(res)
		resOut := builder.String()
		builder.Reset()
		color.Red(resOut)
	}
}

// 生成探测的ip
func ConstructCheck(checkList chan string, ipList []string) {
	for _, ip := range ipList {
		checkList <- ip
	}
	close(checkList)
}

// 将输入的target参数解析成ip地址
func AnalyzeTarget(address string, ipList []string) []string {
	if net.ParseIP(address) == nil {
		if strings.Contains(address, "/") {
			ip, ipNet, err := net.ParseCIDR(address)
			if err != nil {
				ip = net.ParseIP(address)
				ipList = append(ipList, ip.String())
			} else {
				for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); incIP(ip) {
					builder.Reset()
					builder.WriteString(ip.String())
					ip2target := builder.String()
					builder.Reset()
					ipList = append(ipList, ip2target)
				}
			}
		} else if strings.Contains(address, "-") {
			// 192.168.0.1-255
			// ipTmp[0]:192.168.0.1 ipTmp[1]:255
			// ipSplit[0]:192 ipSplit[1]:168 ipSplit[2]:0 ipSplit[3]:1
			ipTmp := strings.Split(address, "-")
			ipSplit := strings.Split(ipTmp[0], ".")
			ipStart, _ := strconv.Atoi(ipSplit[3])
			ipOver, _ := strconv.Atoi(ipTmp[1])

			for c := ipStart; c <= ipOver; c++ {
				builder.Reset()
				builder.WriteString(ipSplit[0])
				builder.WriteString(".")
				builder.WriteString(ipSplit[1])
				builder.WriteString(".")
				builder.WriteString(ipSplit[2])
				builder.WriteString(".")
				builder.WriteString(strconv.Itoa(c))
				ip2target := builder.String()
				builder.Reset()
				ipList = append(ipList, ip2target)
			}
		} else {
			color.Red("[-] %s is not an IP Address or CIDR Network", address)
		}
	} else {
		ipList = append(ipList, address)
	}
	return ipList
}

func OutResult() {
	targetRes := make(map[string]string)
	targetRes = TargetSort()
	for _, ip := range ipRes {
		resPort := targetRes[ip]
		if resPort != "" {
			resOut := ""
			if strings.Contains(resPort, ",") {
				res2Int := []int{}
				portSort = strings.Split(resPort, ",")
				res2Int = String2Int(portSort)
				sort.Ints(res2Int)
				for _, port := range res2Int {
					builder.Reset()
					builder.WriteString(" [+] ")
					builder.WriteString(ip)
					builder.WriteString(":")
					builder.WriteString(strconv.Itoa(port))
					resOut = builder.String()
					builder.Reset()
					color.Red(resOut)
					targetAll = append(targetAll, resOut)
				}
			} else {
				builder.Reset()
				builder.WriteString(" [+] ")
				builder.WriteString(ip)
				builder.WriteString(":")
				builder.WriteString(resPort)
				resOut = builder.String()
				builder.Reset()
				color.Red(resOut)
				targetAll = append(targetAll, resOut)
			}
		}
	}
}

func OutTitle() {
	targetRes := make(map[string]string)
	targetRes = TitleSort()
	for _, target := range targetAll {
		target = target[5:]
		url := ConstructUrl(target)
		// <title>nps error</title>-@-;-?-404
		resTitle := targetRes[url]
		if resTitle != "" {
			titleTmp := strings.Split(resTitle, "-@-;-?-")
			// <title>nps error</title> 	404
			title := titleTmp[0]
			status := titleTmp[1]
			route := titleTmp[2]
			builder.Reset()
			builder.WriteString(" [+] Url: ")
			builder.WriteString(url)
			builder.WriteString(route)
			resOut := builder.String()
			builder.Reset()
			color.Red(resOut)
			builder.Reset()
			builder.WriteString("  WebTitle: ")
			builder.WriteString(title)
			resOut = builder.String()
			builder.Reset()
			color.Red(resOut)
			builder.WriteString("  StatusCode: ")
			builder.WriteString(status)
			resOut = builder.String()
			builder.Reset()
			color.Red(resOut)
		}
	}

}

func String2Int(strArr []string) []int {
	res := make([]int, len(strArr))

	for index, val := range strArr {
		res[index], _ = strconv.Atoi(val)
	}

	return res
}

func ReadTarget(ipList []string, filename string) []string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		color.Red("[-] Read failed")
	}
	for _, byteTarget := range bytes.Split(content, []byte("\n")) {
		target := Bytes2String(byteTarget)
		if target != "" {
			ipList = append(ipList, target)
		}
	}
	return ipList
}

// 接受数据
func ReceiveResult(resList chan string) {
	for res := range resList {
		aliveTarget = append(aliveTarget, res)
		wg.Done()
	}
}

func ScanTask(runTimes int, taskList chan string, resList chan string) {
	for pool := 0; pool < runTimes; pool++ {
		go func() {
			for host := range taskList {
				PortScan(host, resList)
				wg.Done()
			}
		}()
	}
}

// 根据mode生成对应的address
func AddTarget(taskList chan string, resList chan string, aliveList []string, mode string, portRange string) {
	for _, ip := range aliveList {
		if mode == "all" {
			for port := 0; port < 65536; port++ {
				wg.Add(1)
				builder.Reset()
				builder.WriteString(ip)
				builder.WriteString(":")
				builder.WriteString(strconv.Itoa(port))
				target := builder.String()
				builder.Reset()
				taskList <- target
			}
		} else if mode == "lite" {
			for port := 0; port < len(LitePorts); port++ {
				wg.Add(1)
				builder.Reset()
				builder.WriteString(ip)
				builder.WriteString(":")
				builder.WriteString(LitePorts[port])
				target := builder.String()
				builder.Reset()
				taskList <- target
			}
		} else if mode == "top100" {
			for port := 0; port < len(Top100Ports); port++ {
				wg.Add(1)
				builder.Reset()
				builder.WriteString(ip)
				builder.WriteString(":")
				builder.WriteString(Top100Ports[port])
				target := builder.String()
				builder.Reset()
				taskList <- target
			}
		} else if mode == "top1000" {
			for port := 0; port < len(Top1000Ports); port++ {
				wg.Add(1)
				builder.Reset()
				builder.WriteString(ip)
				builder.WriteString(":")
				builder.WriteString(Top1000Ports[port])
				target := builder.String()
				builder.Reset()
				taskList <- target
			}
		} else if mode == "custom" {
			AnalyzePort(portRange)
			for port := 0; port < len(CustomPorts); port++ {
				wg.Add(1)
				builder.Reset()
				builder.WriteString(ip)
				builder.WriteString(":")
				builder.WriteString(CustomPorts[port])
				target := builder.String()
				builder.Reset()
				taskList <- target
			}
		}
	}
	wg.Wait()
	close(taskList)
	close(resList)
}

// 用户指定扫描端口范围(1,2,3-5,10), 要在AddTarget之前调用，以生成CustomPorts端口字典
func AnalyzePort(portRange string) {
	if strings.Contains(portRange, ",") {
		portTmp := strings.Split(portRange, ",")
		// portTmp: 1 2 3-5 10
		for _, atom := range portTmp {
			if strings.Contains(atom, "-") {
				portTmp := strings.Split(atom, "-")
				portStart, _ := strconv.Atoi(portTmp[0])
				portOver, _ := strconv.Atoi(portTmp[1])
				for port := portStart; port < portOver; port++ {
					CustomPorts = append(CustomPorts, strconv.Itoa(port))
				}
			} else {
				CustomPorts = append(CustomPorts, atom)
			}
		}
	} else if strings.Contains(portRange, "-") && !strings.Contains(portRange, ",") {
		portTmp := strings.Split(portRange, "-")
		portStart, _ := strconv.Atoi(portTmp[0])
		portOver, _ := strconv.Atoi(portTmp[1])
		for port := portStart; port < portOver; port++ {
			CustomPorts = append(CustomPorts, strconv.Itoa(port))
		}
	} else {
		CustomPorts = append(CustomPorts, portRange)
	}
	CustomPorts = RemoveDuplication_map(CustomPorts)
}

// 端口去重
func RemoveDuplication_map(arr []string) []string {
	set := make(map[string]struct{}, len(arr))
	j := 0
	for _, v := range arr {
		_, ok := set[v]
		if ok {
			continue
		}
		set[v] = struct{}{}
		arr[j] = v
		j++
	}
	return arr[:j]
}

func incIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func String2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
