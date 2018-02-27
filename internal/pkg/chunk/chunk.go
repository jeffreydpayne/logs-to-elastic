package chunk

func Chunk(chunkSize int, length int, f func(int, int)) {
	for i := 0; i < length; i += chunkSize {
		end := i + chunkSize
		if end > length {
			end = length
		}

		f(i, end)
	}
}