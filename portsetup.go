// portsetup.go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
	"net/http"
	"flag"
	"github.com/stvp/go-toml-config"
	"os"
	"os/user"
)

var (
	username                = config.String("user", "anonymous")
	address             = config.String("address", "localhost")
	instance            = config.Int("instance", 0)
	port	                = config.Int("port", 0)
)

var cfgFile string
var ipaddr string
var verbose bool
var command string

type TAPinfo struct {
	Tap    string
	Ip     string
	Port   int
	Status string
	Reason string
	Name   string
}


var Usage = func() {
    fmt.Fprintf(os.Stderr, "Usage of %s\n", os.Args[0])
    flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\nConfig file:\naddress = <address to TAPmanager. Default: localhost>\nuser = <user signum. Mandatory>\ninstance = <which user's sim instance. Default:0>\n")
}

func CToGoString(c []byte) string {
    n := -1
    for i, b := range c {
        if b == 0 {
            break
        }
        n = i
    }
    return string(c[:n+1])
}

func main() {
	flag.StringVar(&cfgFile, "c", "portsetup.cfg", "portsetup config file")
	flag.BoolVar(&verbose,"v", false, "Verbose")
	flag.IntVar(port, "p", 0, "Port (MANDATORY)")
	flag.StringVar(&command, "e", "help", "Execute command (NOTE: must be last parameter): \n help\n allocate\n remove\n ip\n port\n list\n")

	
	flag.Usage = Usage
    flag.Parse()	

	usr, err := user.Current()
    if err != nil {
        log.Fatal( err )
    }

	*username = usr.Username
		
	if err := config.Parse(cfgFile); err != nil {
		if (verbose) { fmt.Printf("No %s, using defaults. Err: %s", cfgFile, err.Error()) }
	}
	
	if *port == 0 {
		fmt.Println("You must set port")
		os.Exit(1)
	} 
	
	if command == "help" {
		flag.PrintDefaults()
		os.Exit(0)
	} else
	if command == "allocate" {
		sendstr := fmt.Sprintf("http://%s:%d/allocate/%s_%d",*address,*port,*username,*instance)
//		fmt.Println(sendstr)
		resp, err := http.Get(sendstr)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		var data TAPinfo
		err = json.Unmarshal(body, &data)
    	if err != nil {
    	    fmt.Printf("%T\n%s\n%#v\n",err, err, err)
    	    switch v := err.(type){
    	        case *json.SyntaxError:
    	            fmt.Println(string(body[v.Offset-40:v.Offset]))
    	    }
    	}
		if (verbose) {
			fmt.Printf("Ip:%s Name:%s port:%d Reason:%s Status:%s tap:%s\n",data.Ip, data.Name, data.Port, data.Reason, data.Status, data.Tap)					
		}
		if (data.Status == "OK") {
			os.Exit(0)
		} else
		{
      fmt.Printf("Error in allocate");
			os.Exit(1)
		}
	} else
	if command == "remove" {
		sendstr := fmt.Sprintf("http://%s:%d/remove/%s_%d",*address,*port,*username,*instance)
//		fmt.Println(sendstr)
		resp, err := http.Get(sendstr)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		var data TAPinfo
		err = json.Unmarshal(body, &data)
    	if err != nil {
    	    fmt.Printf("%T\n%s\n%#v\n",err, err, err)
    	    switch v := err.(type){
    	        case *json.SyntaxError:
    	            fmt.Println(string(body[v.Offset-40:v.Offset]))
    	    }
    	}
		if (verbose) {
			fmt.Printf("Ip:%s Name:%s port:%d Reason:%s Status:%s tap:%s\n",data.Ip, data.Name, data.Port, data.Reason, data.Status, data.Tap)					
		}
		if (data.Status == "OK") {
			os.Exit(0)
		} else
		{
			os.Exit(1)
		}
	} else
	if command == "ip" {
		sendstr := fmt.Sprintf("http://%s:%d/ip/%s_%d",*address,*port,*username,*instance)
//		fmt.Println(sendstr)
		resp, err := http.Get(sendstr)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		var data TAPinfo
		err = json.Unmarshal(body, &data)
    	if err != nil {
    	    fmt.Printf("%T\n%s\n%#v\n",err, err, err)
    	    switch v := err.(type){
    	        case *json.SyntaxError:
    	            fmt.Println(string(body[v.Offset-40:v.Offset]))
    	    }
    	}
		if (verbose) {
			fmt.Printf("Ip:%s Name:%s port:%d Reason:%s Status:%s tap:%s\n",data.Ip, data.Name, data.Port, data.Reason, data.Status, data.Tap)					
		}
		if (data.Status == "OK") {
			fmt.Printf("%s\n",data.Ip)
			os.Exit(0)
		} else
		{
			os.Exit(1)
		}
	} else
	if command == "port" {
		sendstr := fmt.Sprintf("http://%s:%d/port/%s_%d",*address,*port,*username,*instance)
//		fmt.Println(sendstr)
		resp, err := http.Get(sendstr)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		var data TAPinfo
		err = json.Unmarshal(body, &data)
    	if err != nil {
    	    fmt.Printf("%T\n%s\n%#v\n",err, err, err)
    	    switch v := err.(type){
    	        case *json.SyntaxError:
    	            fmt.Println(string(body[v.Offset-40:v.Offset]))
    	    }
    	}
		if (verbose) {
			fmt.Printf("Ip:%s Name:%s port:%d Reason:%s Status:%s tap:%s\n",data.Ip, data.Name, data.Port, data.Reason, data.Status, data.Tap)					
		}
		if (data.Status == "OK") {
			fmt.Printf("%d\n",data.Port)
			os.Exit(0)
		} else
		{
			os.Exit(1)
		}
	} else
	if command == "list" {
		sendstr := fmt.Sprintf("http://%s:%d/list",*address,*port)
//		fmt.Println(sendstr)
		resp, err := http.Get(sendstr)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		var data TAPinfo
		
		dec := json.NewDecoder(strings.NewReader(CToGoString(body)))
		for {
			if err := dec.Decode(&data); err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Ip:%s Name:%s port:%d Reason:%s Status:%s tap:%s\n",data.Ip, data.Name, data.Port, data.Reason, data.Status, data.Tap)					
		}
		
	} else
	{
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}	
	
}
