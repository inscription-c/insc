package inscription

import (
	"bytes"
	"io"
)

// Protocol is an interface that defines the methods that a protocol should implement.
// It includes methods for getting the name of the protocol, checking the protocol, getting the length of the protocol, and getting the chunks of the protocol.
type Protocol interface {
	Name() string
	Check() error
	Len() int
	Chunks(size uint64) ([]byte, error)
}

// Reader is a struct that represents a reader.
// It contains the data to be read and a bytes.Reader to read the data.
type Reader struct {
	data   []byte
	reader *bytes.Reader
}

// Reset is a method of the Reader struct.
// It resets the data and the reader of the Reader.
func (r *Reader) Reset(d []byte) {
	r.data = d
	r.reader = nil
}

// Data is a method of the Reader struct.
// It returns the data of the Reader.
func (r *Reader) Data() []byte {
	return r.data
}

// Len is a method of the Reader struct.
// It returns the length of the data of the Reader.
func (r *Reader) Len() int {
	if r.data == nil {
		return 0
	}
	return len(r.data)
}

// Chunks is a method of the Reader struct.
// It returns the chunks of the data of the Reader.
func (r *Reader) Chunks(size uint64) ([]byte, error) {
	if len(r.Data()) == 0 {
		return nil, io.EOF
	}
	if r.reader == nil {
		r.reader = bytes.NewReader(r.data)
	}
	res := make([]byte, size)
	n, err := r.reader.Read(res)
	if err != nil {
		if err == io.EOF {
			if _, err := r.reader.Seek(0, 0); err != nil {
				return nil, err
			}
			return nil, err
		}
		return nil, err
	}
	return res[:n], nil
}

// DefaultProtocol is a struct that represents a default protocol.
// It embeds the Reader struct.
type DefaultProtocol struct {
	Reader
}

// Name is a method of the DefaultProtocol struct.
// It returns the name of the DefaultProtocol.
func (d *DefaultProtocol) Name() string {
	return "default"
}

// Check is a method of the DefaultProtocol struct.
// It checks the DefaultProtocol.
// Currently, it does not perform any operations and always returns nil.
func (d *DefaultProtocol) Check() error {
	return nil
}
