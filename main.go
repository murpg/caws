package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
)

type Config struct {
	CommandBoxPath string `json:"commandBoxPath"`
	CfConfig       string `json:"cfconfig"`
}
type program struct{}

const boxBinary = "box.exe"

var (
	pidRegex    = regexp.MustCompile(`.*PID:(\d+).*`)
	statusRegex = regexp.MustCompile(`.*(running|stopped).*`)
)

func main() {
	startCommand := flag.NewFlagSet("start", flag.ExitOnError)
	stopCommand := flag.NewFlagSet("stop", flag.ExitOnError)
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("Commands are 'start' or 'stop'")
		os.Exit(1)
	}

	config, _ := loadConfiguration("config.json")

	switch os.Args[1] {

	case "start":
		startCommand.Parse(os.Args[2:])
		if startCommand.Parsed() {
			pidServerStatus, serverIsRunning := runBoxCommand(config, "server", "status")
			if serverIsRunning {
				fmt.Printf("Server with pid (%d) is already running. Nothing to do", pidServerStatus)
				os.Exit(127)
			}
			runBoxCommand(config, "server", "start")
			runBoxCommand(config, "install", "commandbox-cfconfig")
			runBoxCommand(config, "cfconfig", "import", config.CfConfig)
		}

	case "stop":
		stopCommand.Parse(os.Args[2:])
		if stopCommand.Parsed() {
			runBoxCommand(config, "server", "stop")
		}
	}
}

func loadConfiguration(filename string) (Config, error) {
	var config Config
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err
}

func runBoxCommand(config Config, boxCommands ...string) (int, bool) {
	var pid int
	var serverIsRunning bool
	var box = boxBinary
	if len(config.CommandBoxPath) != 0 {
		box = config.CommandBoxPath + "\\" + boxBinary
	}

	cmd := exec.Command(box, boxCommands...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal("Cannot connect to the command's standard output", err)
		os.Exit(2)
	}
	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			outputLine := scanner.Text()
			matchPid := pidRegex.FindStringSubmatch(outputLine)
			if len(matchPid) > 0 {
				pid, err = strconv.Atoi(matchPid[1])
				if err != nil {
					pid = 0
				}
			}
			matchStatus := statusRegex.FindStringSubmatch(outputLine)
			if len(matchStatus) > 0 {
				if "running" == matchStatus[1] {
					serverIsRunning = true
				} else {
					serverIsRunning = false
				}
			}
			fmt.Println(outputLine)
		}
	}()
	err = cmd.Start()
	if err != nil {
		log.Fatal("Cannot start command", err)
		os.Exit(2)
	}
	err = cmd.Wait()
	if err != nil {
		log.Fatal("Error wating for cmd", err)
		os.Exit(2)
	}

	return pid, serverIsRunning
}
