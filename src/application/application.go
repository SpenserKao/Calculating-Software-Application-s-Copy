package application

import (
	"errors"
)

type ComputerTypeCount struct {
	LaptopCount  int
	DesktopCount int
}

type application struct {
	m map[string]interface{}
}

func NewUsage() *application {
	a := &application{}
	a.m = make(map[string]interface{})
	return a
}

func (a *application) Contains(key string) bool {
	_, c := a.m[key]
	return c
}

func (a *application) GetVal(key string) (interface{}, error) {
	i, ok := a.m[key]
	if ok {
		return i, nil
	} else {
		return nil, errors.New("Key '" + key + "' NOT in entry ")
	}
}

func (a *application) Add(key string, val interface{}) {
	a.m[key] = val
}

func (a *application) CalculateCopyNumber() int {
	copyTotalNo := 0
	for _, app := range a.m {
		var appx ComputerTypeCount
		// Type Assertion, again.
		appx.DesktopCount = app.(ComputerTypeCount).DesktopCount
		appx.LaptopCount = app.(ComputerTypeCount).LaptopCount
		copyTotalNo += sumUpOneUserID(appx)
	}
	return copyTotalNo
}

func sumUpOneUserID(a ComputerTypeCount) int {
	if a.LaptopCount > 0 && a.DesktopCount > 0 {
		return a.DesktopCount
	} else {
		return a.LaptopCount + a.DesktopCount
	}
}
