package main

// #cgo pkg-config: libavformat libavcodec libavutil libswscale
// #include <libavcodec/avcodec.h>
// #include <libavformat/avformat.h>
// #include <libavutil/avconfig.h>
// #include <libswscale/swscale.h>
// #include <libavcodec/bsf.h>
import "C"
import (
	"fmt"
	"github.com/zergon321/reisen"
	"path/filepath"
)

func main() {
	// check video path exist
	const config = "demo.mkv"
	_, err := filepath.Abs(config)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("File exists")
	// Open the media file by its name
	media, err := reisen.NewMedia("demo.mp4")
	handleError(err)
	defer media.Close()
	dur, err := media.Duration()
	handleError(err)

	//Print the media properties
	fmt.Println("Duration:", dur)
	fmt.Println("Format name:", media.FormatName())
	fmt.Println("Format long name", media.FormatLongName())
	fmt.Println("MIME type:", media.FormatMIMEType())
	fmt.Println("Number of streams:", media.StreamCount())
	fmt.Println()
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
