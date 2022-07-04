package bite

import (
	"bytes"
	"encoding/binary"
)

type Reader struct {
	Value []byte
}

type Writer struct {
	Value []byte
}

var BinaryEncoding = binary.BigEndian

func (reader *Reader) Read(length int) []byte {
	cursor := reader.Value[:length]
	reader.Value = reader.Value[length:]

	return cursor
}

func (reader *Reader) ReadSingle() byte {
	return reader.Read(1)[0]
}

func NewReader(raw []byte) Reader {
	return Reader{raw}
}

func (writer *Writer) Write(content any) {
	switch any(content).(type) {
	case []byte:
		writer.Value = append(writer.Value, any(content).([]byte)...)

	case byte:
		writer.Value = append(writer.Value, any(content).(byte))

	case string:
		writer.Value = append(writer.Value, []byte(any(content).(string))...)

	case int:
		buffer := new(bytes.Buffer)

		binary.Write(buffer, BinaryEncoding, int8(any(content).(int)))

		writer.Value = append(writer.Value, buffer.Bytes()...)

	default:
		panic("type cannot be written")
	}
}

func (writer *Writer) WriteWithLength(content any, length int) {
	before := len(writer.Value)
	padding := true

	switch any(content).(type) {
	case []byte, byte, string:
		writer.Write(content)

	case int:
		padding = false

		buffer := new(bytes.Buffer)

		binary.Write(buffer, BinaryEncoding, int8(any(content).(int)))

		for len(writer.Value)+buffer.Len() < (before + length) {
			writer.Value = append(writer.Value, 0)
		}

		writer.Value = append(writer.Value, buffer.Bytes()...)
	}

	if padding == true {
		for len(writer.Value) < (before + length) {
			writer.Value = append(writer.Value, 0)
		}
	}
}

func NewWriter() Writer {
	return Writer{}
}
