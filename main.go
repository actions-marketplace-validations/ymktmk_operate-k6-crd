package main

import (
	"fmt"
	"os"

	"github.com/ymktmk/apply-k6-crd/pkg/k6"
)

func main() {

	method := os.Getenv("INPUT_METHOD")
	if len(method) == 0 {
		fail("the INPUT_METHOD has not been set")
	}

	vus := os.Getenv("INPUT_VUS")
	duration := os.Getenv("INPUT_DURATION")
	rps := os.Getenv("INPUT_RPS")
	parallelism := os.Getenv("INPUT_PARALLELISM")

	templateFile := os.Getenv("INPUT_TEMPLATE_FILE")
	if len(templateFile) == 0 {
		fail("the INPUT_TEMPLATE_FILE has not been set")
	}

	k, err := k6.NewK6(templateFile, vus, duration, rps, parallelism)
	if err != nil {
		fail(err.Error())
	}

	switch method {
	case "create":
		err = k.CreateK6()
		if err != nil {
			fail(err.Error())
		}
	case "delete":
		err = k.DeleteK6()
		if err != nil {
			fail(err.Error())
		}
	default:
		fail("the INPUT_METHOD is create or delete")
	}

}

func fail(err string) {
	fmt.Printf("Error: %s\n", err)
	os.Exit(1)
}
