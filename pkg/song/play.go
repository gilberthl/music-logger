package song

import (
	"fmt"
	"github.com/kbinani/midi"
	"github.com/rakyll/portmidi"
	"log"
	"os"
	"time"
)

func Start(deviceID portmidi.DeviceID, sig <-chan interface{}) {
	portmidi.Initialize()
	fmt.Printf("Number of MIDI devices: %#v \n", portmidi.Info(deviceID))

	out, err := portmidi.NewOutputStream(deviceID, 1024, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	playSecondNote := func(d time.Duration) {
		out.WriteShort(0x90, 61, 100)
		time.Sleep(time.Millisecond * 100)
		out.WriteShort(0x80, 61, 100)
		time.Sleep(d)
	}

	for {
		start := time.Now()
		go func() {
			out.WriteShort(0x90, 60, 100)
			time.Sleep(time.Millisecond * 500)
			out.WriteShort(0x80, 60, 100)
		}()

		select {
		case <-sig:
			fmt.Println("SIG")
			playSecondNote(time.Now().Sub(start))
		case <-time.After(time.Second * 2):
			playSecondNote(time.Second * 2)
		}
	}
}

func GetSMF() []midi.Track {
	f, err := os.Open("../data/Tetris - Tetris Main Theme.mid")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	file, err := midi.Read(f)
	if err != nil {
		panic(err)
	}
	return file.Tracks

	// for i, track := range file.Tracks {
	// 	fmt.Printf("track#%d: %5d events\n", i, len(track.Events))
	// 	for _, event := range track.Events {
	// 		go func(event midi.Event) {
	// 			time.Sleep(time.Millisecond * time.Duration(event.Tick))
	// 			ints := GetMessageInts(event)
	// 			fmt.Println(ints)
	// 			// out.WriteSysExBytes(portmidi.Timestamp(event.Tick), event.Messages)
	// 			out.WriteShort(0x90, ints[1], ints[2])
	// 		}(event)
	// 	}
	// }
}

func GetMessageInts(e midi.Event) [3]int64 {
	bytes := e.Messages
	ints := [3]int64{}
	for i, b := range bytes {
		if i < 3 {
			ints[i] = int64(b)
		}
	}
	return ints
}
