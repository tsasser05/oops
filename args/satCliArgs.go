package oops

import (
	"fmt"
)

type satCliArgs struct {
	SatToken string
}

// Constructor
func NewSatCliArgs() *satCliArgs {
	sca := new(satCliArgs)
	return sca
}

// Get
func (sca *satCliArgs) GetSatToken() string {
	return sca.SatToken
}

// Set
func (sca *satCliArgs) SetSatToken(st string) {
	sca.SatToken = st
}

// Display()
func (sca satCliArgs) ShowDetails() {
	fmt.Printf("\t%s\n", sca.SatToken)
}
