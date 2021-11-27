package Module

import (
	"errors"
	"math"
	"net"
	"sort"
	"strings"
)

var ipTmp = make([]int, 0, 256)
var ipRes = make([]string, 0, 256)
var targetSort map[string]string


func IPString2Int(ip string) (int, error) {
	b := net.ParseIP(ip).To4()
	if b == nil {
		return 0, errors.New("invalid ipv4 format")
	}

	return int(b[3]) | int(b[2])<<8 | int(b[1])<<16 | int(b[0])<<24, nil
}

func Int2IPString(i int) (string, error) {
	if i > math.MaxUint32 {
		return "", errors.New("beyond the scope of ipv4")
	}

	ip := make(net.IP, net.IPv4len)
	ip[0] = byte(i >> 24)
	ip[1] = byte(i >> 16)
	ip[2] = byte(i >> 8)
	ip[3] = byte(i)

	return ip.String(), nil
}

func ResSort(res []string) []string {
	for _, ip4 := range res {
		ipint, _ := IPString2Int(ip4)
		ipTmp = append(ipTmp, ipint)
	}
	sort.Ints(ipTmp)

	for _, ipint := range ipTmp {
		ip4, _ := Int2IPString(ipint)
		ipRes = append(ipRes, ip4)
	}
	return ipRes
}

func TargetSort() map[string]string {
	targetSort = make(map[string]string)
	for _, address := range aliveTarget {
		addTemp := strings.Split(address, ":")
		ip := addTemp[0]
		port := addTemp[1]
		value, ok := targetSort[ip]
		if ok {
			builder.Reset()
			builder.WriteString(value)
			builder.WriteString(",")
			builder.WriteString(port)
			portTmp := builder.String()
			builder.Reset()
			targetSort[ip] = portTmp
		} else {
			targetSort[ip] = port
		}
	}
	return targetSort
}

func TitleSort() map[string]string {
	targetSort = make(map[string]string)
	for _, webtitle := range targetTitle {
		// http://192.168.1.101:80 		<title>nps error</title>	404
		addTemp := strings.Split(webtitle, "-@-;-?-")
		if len(addTemp) == 4  {
			target := addTemp[0]
			title := addTemp[1]
			status := addTemp[2]
			route := addTemp[3]

			builder.Reset()
			builder.WriteString(title)
			builder.WriteString("-@-;-?-")
			builder.WriteString(status)
			builder.WriteString("-@-;-?-")
			builder.WriteString(route)
			res := builder.String()
			builder.Reset()

			targetSort[target] = res
		}
	}
	return targetSort
}