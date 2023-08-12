package main

import (
	"fmt"

	"github.com/NouKinSG/cloud-station-g7/cli"
)

func main() {
	if err := cli.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
