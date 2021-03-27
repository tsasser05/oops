package main

import (
	"fmt"
	oops "oops/args"
)

func main() {

	// flowsCliArgs
	fmt.Println("\nTesting flowsCliArgs")
	fca := oops.NewFlowsCliArgs()
	fca.SetFlowName("MyFlow")
	fca.ShowDetails()
	flowName := fca.GetFlowName()
	fmt.Printf("\tThe flowName is %s\n", flowName)

	// satCliArgs
	fmt.Println("\nTesting satCliArgs")
	sca := oops.NewSatCliArgs()
	sca.SetSatToken("sat_token_goes_here")
	sca.ShowDetails()
	satToken := sca.GetSatToken()
	fmt.Printf("\tThe Sat Token is %s\n", satToken)

	// composition
	fmt.Println("\nTesting composition")
	cca := oops.NewCompCliArgs()
	cca.SetFlowName("OtherFlowName")
	cca.SetSatToken("big_long_sat_token_goes_here")
	cca.ShowDetails()
}
