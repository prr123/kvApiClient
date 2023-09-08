// apiclogin.go
// http client program to manipulate the azulkv db
// author: prr azul software
// date: 2 Sept 2023
// copyright 2023 prr, azulsoftware
//

package main

import (
    "log"
    "fmt"
    "os"
    "io/ioutil"
    "net/http"

	"bytes"


    util "github.com/prr123/utility/utilLib"
)

func main() {

    numarg := len(os.Args)
    dbg := false

    flags:=[]string{"dbg", "db", "user", "pwd", "port"}


    useStr := "./apiclogin /user=username /pwd=password  [/port=portstr] [/db=database] [/dbg]"
    helpStr := "program to test login via cmd line and json body:\n"

    if numarg > len(flags) +1 {
        fmt.Println("too many arguments in cli!")
		fmt.Printf("usage is: %s\n", useStr)
        os.Exit(-1)
    }

    if numarg == <2 {
		fmt.Printf("insufficient arguments!\n")
		fmt.Printf("usage is: %s\n", useStr)
		os.Exit(1)
	}

	if (numarg == 2) && (os.Args[1] == "help") {
		fmt.Printf("help: %s\n", helpStr)
		fmt.Printf("usage is: %s\n", useStr)
	}


	if dbg {
		fmt.Println("dbg -- flags:")
		for i:=0; i<len(flags); i++ {fmt.Printf(" %d: %s\n", i+1, flags[i])}
	}

    // default file
    flagMap, err := util.ParseFlagsStart(os.Args, flags, 2)
    if err != nil {log.Fatalf("util.ParseFlags: %v\n", err)}

    _, ok := flagMap["dbg"]
    if ok {dbg = true}
    if dbg {
		fmt.Printf("dbg -- flag list:\n")
        for k, v :=range flagMap {
            fmt.Printf("  flag: /%s value: %s\n", k, v)
        }
		fmt.Printf("dbg -- end flag list:\n")
    }

	userStr:=""
    usval, ok := flagMap["user"]
    if ok {
		if usval.(string) == "none" {log.Fatalf("error: user flag requires user name!")}
		userStr = usval.(string)
    } else {
		log.Fatalf("error: user flag is required!")
	}

	pwdStr:=""
    pwdval, ok := flagMap["pwd"]
    if ok {
		if pwdval.(string) == "none" {log.Fatalf("error: pwd flag requires password!")}
		pwdStr = pwdval.(string)
    } else {
		log.Fatalf("error: pwd flag is required!")
	}

	portStr:="10901"
    portval, ok := flagMap["port"]
    if ok {
		if portval.(string) == "none" {log.Fatalf("error: port flag requires value!")}
		portStr = portval.(string)
    } else {
		log.Printf("no port specified! Defaulting to: %s\n", portStr)
	}

	dbPath := "testDb.dat"
    dbval, ok := flagMap["db"]
    if ok {
		if dbval.(string) == "none" {log.Fatalf("error: db flag requires db path!")}
		dbPath = dbval.(string)
    }

	cmdStr := os.Args[1]
	switch cmdStr {
	case "signin":

	case "welcome":

	case "refresh":

	case "logout":

	default:
		log.Fatalf("invalid cmd: %s\n", cmdStr)

	}


	if dbg {

		fmt.Printf("command: %s\n", cmdStr)
		fmt.Printf("dbPath:  %s\n", dbPath)
		fmt.Printf("port:    %s\n", portStr)
		fmt.Printf("user:    %s\n", userStr)
		fmt.Printf("pwd:     %s\n", pwdStr)
		for k, v :=range flagMap {
			if k != "dbg" && k !="db" {
				fmt.Printf("flag: %s value: %s\n", k, v)
			}
		}
	}

	parStr := "user="+userStr+"&pwd="+pwdStr
    url := "http://89.116.30.49:" + portStr + "/" + cmdStr

	if len(parStr) > 0 {
		url += "?" + parStr
	}
    log.Printf("Destination: %s\n", url)

    // Create a Bearer string by appending string access token
	tokStr := "abcdefghijklmnop"
    bearer := "Bearer " + tokStr

	// add body
//	msg := `{"user": "hello, server!"}`

 	jsonBody := []byte(msg)
 	bodyReader := bytes.NewReader(jsonBody)

    // Create a new request using http
    req, err := http.NewRequest("GET", url, bodyReader)

    // add authorization header to the req
    req.Header.Add("Authorization", bearer)

    // add authorization header to the req
//    req.Header.Add("Content-Length", )
 // set the content length
    req.ContentLength = int64(len(msg))

    // Send req using http Client
    client := &http.Client{}

    // synchronous client is blocked until response comes
    resp, err := client.Do(req)
    if err != nil {
        log.Fatalf("Error on response: %v\n", err)
    }
    defer resp.Body.Close()

	// get response code
	log.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))

	// read body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Println("Error while reading the response bytes: %v", err)
    }
    log.Printf("resp body: %s\n", string([]byte(body)))

	log.Println("success end apic!")
}
