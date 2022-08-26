package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	envs := make(map[string]EnvValue, 0)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		info, err := file.Info()
		if err != nil {
			return nil, err
		}
		if info.Size() <= 0 {
			envs[file.Name()] = EnvValue{NeedRemove: true}
			continue
		}

		f, err := os.Open(dir + "/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}
		envs[file.Name()] = readFromFile(f)
		f.Close()
	}
	return envs, nil
}

func readFromFile(file io.Reader) EnvValue {
	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return EnvValue{
			NeedRemove: true,
		}
	}
	val := scanner.Text()
	val = strings.TrimRight(val, "\t")
	val = strings.TrimRight(val, " ")
	val = strings.Replace(val, "\x00", "\n", 1)
	return EnvValue{
		Value:      val,
		NeedRemove: false,
	}
}
