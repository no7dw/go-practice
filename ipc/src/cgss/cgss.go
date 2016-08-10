package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"cg"
	"ipc"
)
var centerClient *cg.centerClient

func startCenterService() error {
	server := ipc.NewIpcServer(&cg.CenterServer())
	client := ipc.NewIpcClient(server)
	centerClient = &cg.centerClient(client);
}