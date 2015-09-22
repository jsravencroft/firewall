package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)


func handler(w http.ResponseWriter, r *http.Request, config *ServerConfiguration) {
	socket := strings.Split(r.RemoteAddr, ":")
	ip := socket[0] + "/32"
	fmt.Println("Match alright: ", ip)
	cmd := exec.Command("/sbin/iptables", "-A", "INPUT", "-s", ip,  "-p", "tcp",
	 "-m", "tcp", "--dport", "22", "-m", "state", "--state", "NEW,ESTABLISHED", "-j", "ACCEPT")
	err := cmd.Run()
	fmt.Println(err)
}



func main() {

	config := Config()
	PrintConfiguration(config)

	mux := http.NewServeMux()

	path := "/" + config.secret
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", config.name + "/" + config.version);
		handler(w, r, &config)
	})

	http.ListenAndServe(":" + config.port, mux)
}
