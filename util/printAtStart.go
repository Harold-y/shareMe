package util

import (
	"fmt"
	"strconv"
	"time"
)

func inform(info string) {
	currentTime := time.Now()
	fmt.Println("[", currentTime.Format("2006.01.02 15:04:05"), "]", info)
}

func InformAtStart() {
	fmt.Println("ShareMe Application Started! Author: Harold Ye")
	inform("Port: " + Port)
	inform("MaxUploadMem: " + strconv.Itoa(MaxUploadSize))
	inform("UploadFolder: " + UploadFolder)
}
