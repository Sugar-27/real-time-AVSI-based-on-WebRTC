package framework

import (
	"errors"
	"fmt"
	"signaling/src/framework/xrpc"
	"strconv"
	"strings"
	"time"
)

var xrpcClients map[string]*xrpc.Client = make(map[string]*xrpc.Client)

func loadXrpc() error {
	sections := configFile.GetSectionList()
	for _, section := range sections {
		if !strings.HasPrefix(section, "xrpc.") {
			continue
		}

		mSection, err := configFile.GetSection(section)
		if err != nil {
			return err
		}

		// server
		values, ok := mSection["server"]
		if !ok {
			return errors.New("no server field in config file")
		}

		arrServer := strings.Split(values, ",")
		// 增加健壮性，防止空格引起歧义
		for i, server := range arrServer {
			arrServer[i] = strings.TrimSpace(server)
		}

		client := xrpc.NewClient(arrServer)
		if values, ok := mSection["connectTimeout"]; ok {
			if connectTimeout, err := strconv.Atoi(values); err == nil {
				client.ConnectTimeout = time.Duration(connectTimeout) * time.Millisecond
			}
		}
		if values, ok := mSection["readTimeout"]; ok {
			if readTimeout, err := strconv.Atoi(values); err == nil {
				client.ReadTimeout = time.Duration(readTimeout) * time.Millisecond
			}
		}
		if values, ok := mSection["writeTimeout"]; ok {
			if writeTimeout, err := strconv.Atoi(values); err == nil {
				client.WriteTimeout = time.Duration(writeTimeout) * time.Millisecond
			}
		}

		xrpcClients[section] = client
	}
	return nil
}

func Call(serviceName string, request interface{}, response interface{}, logId uint32) error {
	fmt.Println("call " + serviceName)
	return nil
}
