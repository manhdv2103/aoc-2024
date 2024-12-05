package aoc

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Process(solveP1 func(string) string, solveP2 func(string) string, part string, willSubmit bool) {
	day := getDay()
	input := getInput(day)
	var answer string

	var start time.Time
	var elapsed time.Duration
	if part == "1" {
		start = time.Now()
		answer = solveP1(input)
		elapsed = time.Since(start)
	} else {
		start = time.Now()
		answer = solveP2(input)
		elapsed = time.Since(start)
	}

	fmt.Println("Day " + day)
	fmt.Println("Part " + part + ":")
	fmt.Println(answer)
	fmt.Println("\nBenchmark: " + elapsed.String())

	if answer == "" {
		fmt.Println("\nMissing answer")
		return
	}

	copy(answer)
	if willSubmit {
		fmt.Println()
		submit(day, part, answer)
	}
}

func getDay() string {
	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	day, err := strconv.Atoi(filepath.Base(wd)[3:])

	if err != nil {
		panic(err)
	}

	return strconv.Itoa(day)
}

func getInput(day string) string {
	inputFile, err := os.ReadFile("input")

	if errors.Is(err, os.ErrNotExist) {
		return downloadInput(day)
	} else if err != nil {
		panic(err)
	}

	input := string(inputFile)
	if input == "" {
		return downloadInput(day)
	}

	return input
}

func downloadInput(day string) string {
	fmt.Println("Downloading input...")

	cookie := getCookie()
	out, err := exec.Command(
		"curl",
		"-b",
		cookie,
		"-A",
		"curl by manhdv2103@gmail.com",
		"https://adventofcode.com/2024/day/"+day+"/input",
	).Output()

	if err != nil {
		panic(err)
	}

	os.WriteFile("input", out, 0644)
	return string(out)
}

func getCookie() string {
	cookieData, err := os.ReadFile("../cookie")

	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(cookieData))
}

func copy(str string) {
	cmd := exec.Command("xclip", "-selection", "clipboard")
	cmd.Stdin = strings.NewReader(str)
	err := cmd.Run()

	if err != nil {
		panic(err)
	}
}

func submit(day string, part string, answer string) {
	cookie := getCookie()

	fmt.Println("Submitting answer...")
	out, err := exec.Command(
		"curl",
		"-b",
		cookie,
		"-A",
		"curl by manhdv2103@gmail.com",
		"-X",
		"POST",
		"-d",
		"level="+part+"&answer="+answer,
		"https://adventofcode.com/2024/day/"+day+"/answer",
	).Output()

	if err != nil {
		panic(err)
	}

	cmd := exec.Command("xmllint", "--html", "--xpath", "normalize-space(//article/p)", "-")
	cmd.Stdin = bytes.NewReader(out)
	resultOut, err := cmd.Output()

	if err != nil {
		panic(err)
	}

	response := strings.Replace(
		strings.Replace(
			regexp.MustCompile(` \[.*\]`).ReplaceAllString(string(resultOut), ""),
			"; you have to wait after submitting an answer before trying again",
			"",
			1,
		),
		" If you're stuck, make sure you're using the full input data; there are also some general tips on the about page, or you can ask for hints on the subreddit. Please wait one minute before trying again.",
		"",
		1,
	)

	fmt.Println(response)
}
