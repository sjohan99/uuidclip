package main

import (
	"fmt"
	"strings"

	"github.com/atotto/clipboard"
	hook "github.com/robotn/gohook"
)

func getConvertedId(input string) (string, error) {
	if len(input) == 24 && !strings.Contains(input, "-") {
		return convertToUUID(input), nil
	}
	if len(input) == 36 && strings.Count(input, "-") == 4 {
		return convertToObjectId(input), nil
	}
	return "", fmt.Errorf("input does not match either objectId or GUID format")
}

func convertToObjectId(input string) string {
	s := strings.ReplaceAll(input, "-", "")
	return s[:24]
}

func convertToUUID(input string) string {
	return input[0:8] + "-" + input[8:12] + "-" + input[12:16] + "-" + input[16:20] + "-" + input[20:24] + "00000000"
}

func processClipboard() {
	content, err := clipboard.ReadAll()
	if err != nil {
		fmt.Println("Error reading clipboard:", err)
		return
	}

	newContent, err := getConvertedId(content)
	if err != nil {
		fmt.Println("Error transforming text:", err)
		return
	}

	err = clipboard.WriteAll(newContent)
	if err != nil {
		fmt.Println("Error writing to clipboard:", err)
		return
	}
}

func main() {
	s := hook.Start()
	defer hook.End()
	hook.Register(hook.KeyDown, []string{".", "ctrl", "shift"}, func(e hook.Event) {
		processClipboard()
	})
	fmt.Println("Listening... Press Ctrl+Shift+. to convert clipboard content between ObjectId and UUID.")
	<-hook.Process(s)
}
