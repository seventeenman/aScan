package Module

import (
	"bytes"
	"crypto/tls"
	"github.com/valyala/fasthttp"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var WebPort = [205]string{"80", "81", "82", "83", "84", "85", "86", "87", "88", "89", "90", "91", "92", "98", "99", "443", "800", "801", "808", "880", "888", "889", "1000", "1010", "1080", "1081", "1082", "1118", "1888", "2008", "2020", "2100", "2375", "2379", "3000", "3008", "3128", "3505", "5555", "6080", "6648", "6868", "7000", "7001", "7002", "7003", "7004", "7005", "7007", "7008", "7070", "7071", "7074", "7078", "7080", "7088", "7200", "7680", "7687", "7688", "7777", "7890", "8000", "8001", "8002", "8003", "8004", "8006", "8008", "8009", "8010", "8011", "8012", "8016", "8018", "8020", "8028", "8030", "8038", "8042", "8044", "8046", "8048", "8053", "8060", "8069", "8070", "8080", "8081", "8082", "8083", "8084", "8085", "8086", "8087", "8088", "8089", "8090", "8091", "8092", "8093", "8094", "8095", "8096", "8097", "8098", "8099", "8100", "8101", "8108", "8118", "8161", "8172", "8180", "8181", "8200", "8222", "8244", "8258", "8280", "8288", "8300", "8360", "8443", "8448", "8484", "8800", "8834", "8838", "8848", "8858", "8868", "8879", "8880", "8881", "8888", "8899", "8983", "8989", "9000", "9001", "9002", "9008", "9010", "9043", "9060", "9080", "9081", "9082", "9083", "9084", "9085", "9086", "9087", "9088", "9089", "9090", "9091", "9092", "9093", "9094", "9095", "9096", "9097", "9098", "9099", "9100", "9200", "9443", "9448", "9800", "9981", "9986", "9988", "9998", "9999", "10000", "10001", "10002", "10004", "10008", "10010", "10250", "12018", "12443", "14000", "16080", "18000", "18001", "18002", "18004", "18008", "18080", "18082", "18088", "18090", "18098", "19001", "20000", "20720", "21000", "21501", "21502", "28018", "20880"}
var Charsets = []string{"utf-8", "gbk", "gb2312"}
var webList = make(chan string, 256*len(WebPort))
var titleList = make(chan string, 256*len(WebPort))
var targetTitle []string

func ConstructUrl(target string) string {
	splitTmp := strings.Split(target, ":")
	ip := splitTmp[0]
	port := splitTmp[1]
	protocol := ConstructProtocol(port)
	builder.Reset()
	builder.WriteString(protocol)
	builder.WriteString("://")
	builder.WriteString(ip)
	builder.WriteString(":")
	builder.WriteString(port)
	url := builder.String()
	builder.Reset()
	return url
}

func ConstructProtocol(port string) string {
	if port == "443" {
		return "https"
	} else {
		return "http"
	}
}

func FindTitle(target string) {
	wg.Add(1)
	url := ConstructUrl(target)

	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseResponse(response)
		fasthttp.ReleaseRequest(request)
	}()

	client := &fasthttp.Client{
		Name:                "CLIENT",
		MaxConnsPerHost:     5,
		MaxIdleConnDuration: time.Second * 5,
		ReadBufferSize:      10 * 1024 * 1024,
		WriteBufferSize:     1 * 1024,
		ReadTimeout:         time.Second * 3,
		TLSConfig:           &tls.Config{InsecureSkipVerify: true},
	}

	request.SetRequestURI(url)
	//request.SetConnectionClose()
	request.Header.SetContentType("application/x-www-form-urlencoded")
	request.Header.SetMethod("GET")
	request.Header.Set("User-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:90.0) Gecko/20100101 Firefox/90.0")
	request.Header.Set("Accept", "*/*")

	if err := client.Do(request, response); err != nil {
		if strings.Contains(err.Error(), "timed out") {
			GenerateRes(url, "[timeout]", response.StatusCode(), "")
			return
		} else {
			GenerateRes(url, "[wrong]", response.StatusCode(), "")
			return
		}
	}

	respHeader := response.Header.String()
	respBody := string(response.Body())

	// 对302,301跳转做处理
	if strings.Contains(respHeader, "Location:") || strings.Contains(respBody, "window.location.replace") ||
		strings.Contains(respBody, "window.location.href") || strings.Contains(respBody, "window.navigate") {

		route := ""

		if strings.Contains(respHeader, "Location:") {
			re := regexp.MustCompile("(?ims)^Location: (.*)")
			find := re.FindAllStringSubmatch(string(respHeader), 1)

			if find != nil {
				location := find[0][1]
				if strings.Contains(location, "\r\n") {
					route = strings.Split(location, "\r\n")[0]
				} else {
					route = location
				}
			}
		}

		if strings.Contains(respBody, "window.location.replace") {
			re := regexp.MustCompile("(?ims)window.location.replace\\((.*?)\\);")
			find := re.FindAllStringSubmatch(string(response.Body()), 1)
			if find != nil {
				location := find[0][1]
				route = strings.Trim(location, "'")
				route = strings.Trim(route, "\"")
			}
		}

		if strings.Contains(respBody, "window.location.href") {
			re := regexp.MustCompile("(?ims)window.location.href=(.*?);")
			find := re.FindAllStringSubmatch(string(response.Body()), 1)
			if find != nil {
				location := find[0][1]
				route = strings.Trim(location, "'")
				route = strings.Trim(route, "\"")
			}
		}

		if strings.Contains(respBody, "window.navigate") {
			re := regexp.MustCompile("(?ims)window.navigate\\((.*?)\\);")
			find := re.FindAllStringSubmatch(string(response.Body()), 1)
			if find != nil {
				location := find[0][1]
				route = strings.Trim(location, "'")
				route = strings.Trim(route, "\"")
			}
		}

		route = strings.TrimSpace(route)

		//fmt.Println(route)

		// 判断是否跳转到其它ip
		ipReg := `^((0|[1-9]\d?|1\d\d|2[0-4]\d|25[0-5])\.){3}(0|[1-9]\d?|1\d\d|2[0-4]\d|25[0-5])`
		match1, _ := regexp.MatchString(ipReg, route)

		ipwithprotocolReg := `^(http|https)((0|[1-9]\d?|1\d\d|2[0-4]\d|25[0-5])\.){3}(0|[1-9]\d?|1\d\d|2[0-4]\d|25[0-5])`
		match2, _ := regexp.MatchString(ipwithprotocolReg, route)
		// 判断是否跳转到其它域名
		domainReg := `^(http|https)`
		match3, _ := regexp.MatchString(domainReg, route)

		if match1 || match2 || match3 {
			localurl := route
			request.SetRequestURI(localurl)
			if err := client.Do(request, response); err != nil {
				if strings.Contains(err.Error(), "timed out") {
					GenerateRes(url, "[timeout]", 302, " -> "+route)
					return
				} else {
					GenerateRes(url, "[wrong]", 302, " -> "+route)
					return
				}
			}
			respBodyLocal := string(response.Body())

			encode := GetEncoding(respBodyLocal, response)

			re := regexp.MustCompile("(?ims)<title(.*?)</title>")
			find := re.FindAllStringSubmatch(respBodyLocal, 1)

			if find == nil {
				GenerateRes(url, "[None]", 302, " -> "+route)
			} else {
				titleRaw := find[0][0]
				title := titleRaw
				if encode == "gbk" || encode == "gb2312" {
					titleGBK, err := Decodegbk([]byte(title))
					if err == nil {
						title = string(titleGBK)
					}
				}
				if strings.Contains(title, "\n") {
					title = strings.Split(title, "\n")[0]
				}
				GenerateRes(url, title, 302, " -> "+route)
			}
			return
		}

		if len(route) > 0 {
			if route[:1] != "/" {
				route = "/" + route
			}
		}

		builder.Reset()
		builder.WriteString(url)
		builder.WriteString(route)
		uri := builder.String()
		builder.Reset()
		request.SetRequestURI(uri)
		if err := client.Do(request, response); err != nil {
			if strings.Contains(err.Error(), "timed out") {
				GenerateRes(url, "[timeout]", response.StatusCode(), route)
				return
			} else {
				GenerateRes(url, "[wrong]", response.StatusCode(), route)
				return
			}
		}
		respBody2 := string(response.Body())

		encode := GetEncoding(respBody2, response)

		re := regexp.MustCompile("(?ims)<title(.*?)</title>")
		find := re.FindAllStringSubmatch(respBody2, 1)
		if find == nil {
			GenerateRes(url, "[None]", 0, route)
		} else {
			titleRaw := find[0][0]
			title := titleRaw
			if encode == "gbk" || encode == "gb2312" {
				titleGBK, err := Decodegbk([]byte(title))
				if err == nil {
					title = string(titleGBK)
				}
			}
			if strings.Contains(title, "\n") {
				title = strings.Split(title, "\n")[0]
			}
			GenerateRes(url, title, response.StatusCode(), route)
		}
		return
	}

	respbody := string(response.Body())
	re := regexp.MustCompile("(?ims)<title(.*?)</title>")
	find := re.FindAllStringSubmatch(respbody, 1)

	encode := GetEncoding(respbody, response)

	if find == nil {
		GenerateRes(url, "None", response.StatusCode(), "")
	} else {
		titleRaw := find[0][0]
		title := titleRaw
		if encode == "gbk" || encode == "gb2312" {
			titleGBK, err := Decodegbk([]byte(title))
			if err == nil {
				title = string(titleGBK)
			}
		}
		if strings.Contains(title, "\n") {
			title = strings.Split(title, "\n")[0]
		}
		status := response.StatusCode()
		GenerateRes(url, title, status, "")
	}
}

