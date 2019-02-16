package main

type countWriter struct {
	bytesWritten uint64
}

func (w *countWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	w.bytesWritten += uint64(n)
	return n, nil
}
