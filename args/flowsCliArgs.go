package oops

import "fmt"

type flowsCliArgs struct {
	FlowName string
}

// Constructor
func NewFlowsCliArgs() *flowsCliArgs {
	fca := new(flowsCliArgs)
	return fca
}

// Get
func (fca flowsCliArgs) GetFlowName() string {
	return fca.FlowName
}

// Set
func (fcs *flowsCliArgs) SetFlowName(fn string) {
	fcs.FlowName = fn
}

// Display()
func (fcs flowsCliArgs) ShowDetails() {
	fmt.Printf("\t%s\n", fcs.FlowName)
}
