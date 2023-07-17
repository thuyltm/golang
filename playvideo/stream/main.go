package main

import (
	"github.com/faiface/beep"
	"github.com/hajimehoshi/ebiten"
	"github.com/zergon321/reisen"
	"image"
	"time"
)

const (
	width                             = 1280
	height                            = 720
	frameBufferSize                   = 1024
	sampleRate                        = 44100
	channelCount                      = 2
	bitDepth                          = 8
	sampleBufferSize                  = 32 * channelCount * bitDepth * 1024
	SpeakerSampleRate beep.SampleRate = 44100
)

// readVideoAndAudio reads video and audio frames
// from the opened media and sends the decoded
// data to the channels to be played
func readVideoAndAudio(media *reisen.Media) (<-chan *image.RGBA,
	<-chan [2]float64, chan error, error) {
	frameBuffer := make(chan *image.RGBA, frameBufferSize)
	sampleBuffer := make(chan [2]float64, sampleBufferSize)
	errs := make(chan error)
	err := media.OpenDecode()
	if err != nil {
		return nil, nil, nil, err
	}
	videoStream := media.VideoStreams()[0]
	err = videoStream.Open()
	if err != nil {
		return nil, nil, nil, err
	}
	audioStream := media.AudioStreams()[0]
	err = audioStream.Open()

	if err != nil {
		return nil, nil, nil, err
	}
	go func() {
		for {

		}
	}()
	return frameBuffer, sampleBuffer, errs, nil
}

//streamSamples creates a new custom streamer for
//playing audio samples provided by the source channel
func streamSamples(sampleSource <-chan [2]float64) beep.Streamer {
	return beep.StreamerFunc(func(samples [][2]float64) (n int, ok bool)) {
		numRead := 0
		for i := 0; i < len(samples); i++ {

		}

		if numRead < len(samples) {
			return numRead, false
		}
		return numRead, true
	}
}

type Game struct {
	videoSprite            *ebiten.Image
	ticker                 <-chan time.Time
	errs                   <-chan error
	frameBuffer            <-chan *image.RGBA
	fps                    int
	videoTotalFramesPlayed int
	videoPlaybackFPS       int
	perSecond              <-chan time.Time
	last                   time.Time
	deltaTime              float64
}

func main() {
	game := &Game{}
	err := game.Start("rtmp://localhost/live/stream")
	handleError(err)
	ebiten.SetWindowSize(width, height)

}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
