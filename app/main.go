package main

import (
	"bufio"
	"context"
	_ "embed"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
	"time"

	sample "sample-japanese-add-rubi"
)

type LineData struct {
	Furigana string
	Surface  string
}

//go:embed assets/result.md
var resultTpl string

func main() {
	var logger io.WriteCloser

	logger, err := os.OpenFile("./debug.log", os.O_APPEND|os.O_WRONLY, 0775)
	if err != nil {
		logger = os.Stdout
	}
	defer logger.Close()

	stdout(logger, "begin analyzing")

	keyfh, err := os.Open("./key.txt")
	if err != nil {
		stderr(logger, "open key file : %v", err)
		return
	}
	defer keyfh.Close()

	scanner := bufio.NewScanner(keyfh)
	if !scanner.Scan() {
		stderr(logger, "read 1 line of key file : %v", scanner.Err())
		return
	}

	target, err := os.Open("./list.txt")
	if err != nil {
		stderr(logger, "read target file : %v", err)
		return
	}
	defer target.Close()

	resultfh, err := os.Create("./result.md")
	if err != nil {
		stderr(logger, "create result file : %v", err)
		return
	}
	defer resultfh.Close()

	ctx := context.Background()
	client, _ := sample.NewClient(scanner.Text())
	defer client.Close()

	var lineData LineData
	var results []LineData
	var furigana, surface strings.Builder
	var req sample.Request
	targetScanner := bufio.NewScanner(target)
	for targetScanner.Scan() {
		req.Text = targetScanner.Text()
		if len(req.Text) == 0 {
			continue
		}

		res, err := client.Analyze(ctx, &req)
		if err == nil {
			if res.Error == nil {
				furigana.Reset()
				surface.Reset()
				for _, word := range res.Result {
					furigana.WriteString(word.Furigana)
					surface.WriteString(word.Surface)
				}
				lineData.Furigana = furigana.String()
				lineData.Surface = surface.String()
			} else {
				lineData.Furigana = "Failed"
				lineData.Surface = req.Text
				stderr(logger, "RES ERROR : %v\n", res.Error)
			}
		} else {
			lineData.Furigana = "Failed"
			lineData.Surface = req.Text
			stderr(logger, "ERROR : %v\n", err)
		}

		results = append(results, lineData)
	}

	tpl, err := template.New("result").Parse(resultTpl)
	if err != nil {
		stderr(logger, "parse template : %v", err)
		return
	}

	err = tpl.ExecuteTemplate(resultfh, "result", map[string]interface{}{
		"results": results,
	})

	if err != nil {
		stderr(logger, "write result : %v", err)
	} else {
		stdout(logger, "successfully finished")
	}
}

func stdout(logger io.Writer, format string, args ...interface{}) {
	writeLog(logger, false, fmt.Sprintf(format, args...))
}

func stderr(logger io.Writer, format string, args ...interface{}) {
	writeLog(logger, true, fmt.Sprintf(format, args...))
}

func writeLog(logger io.Writer, isErr bool, message string) {
	var prefix string
	if isErr {
		prefix = "[ERR]"
	} else {
		prefix = "[OUT]"
	}

	tm := time.Now()
	fmt.Fprintf(logger, "%s %s %s\n", tm.Format("2006/01/02 15:04:05.000"), prefix, message)
}
