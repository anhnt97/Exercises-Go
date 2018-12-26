package main
import (
	"testing"
	// "bytes"
	// "image/png"
	"errors"
)
type ErrorWriter struct {

}
func (e *ErrorWriter) Write(b []byte) (int ,error) {
	return 0, errors.New("Expected error")
}
// func TestQRCodePropagatesErrors(t *testing.T)  {
// 	w := new(ErrorWriter)
// 	err := GenerateQRCode(w, "555-2368")
// 	if err == nil || err.Error() != "Expected error" {
// 		t.Errorf("Error not propagated correctly, got %v", err)
// 	}	
// }
// func TestQRCodeReturnsValue(t *testing.T){
// 	buffer := new(bytes.Buffer)
// 	GenerateQRCode(buffer, "555-2368")
// 	if buffer.Len() == 0 {
// 		t.Errorf("No QRCode generated")
// 	}
// }

// func TestQRCodeGeneratesPNG(t *testing.T){
// 	buffer := new(bytes.Buffer)
// 	GenerateQRCode(buffer, "555-2368")
// 	_, err := png.Decode(buffer)
// 	if err != nil {
// 		t.Errorf("Generated is not a PNG: %s",err)
// 	}
// }
func TestVersionDeterminesSize(t *testing.T){
	table := []struct {
		version int
		expected int
	}{
		{1, 21},
		{2, 25},
		{6, 41},
		{7, 45},
		{14, 73},
		{40, 177},
	}
	for _, test := range table {
		size := Version(test.version).PatternSize()
		if size != test.expected {
			t.Errorf("Version %d, expected %d but got %d", test.version, test.expected, size)
		}
	}
}