package logger

import (
	"fmt"
	"log"
)

const (
	Black = (iota + 30)
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

// Returns a proper string to output for colored logging
func colorString(text string, color int) string {
	return fmt.Sprintf("\033[%dm%s\033[m", int(color), text)
}

func Info(text string) {
	text = "[INFO]: " + text
	log.Println(colorString(text, Cyan))
}

func Warn(text string) {
	text = "[WARNING]: " + text
	log.Println(colorString(text, Yellow))
}

func Debug(text string) {
	text = "[DEBUG]: " + text
	log.Println(colorString(text, Magenta))
}

func Success(text string) {
	text = "[INFO]: " + text
	log.Println(colorString(text, Green))
}

func Error(text string) {
	text = "[ERROR]: " + text
	log.Println(colorString(text, Red))
}
