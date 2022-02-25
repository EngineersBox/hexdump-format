package hexdump

import "fmt"

func charFormatBytes(data []byte) string {
	buf := ""
	for _, b := range data {
		char := "."
		if b >= 0x20 && b <= 0x7f {
			char = fmt.Sprintf("%c", b)
		}
		buf += char
	}
	return buf
}

func createByteSeparatedLiteral(data []byte) string {
	buf := ""
	for i := 0; i < len(data); i++ {
		buf += fmt.Sprintf(" %02x", data[i])
	}
	return buf
}

func CreateHexdumpText(data []byte) string {
	dwordCount := len(data) / 0x10
	residual := len(data) % 0x10
	buf := `Addr/Off   00 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f  Decoded Text
-----------------------------------------------------------------------------` + "\n"
	for i := 0; i < dwordCount; i++ {
		buf += fmt.Sprintf(
			"%s: %s  %s\n",
			fmt.Sprintf("%08x", i*0x10),
			createByteSeparatedLiteral(data[i*0x10:(i+1)*0x10]),
			charFormatBytes(data[i*0x10:(i+1)*0x10]),
		)
	}
	residualByteString := fmt.Sprintf("%08x: ", (dwordCount)*0x10)
	bytesString := ""
	asciiConverted := ""
	for i := 0; i < residual; i++ {
		bytesString += fmt.Sprintf(" %02x", data[(dwordCount*0x10)+i])
		asciiConverted += charFormatBytes([]byte{data[(dwordCount*0x10)+i]})
	}
	residualByteString += fmt.Sprintf(
		"%-48s  %s",
		bytesString,
		asciiConverted,
	)
	return buf + residualByteString
}
