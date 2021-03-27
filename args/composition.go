package oops

type compCliArgs struct {
	satCliArgs
	flowsCliArgs
}

// Constructor
func NewCompCliArgs() *compCliArgs {
	cca := new(compCliArgs)
	return cca
}

// Print
func (cca compCliArgs) ShowDetails() {
	cca.flowsCliArgs.ShowDetails()
	cca.satCliArgs.ShowDetails()
}
