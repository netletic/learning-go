package reader_test

import (
	"io"
	"testing"
	"testing/iotest"

	"github.com/netletic/reader"
)

// type errReader struct{}

//	func (errReader) Read(p []byte) (n int, err error) {
//		return 0, io.ErrUnexpectedEOF
//	}

func TestReadAll_ReturnsAnyReadError(t *testing.T) {
	// input := errReader{} // commented this out since we use built-in io.ErrUnexpectedEOF now
	// but we could've handrolled this with the errReader struct{} above
	input := iotest.ErrReader(io.ErrUnexpectedEOF)
	_, err := reader.ReadAll(input)
	if err == nil {
		t.Error("want error for broken reader, got nil")
	}
}
