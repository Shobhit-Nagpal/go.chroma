package wav

import "encoding/binary"

const HEADER_LENGTH = 44

func extractAudioData(audio []byte) []int16 {
	audioData := audio[HEADER_LENGTH:]
	audioSamples := make([]int16, len(audioData)/2)

	for i := 0; i < len(audioSamples); i++ {
		audioSamples[i] = int16(binary.LittleEndian.Uint16(audioData[i*2 : (i+1)*2]))
	}

	return audioSamples
}

func normalizeAudio(samples []int16) []float64 {
	normalized := make([]float64, len(samples))
	for i, sample := range samples {
		normalized[i] = float64(sample) / 32768.0
	}
	return normalized
}
