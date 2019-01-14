package main

import (
	"fmt"

	sdk "github.com/douglaszuqueto/strava/sdk"
)

func main() {
	fmt.Println("Strava client")

	sdk.New()
}
