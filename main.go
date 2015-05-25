package main

import (
	"flag"
	"fmt"
	"net/http"
	"runtime"
	"strings"

	"github.com/swanwish/go-helper/config"
	"github.com/swanwish/go-helper/logs"
	"github.com/swanwish/go-helper/utils"
	. "github.com/swanwish/sb-admin-go/config"
	"github.com/swanwish/sb-admin-go/handlers"
	"github.com/swanwish/sb-admin-go/views"
)

const (
	DEFAULT_PORT int64 = 8080
)

var (
	port int64
)

func parseCmdLineArgs() {
	flag.Int64Var(&port, "port", DEFAULT_PORT, "The port to listen")
	flag.Parse()
}

func main() {
	config.Load("conf/app.ini")
	viewsConfigFile, err := config.Get("views_config_file")
	if err == nil {
		views.LoadViews(viewsConfigFile)
	}
	appMode, err := config.Get("app_mode")
	if err == nil {
		if appMode == "product" {
			logs.Infof("Application will run in product mode")
			ProductMode = true
		}
	}

	runtime.GOMAXPROCS(runtime.NumCPU())
	parseCmdLineArgs()

	handlers.InitHandlers()

	if port == 0 {
		port = DEFAULT_PORT
	}

	localIps, err := utils.GetLocalIPAddrs()
	if err != nil {
		fmt.Println("Failed to get local ip addresses.")
		return
	}

	fmt.Printf("Service listen on port \x1b[31;1m%d\x1b[0m and server ip addresses are \x1b[31;1m%s\x1b[0m\n", port, strings.Join(localIps, ", "))

	httpAddr := fmt.Sprintf(":%d", port)
	if err := http.ListenAndServe(httpAddr, nil); err != nil {
		fmt.Printf("http.ListendAndServer() failed with %s\n", err)
	}
	fmt.Printf("Exited\n")

}
