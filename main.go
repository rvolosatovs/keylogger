package main

import (
	"flag"
	"log"
	"os"

	evdev "github.com/gvalkov/golang-evdev"
)

func init() {
	log.SetFlags(0)
}

func main() {
	flag.Parse()

	dev, err := evdev.Open(flag.Arg(0))
	if err != nil {
		log.Fatalf("Failed to open device: %s", err)
	}

	for {
		evs, err := dev.Read()
		if err != nil {
			log.Fatalf("Failed to read input: %s", err)
		}

		var s string
		for _, ev := range evs {
			if ev.Type != evdev.EV_KEY {
				continue
			}

			ke := evdev.NewKeyEvent(&ev)
			if ke.State != evdev.KeyDown {
				continue
			}
			s += evdev.KEY[int(ke.Scancode)] + " "
		}

		if len(s) == 0 {
			continue
		}

		b := []byte(s)
		b[len(b)-1] = '\n'

		if _, err := os.Stdout.Write(b); err != nil {
			log.Fatalf("Failed to write keycode to output: %s", err)
		}
	}
}
