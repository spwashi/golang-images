package main

import (
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"log"
	"os"
)

func initAudioContext() *audio.Player {
	sampleRate := 44100
	// Create an audio context.
	audioContext := audio.NewContext(sampleRate)

	// Load the sound file.
	file, err := os.Open("assets/bonk.mp3") // Replace with your sound file path
	if err != nil {
		log.Fatal(err)
	}

	// Decode the sound file.
	stream, err := mp3.DecodeWithSampleRate(sampleRate, file)
	if err != nil {
		log.Fatal(err)
	}

	// Create an audio player for the decoded stream.
	audioPlayer, err := audioContext.NewPlayer(stream)
	if err != nil {
		log.Fatal(err)
	}
	return audioPlayer
}
