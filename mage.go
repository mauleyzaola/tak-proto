// +build mage

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/magefile/mage/sh"
)

var buildDir = filepath.Join("./build", "domain")

// Generate protobuf files
func Generate() error {
	if err := os.MkdirAll(buildDir, 0755); err != nil {
		return err
	}

	p := "./domain"
	var (
		args1, args2             []string
		protoFiles, serviceFiles []string
	)
	args1 = []string{
		fmt.Sprintf("--go_out=%s", buildDir),
		"--proto_path=" + p,
	}
	args2 = []string{
		fmt.Sprintf("--go_out=plugins=grpc:%s", buildDir),
		"--proto_path=" + p,
	}

	files, err := ioutil.ReadDir(p)
	if err != nil {
		return err
	}
	for _, file := range files {
		name := file.Name()
		fullName := filepath.Join(p, name)
		if strings.Contains(name, "service") {
			serviceFiles = append(serviceFiles, fullName)
		} else if filepath.Ext(name) == ".proto" {
			protoFiles = append(protoFiles, fullName)
		}
	}
	if len(protoFiles) != 0 {
		args1 = append(args1, protoFiles...)
		if err = sh.Run("protoc", args1...); err != nil {
			return err
		}
	}

	if len(serviceFiles) != 0 {
		args2 = append(args2, serviceFiles...)
		err = sh.Run("protoc", args2...)
	}

	return err
}

// Clean protobuf generated files
func Clean() error {
	if _, err := os.Stat(buildDir); err != nil {
		return nil
	}

	files, err := ioutil.ReadDir(buildDir)
	if err != nil {
		return err
	}
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".pb.go") {
			continue
		}
		if err = os.Remove(filepath.Join(buildDir, file.Name())); err != nil {
			return err
		}
	}
	return nil
}
