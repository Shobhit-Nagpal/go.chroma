package wav

import "fmt"

type Wav struct {
	Id            []byte
	TotalLength   int //Size of file  - 8
	WaveFmt       []byte
	Format        int
	PCM           int16
	Channels      int16
	Freq          int   //Sampling frequency
	BPS           int   //Bytes per second
	BPC           int16 //Bytes per capture
	BitsPerSample int16
	Data          [4]byte
	BytesInData   int
}

const ID_BYTES = 4
const WAVEFMT_BYTES = ID_BYTES + 1 + 8

func NewWav(audio []byte) *Wav {
	return &Wav{
		Id:          audio[:ID_BYTES],
		TotalLength: int(audio[ID_BYTES+1]),
		WaveFmt:     audio[ID_BYTES+1 : WAVEFMT_BYTES],
		Format:      int(audio[WAVEFMT_BYTES+1]),
	}
}

func (w *Wav) StringifyId() {
	fmt.Printf("%s\n", w.Id)
}

func (w *Wav) Length() int {
	return w.TotalLength
}

func (w *Wav) GetFormat() int {
	return w.Format
}

func (w *Wav) GetBPS() int {
	return w.BPS
}

func (w *Wav) GetBPC() int16 {
	return w.BPC
}

func (w *Wav) GetFreq() int {
	return w.Freq
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
