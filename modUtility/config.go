package modUtility

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// x api key: for operator only;
const CG_Key_HttpPort = "HTTPPORT"

const GD_Default_HttpPort = 10000

var G_HttpPort = 0

func config_Initialize() error {
	portStr := os.Getenv(CG_Key_HttpPort)
	portStr = strings.TrimSpace(portStr)
	fmt.Println("get env httpport: ", portStr)
	if portStr != "" {
		port, err := strconv.Atoi(portStr)
		if err != nil {
			fmt.Println("httpport error : ", portStr, err)
		} else {
			G_HttpPort = port
		}
	}
	if G_HttpPort < 10 {
		portStr = os.Getenv("PORT")
		portStr = strings.TrimSpace(portStr)
		fmt.Println("get env port: ", portStr)
		if portStr != "" {
			port1, err := strconv.Atoi(portStr)
			if err != nil {
				fmt.Println("httpport error : ", portStr, err)
			} else {
				G_HttpPort = port1
			}
		}
	}
	if G_HttpPort < 10 {
		G_HttpPort = GD_Default_HttpPort
	}

	return nil
}
