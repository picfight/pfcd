package main

import (
	"encoding/json"
	"github.com/jfixby/pin/lang"
	"github.com/picfight/pfcd/pfcdbuilder/policy"
	"io/ioutil"
	"os"
	"testing"
)

func TestPolicy(t *testing.T) {
	p := &policy.PackagePolicy{
		PackageName: "test_pkg",
		Files: []policy.FilePolicy{
			{
				FileName: "file_1.go",
			},
			{
				FileName: "file_2.go",
			},
		},
	}

	bytes, err := json.MarshalIndent(p, "", "	")
	lang.CheckErr(err)

	err = ioutil.WriteFile(POLICY_FILE, bytes, os.FileMode(777))
	lang.CheckErr(err)

}
