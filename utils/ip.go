package utils

import (
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
)

// GetIpAddress 获取IP地址信息
func GetIpAddress() (localIp string, err error) {
	localIp = "127.0.0.1"
	if addressList, err := net.InterfaceAddrs(); err != nil {
		return localIp, err
	} else {
		for _, address := range addressList {
			// 检查ip地址判断是否回环地址
			if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
				if ipNet.IP.To4() != nil {
					return ipNet.IP.String(), err
				}
			}
		}
	}

	return localIp, err
}

//GetIPDataInfo 获取ip具体的数据
func GetIPDataInfo(params string) (interface{}, error) {
	db, err := geoip2.Open("./GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ip := net.ParseIP(params)
	record, err := db.City(ip)

	return record, err
}
