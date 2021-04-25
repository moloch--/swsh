package main

import (
	_ "embed"
	"strconv"
	"strings"
	"time"

	tm "github.com/buger/goterm"
)

//go:embed film.txt
var film string

func main() {
	film = strings.ReplaceAll(film, "\\n", "\n")
	film = strings.ReplaceAll(film, "\\\\", "\\")
	film = strings.ReplaceAll(film, "\\'", "'")
	frames := strings.Split(film, "\n")
	tm.Clear()

	speed := time.Millisecond * 64

	for index := 0; index+13 < len(frames); index += 14 {
		tm.Clear()
		tm.MoveCursor(1, 1)
		frameControl, err := strconv.Atoi(frames[index])
		if err != nil {
			frameControl = 1
		}
		tm.Printf("%s\n", frames[index+1])
		tm.Printf("%s\n", frames[index+2])
		tm.Printf("%s\n", frames[index+3])
		tm.Printf("%s\n", frames[index+4])
		tm.Printf("%s\n", frames[index+5])
		tm.Printf("%s\n", frames[index+6])
		tm.Printf("%s\n", frames[index+7])
		tm.Printf("%s\n", frames[index+8])
		tm.Printf("%s\n", frames[index+9])
		tm.Printf("%s\n", frames[index+10])
		tm.Printf("%s\n", frames[index+11])
		tm.Printf("%s\n", frames[index+12])
		tm.Printf("%s\n", frames[index+13])
		tm.Flush()

		for index := 0; index < frameControl; index++ {
			time.Sleep(speed)
		}
	}

}
