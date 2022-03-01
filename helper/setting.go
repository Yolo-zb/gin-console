package helper

import (
	"io/ioutil"
	"golang.org/x/mod/modfile"
)

func GetModuleName() string {
	goModBytes, _ := ioutil.ReadFile("go.mod")
	modName := modfile.ModulePath(goModBytes)
	return modName
}