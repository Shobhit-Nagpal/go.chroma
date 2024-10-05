package main

import (
	"fmt"
	"github.com/Shobhit-Nagpal/go.chroma/internal/wav"
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

	wavAudio := wav.NewWav(audio)

  fmt.Println(wavAudio.Channels())
  fmt.Println(wavAudio.Freq())
}
