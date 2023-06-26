package main

import (
	"os"
	"strings"
)

type File struct {
}

// -------------------------- Write

func (F File) Write(path string, text string) error {
	// create file
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(text)
	if err != nil {
		return err
	}
	return nil
}

func (F File) Append(path string, text string, auto_newline bool) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	defer f.Close()

	// write data
	if auto_newline {
		if _, err = f.WriteString(text + "\n"); err != nil {
			return err
		}
	} else {
		if _, err = f.WriteString(text); err != nil {
			return err
		}
	}

	return nil
}

// - ------------------------- Read

func (F File) Read(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (F File) ReadList(path string) ([]string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")

	return lines, nil
}

func (F File) ReadCSV(path string, separator string) ([][]string, error) {
	var rows [][]string

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		fields := strings.Split(string(line), separator)
		rows = append(rows, fields)
	}
	return rows, nil
}
