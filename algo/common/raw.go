package common

import "errors"

type RawDecoder struct {
	file     []byte
	audioExt string
}

func NewRawDecoder(file []byte) Decoder {
	return &RawDecoder{file: file}
}

func (d *RawDecoder) Validate() error {
	for ext, sniffer := range snifferRegistry {
		if sniffer(d.file) {
			d.audioExt = ext
			return nil
		}
	}
	return errors.New("audio doesn't recognized")
}

func (d RawDecoder) Decode() error {
	return nil
}

func (d RawDecoder) GetCoverImage() []byte {
	return nil
}

func (d RawDecoder) GetAudioData() []byte {
	return d.file
}

func (d RawDecoder) GetAudioExt() string {
	return d.audioExt
}

func (d RawDecoder) GetMeta() Meta {
	return nil
}

func init() {
	RegisterDecoder("mp3", NewRawDecoder)
	RegisterDecoder("flac", NewRawDecoder)
	RegisterDecoder("ogg", NewRawDecoder)
	RegisterDecoder("m4a", NewRawDecoder)
	RegisterDecoder("wav", NewRawDecoder)
	RegisterDecoder("wma", NewRawDecoder)
	RegisterDecoder("aac", NewRawDecoder)
}
