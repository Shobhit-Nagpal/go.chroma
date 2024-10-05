package main

import (
	"fmt"
	"log"
	"os"
)

const AUDIO_FILE = "./audio/big-dawgs-hanumankind.wav"

func main() {

	audio, err := os.ReadFile(AUDIO_FILE)
	if err != nil {
		log.Fatalf("%s", err.Error())
		os.Exit(1)
	}

  wavAudio := NewWav()

}
