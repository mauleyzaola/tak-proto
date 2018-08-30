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

type target struct {
	commands []string
}

const (
	buildDir = "pb"
	baseDir  = "proto"
)

// Generate protobuf files
func Generate() error {
	if err := os.MkdirAll(buildDir, 0755); err != nil {
		return err
	}
	domainFileOutput := fmt.Sprintf("--go_out=%s", buildDir)
	serviceFileOutput := fmt.Sprintf("--go_out=plugins=grpc:%s", buildDir)
	includePath := fmt.Sprintf("--proto_path=%s", baseDir)

	args := []string{domainFileOutput, includePath}

	// first pass, just read domain files
	files, err := ioutil.ReadDir(baseDir)
	if err != nil {
		return err
	}
	for _, file := range files {
		name := filepath.Join(baseDir, file.Name())
		if strings.Contains(name, "service") {
			continue
		}
		args = append(args, name)
	}
	if err = sh.Run("protoc", args...); err != nil {
		return err
	}

	// second pass, service grpc files
	args = []string{serviceFileOutput, includePath}
	for _, file := range files {
		name := filepath.Join(baseDir, file.Name())
		if !strings.Contains(name, "service") {
			continue
		}
		args = append(args, name)
	}
	if err = sh.Run("protoc", args...); err != nil {
		return err
	}

	return nil
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
