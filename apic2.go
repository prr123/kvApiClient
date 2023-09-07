// apic.go
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

	"strings"


    util "github.com/prr123/utility/utilLib"
)

func main() {

    numarg := len(os.Args)
    dbg := false

    useStr := "./apic cmd [/kv=key:value]or[/key=key] [/db=database] [/dbg]"
    helpStr := "program to manipulate the azulkv db\ncommands are:\n"
	helpStr = helpStr + "  add /kv=key:value\n"
	helpStr = helpStr + "  del /key=key\n"
	helpStr = helpStr + "  get [/key=key]\n"
	helpStr = helpStr + "  upd /kv=key:value\n"
	helpStr = helpStr + "  list\n"
	helpStr = helpStr + "  info\n"
	helpStr = helpStr + "  entries\n"
	helpStr = helpStr + "  help\n"

    if numarg > 4 {
        fmt.Println("too many arguments in cli!")
		fmt.Printf("usage is: %s\n", useStr)
        os.Exit(-1)
    }

    if numarg == 1 {
		fmt.Printf("insufficient arguments!\n")
		fmt.Printf("usage is: %s\n", useStr)
		os.Exit(1)
	}


	cmdStr := os.Args[1]
    flags:=[]string{"dbg", "db"}

	switch cmdStr {
	case "add":
		if dbg {fmt.Println("add")}
		flags = []string{"kv", "dbg", "db"}
	case "del":
		if dbg {fmt.Println("del")}
		flags = []string{"key", "dbg", "db"}
	case "upd":
		if dbg {fmt.Println("upd")}
		flags = []string{"kv", "dbg", "db"}
	case "get":
		if dbg {fmt.Println("get")}
		flags = []string{"key", "dbg", "db"}
	case "list":
		if dbg {fmt.Println("list")}
	case "info":
		if dbg {fmt.Println("info")}
	case "entries":
		if dbg {fmt.Println("entries")}
	case "help":
		fmt.Printf("%s", helpStr)
		fmt.Printf("usage is: %s\n", useStr)
		os.Exit(1)

	default:
		fmt.Printf(" command %s is not vald!\n For more information, see: azulkv help", cmdStr)
		fmt.Printf("usage is: %s\n", useStr)
		os.Exit(1)
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

	dbPath := "testDb.dat"
    dbval, ok := flagMap["db"]
    if ok {
		if dbval.(string) == "none" {log.Fatalf("error: db flag requires db path!")}
		dbPath = dbval.(string)
    }


	parStr :=""
	keyStr := ""
	valStr := ""

	switch cmdStr {
	case "add":
		kval, ok := flagMap["kv"]
    	if !ok {
        	fmt.Printf("cli add error: no kv flag\n",)
			fmt.Printf("usage is: %s\n", useStr)
			os.Exit(-1)
		} else {
			if kval.(string) == "none" {log.Fatalf("cli add error: no key:val string provided with kv flag!")}
			kvStr := kval.(string)
			idx := strings.Index(kvStr, ":")
			if idx == -1 {log.Fatalf("cli add error: no key:val seperator provided win kv value string!")}
			keyStr = kvStr[:idx]
			valStr = kvStr[idx+1:]
			if dbg {fmt.Printf("-- add key: %s value %s\n", keyStr, valStr)}
		}
		// process add
		log.Printf("processing add key: %s value: %s\n", keyStr, valStr)
		parStr = fmt.Sprintf("?key=%s&val=%s",keyStr, valStr)
	case "upd":
		kval, ok := flagMap["kv"]
    	if !ok {
        	fmt.Printf("cli upd error: no kv flag\n",)
			fmt.Printf("usage is: %s\n", useStr)
			os.Exit(-1)
		} else {
			if kval.(string) == "none" {log.Fatalf("cli upd error: no key:val string provided with kv flag!")}
			kvStr := kval.(string)
			idx := strings.Index(kvStr, ":")
			if idx == -1 {log.Fatalf("cli upd error: no key:val seperator provided win kv value string!")}
			keyStr = kvStr[:idx]
			valStr = kvStr[idx+1:]
			if dbg {fmt.Printf("-- upd key: %s value %s\n", keyStr, valStr)}
		}
		// process upd
		log.Printf("processing upd key: %s value: %s\n", keyStr, valStr)

	case "del":
		kval, ok := flagMap["key"]
    	if !ok {
        	fmt.Printf("cli del error: no key flag\n",)
			fmt.Printf("usage is: %s\n", useStr)
			os.Exit(-1)
		} else {
			if kval.(string) == "none" {log.Fatalf("cli del error: no key string provided with key flag!")}
			keyStr = kval.(string)
			if dbg {fmt.Printf("-- del key: %s\n", keyStr)}
		}

		// process del
		log.Printf("processing del key: %s\n", keyStr)

	case "get":
		kval, ok := flagMap["key"]
    	if !ok {
        	fmt.Printf("cli get error: no key flag\n",)
			fmt.Printf("usage is: %s\n", useStr)
			os.Exit(-1)
		} else {
			if kval.(string) == "none" {log.Fatalf("cli get error: no key string provided with key flag!")}
			keyStr = kval.(string)
			if dbg {fmt.Printf("-- get key: %s\n", keyStr)}
		}

		// process get
		log.Printf("processing get key: %s\n", keyStr)

	case "list":
		// display all keys
		log.Printf("processing list\n")
	case "info":
		// process db info
		log.Printf("processing info\n")
	case "entries":
		// process entries
		log.Printf("processing entries\n")

	default:
		if dbg {fmt.Printf("default cmd: %s\n", cmdStr)}
		for k, _ :=range flagMap {
            if k != "dbg" {
				fmt.Printf("cli error: invalid flag: %s\n",k)
				fmt.Printf("usage is: %s\n", useStr)
				os.Exit(-1)
			}
        }

	}

	if dbg {
		fmt.Printf("dbPath: %s\n", dbPath)
		fmt.Printf("cmd:    %s\n", cmdStr)
		for k, v :=range flagMap {
			if k != "dbg" && k !="db" {
				fmt.Printf("flag: %s value: %s\n", k, v)
			}
		}
	}

	tokStr := "abcdefghijklmnop"
    baseurl := "http://89.116.30.49:10900/db/"

    // Create a Bearer string by appending string access token
    bearer := "Bearer " + tokStr

	url := baseurl + cmdStr + parStr
    log.Printf("Destination: %s\n", url)

    // Create a new request using http
    req, err := http.NewRequest("GET", url, nil)

    // add authorization header to the req
    req.Header.Add("Authorization", bearer)

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
