package bite

type Reader struct {
	Value []byte
}

func (reader *Reader) Read(length int) []byte {
	cursor := reader.Value[:length]
	reader.Value = reader.Value[length:]

	return cursor
}

func New(raw []byte) Reader {
	return Reader{raw}
}