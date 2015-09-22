package main

import(
	"flag"
	"fmt"
)

type ServerConfiguration struct{
	port string
	path string
	name string
	version string
	secret string
}

func Config() ServerConfiguration{
	configuration := ServerConfiguration{
		port: "993",
		name: "JRFW",
		version: "0.1",
		secret: "SUxvdmVUYWt1bWlBbmRIaWthcnU",
	}

        flag.StringVar(&configuration.port, "port", configuration.port, "default listening port")
        flag.Parse()

	return configuration

}	

func PrintConfiguration(configuration ServerConfiguration){
	fmt.Println(configuration)
}
