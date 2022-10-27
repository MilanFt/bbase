package sound

import (
	"errors"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var buffer *beep.Buffer

func InitSound(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return errors.New("error opening checkout sound")
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		return errors.New("error playing checkout sound")
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	buffer = beep.NewBuffer(format)
	buffer.Append(streamer)
	streamer.Close()

	return nil
}

func PlayCheckoutSound() {
	checkout := buffer.Streamer(0, buffer.Len())
	speaker.Play(checkout)
}