// 判断编码，防止获取标题乱码
func GetEncoding(respbody string, response *fasthttp.Response) string {
	r1, err := regexp.Compile(`(?im)charset=\s*?([\w-]+)`)
	if err != nil {
		return ""
	}
	headerCharset := r1.FindString(string(response.Header.Peek("Content-Type")))
	if headerCharset != "" {
		for _, v := range Charsets {
			if strings.Contains(strings.ToLower(headerCharset), v) == true {
				return v
			}
		}
	}

	r2, err := regexp.Compile(`(?im)<meta.*?charset=['"]?([\w-]+)["']?.*?>`)
	if err != nil {
		return ""
	}
	htmlCharset := r2.FindString(respbody)
	if htmlCharset != "" {
		for _, v := range Charsets {
			if strings.Contains(strings.ToLower(htmlCharset), v) == true {
				return v
			}
		}
	}
	return ""
}

// GBK解码
func Decodegbk(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func GenerateRes(url string, title string, status int, path string) {
	builder.Reset()
	builder.WriteString(url)
	builder.WriteString("-@-;-?-")
	builder.WriteString(title)
	builder.WriteString("-@-;-?-")
	builder.WriteString(strconv.Itoa(status))
	builder.WriteString("-@-;-?-")
	builder.WriteString(path)
	res := builder.String()
	builder.Reset()
	titleList <- res
}

func ReceiveTitle() {
	for res := range titleList {
		targetTitle = append(targetTitle, res)
		wg.Done()
	}
}

func ScanTitle() {
	for target := range webList {
		FindTitle(target)
		wg.Done()
	}
}

func ControlScan(runTimes int) {
	for runtime := 0; runtime < runTimes; runtime++ {
		go func() {
			ScanTitle()
		}()
	}
}

// 添加扫描目标
func AddScan(target string) {
	splitTmp := strings.Split(target, ":")
	port := splitTmp[1]

	for _, scanPort := range WebPort {
		if scanPort == port {
			wg.Add(1)
			webList <- target
		}
	}
}

// 将所有目标导入
func ImportAll() {
	for _, address := range aliveTarget {
		AddScan(address)
	}
	wg.Wait()
	close(webList)
	close(titleList)
}
