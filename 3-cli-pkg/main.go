package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	printArgs(os.Args)
	flagPgkWithChoices()
	flagComplexSelection()

}

func flagComplexSelection() {
	var containers = map[string]string{
		"abc12": "rabbitmq-3:management",
		"abc13": "rabbitmq-3:management",
		"abc14": "rabbitmq-3:management",
		"abc15": "rabbitmq-3:management",
	}

	// containers -id=abc123
	containersCli := flag.NewFlagSet("containers", flag.ExitOnError)
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatal("<containers> You must specify an option for containers")
		os.Exit(2)
	}

	switch flag.Arg(0) {
	case "containers":
		ID := containersCli.String("id", "", "find a container by id")
		ID2 := containersCli.String("id2", "", "find a second container by id")
		containersCli.Parse(os.Args[2:])

		fmt.Println("=====CONTAINERS======")
		// e.g. containers -id abc12 -id2 abc13
		if *ID != "" {
			fmt.Println(containers[*ID])
		}

		if *ID2 != "" {
			fmt.Println(containers[*ID2])
		}

		if *ID == "" && *ID2 == "" {
			fmt.Println(containers)
		}

		fmt.Println("===========")

	}
}

func flagPgkWithChoices() {

	var (
		num    = flag.Int("n", 1234, "Hello I am help message for you :)") //-n INTEGER, --n INTEGER, --n=INTEGER
		bday   = flag.String("day", strconv.Itoa(time.Now().Local().Day()), "...")
		myBool = flag.Bool("g", false, "...")
	)
	flag.Parse()

	if *num == 1111 {
		fmt.Println("> LEGGO_integer")
	}
	if *bday == "12" {
		fmt.Println("> happy birthday")
	}
	if *myBool {
		fmt.Println("> LEGGO_boolean")
	}

	fmt.Println("======DEFAULTS-OVERWRITTEN VALUES======")
	fmt.Println(*num)
	fmt.Println(*bday)
	fmt.Println(*myBool)
	fmt.Println("=============")
}

func printArgs(args []string) {
	for i, arg := range args {
		fmt.Printf("(%v) %v\n", i, arg)
	}
}
