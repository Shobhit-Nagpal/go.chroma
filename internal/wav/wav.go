package wav

import (
	"encoding/binary"
	"fmt"
)

type Wav struct {
	id            []byte // RIFF
	totalLength   uint32 //Size of file
	fileType      string //"WAVE"
	formatLength  uint16
	formatType    uint16 // 1 is PCM
	channels      uint16
	freq          uint32 //Sampling frequency
	byteRate      uint32
	blockAlign    uint16
	bps           uint16 //Bytes per sample
	bpc           int16  //Bytes per capture
	bitsPerSample int16
	dataSize      uint32
	data          []float64 //Normalized audio between 1 and -1
}

func NewWav(audio []byte) *Wav {

	audioSamples := extractAudioData(audio)
	normalizedAudio := normalizeAudio(audioSamples)

	return &Wav{
		id:           audio[:4],
		totalLength:  binary.LittleEndian.Uint32(audio[4:8]),
		fileType:     string(audio[8:12]),
		formatLength: binary.LittleEndian.Uint16(audio[16:20]),
		formatType:   binary.LittleEndian.Uint16(audio[20:22]),
		channels:     binary.LittleEndian.Uint16(audio[22:24]),
		freq:         binary.LittleEndian.Uint32(audio[24:28]),
		byteRate:     binary.LittleEndian.Uint32(audio[28:32]),
		blockAlign:   binary.LittleEndian.Uint16(audio[32:34]),
		bps:          binary.LittleEndian.Uint16(audio[34:36]),
		dataSize:     binary.LittleEndian.Uint32(audio[40:44]),
		data:         normalizedAudio,
	}
}

func (w *Wav) StringifyId() string {
	return fmt.Sprintf("%s", w.id)
}

func (w *Wav) FileType() string {
	return w.fileType
}

func (w *Wav) Length() uint32 {
	return w.totalLength
}

func (w *Wav) FormatLength() uint16 {
	return w.formatLength
}

func (w *Wav) BPS() uint16 {
	return w.bps
}

func (w *Wav) BPC() int16 {
	return w.bpc
}

func (w *Wav) Freq() uint32 {
	return w.freq
}

func (w *Wav) FormatType() uint16 {
	return w.formatType
}

func (w *Wav) Channels() uint16 {
	return w.channels
}

func (w *Wav) DataSize() uint32 {
	return w.dataSize
}

func (w *Wav) ByteRate() uint32 {
	return w.byteRate
}

func (w *Wav) BlockAlign() uint16 {
	return w.blockAlign
}

func (w *Wav) AudioData() []float64 {
	return w.data
}

//Positions   Sample Value         Description
//1 - 4       "RIFF"               Marks the file as a riff file. Characters are each 1. byte long.
//5 - 8       File size (integer)  Size of the overall file - 8 bytes, in bytes (32-bit integer). Typically, you'd fill this in after creation.
//9 -12       "WAVE"               File Type Header. For our purposes, it always equals "WAVE".
//13-16       "fmt "               Format chunk marker. Includes trailing null
//17-20       16                   Length of format data as listed above
//21-22       1                    Type of format (1 is PCM) - 2 byte integer
//23-24       2                    Number of Channels - 2 byte integer
//25-28       44100                Sample Rate - 32 bit integer. Common values are 44100 (CD), 48000 (DAT). Sample Rate = Number of Samples per second, or Hertz.
//29-32       176400               (Sample Rate * BitsPerSample * Channels) / 8.
//33-34       4                    (BitsPerSample * Channels) / 8.1 - 8 bit mono2 - 8 bit stereo/16 bit mono4 - 16 bit stereo
//35-36       16                   Bits per sample
//37-40       "data"               "data" chunk header. Marks the beginning of the data section.
//41-44       File size (data)     Size of the data section, i.e. file size - 44 bytes header.
