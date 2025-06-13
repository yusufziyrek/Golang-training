package main

import "fmt"

type IEncoder interface {
	Encode(value string)
	Decode(value string)
}

type XEncoder struct {
}

type YEncoder struct {
}

func (xEncoder *XEncoder) Encode(value string) {
	fmt.Println("XEncode ile encode edildi.")
}

func (xEncoder *XEncoder) Decode(value string) {
	fmt.Println("XEncode ile decode edildi.")
}

func (yEncoder *YEncoder) Encode(value string) {
	fmt.Println("YEncode ile encode edildi.")
}

func (yEncoder *YEncoder) Decode(value string) {
	fmt.Println("YEncode ile decode edildi.")
}

func main() {

	// interface referansı
	var encoder IEncoder = &YEncoder{} // Polimorfizmden faydalandık. İster x ister y kullan

	encoder.Encode("123456")
	encoder.Decode("cxaln1234123ld")

}
