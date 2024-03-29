// an mp3 player implementation for Horu, made in Golang by Lorenzo Di Cola;
// thanks to faiface project for the ispiration

package player

import (
	"fmt"
	"os"
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func Song() {

	for {
		
		f, _ := os.Open("Highway To Hell.mp3")
		streamer, format, _ := mp3.Decode(f)
		defer streamer.Close()

		speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/3))
		speaker.Play(streamer)
		
		stop := "" 
		fmt.Scan(&stop)

		if stop == "stop" || stop == "ferma" || stop == "basta" {
			break
		}

		select {}
	}
}
