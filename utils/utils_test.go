package utils

import (
	"fmt"
	"strings"
	"testing"
)

func TestStringToBytes(t *testing.T) {
	b := StringToBytes("ent2015")
	fmt.Println(b)

	s := BytesToString(b)
	fmt.Println(s)
}

func TestInt16ToBytes(t *testing.T) {
	b:=Int16ToBytes(1000)
	fmt.Println(b)
	v := BytesToInt16(b)
	fmt.Println(v)
}

func TestUint8ToBytes(t *testing.T) {
	b := Uint8ToBytes(1)
	fmt.Println(b)
	v := BytesToUInt8(b)
	fmt.Println(v)
}

func TestBytesJoin(t *testing.T) {
	b1 := Int16ToBytes(1000)
	fmt.Println(b1)
	b2 := Uint8ToBytes(1)
	fmt.Println(b2)
	b3 := Uint32ToBytes(3000)
	fmt.Println(b3)
	b := BytesJoin(b1,b2,b3)
	fmt.Println(b)


}
func TestPrintCh(t *testing.T) {
	PrintCh()
	b := []byte("hello我爱你中国！")
	BytesTOCh(b)

	list := strings.FieldsFunc("你好, 我是 李德鹏, 你是 哪位", func(r rune) bool {
		if r == '是'{
			return true
		}
		return false
	})
	fmt.Print(list)
}