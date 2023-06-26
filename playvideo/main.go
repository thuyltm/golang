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
	"path/filepath"
	"unsafe"
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
	media, err := NewMedia("demo.mp4")
	handleError(err)
	defer media.Close()

	//Print the media properties
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

// Media is a media file containing
// audio, video and other types of streams.
type Media struct {
	ctx    *C.AVFormatContext
	packet *C.AVPacket
}

// StreamCount returns the number of streams.
func (media *Media) StreamCount() int {
	return int(media.ctx.nb_streams)
}

// FormatName returns the name of the media format.
func (media *Media) FormatName() string {
	if media.ctx.iformat.name == nil {
		return ""
	}

	return C.GoString(media.ctx.iformat.name)
}

// FormatLongName returns the long name
// of the media container.
func (media *Media) FormatLongName() string {
	if media.ctx.iformat.long_name == nil {
		return ""
	}

	return C.GoString(media.ctx.iformat.long_name)
}

// FormatMIMEType returns the MIME type name
// of the media container.
func (media *Media) FormatMIMEType() string {
	if media.ctx.iformat.mime_type == nil {
		return ""
	}

	return C.GoString(media.ctx.iformat.mime_type)
}

// OpenDecode opens the media container for decoding.
//
// CloseDecode() should be called afterwards.
func (media *Media) OpenDecode() error {
	media.packet = C.av_packet_alloc()

	if media.packet == nil {
		return fmt.Errorf(
			"couldn't allocate a new packet")
	}

	return nil
}

// CloseDecode closes the media container for decoding.
func (media *Media) CloseDecode() error {
	C.av_free(unsafe.Pointer(media.packet))
	media.packet = nil

	return nil
}

// Close closes the media container.
func (media *Media) Close() {
	C.avformat_free_context(media.ctx)
	media.ctx = nil
}

// NewMedia returns a new media container analyzer
// for the specified media file.
func NewMedia(filename string) (*Media, error) {
	media := &Media{
		ctx: C.avformat_alloc_context(),
	}

	if media.ctx == nil {
		return nil, fmt.Errorf(
			"couldn't create a new media context")
	}

	fname := C.CString(filename)
	fmt.Println("fname:", fname)
	status := C.avformat_open_input(&media.ctx, fname, nil, nil)
	fmt.Println("status:", status)

	if status < 0 {
		return nil, fmt.Errorf(
			"couldn't open file %s", filename)
	}

	C.free(unsafe.Pointer(fname))

	return media, nil
}
