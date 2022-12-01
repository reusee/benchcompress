package benchcompress

import (
	"crypto/rand"
	"fmt"
	"testing"

	"github.com/DataDog/zstd"
	"github.com/pierrec/lz4/v4"
)

func benchmarkLZ4Compress(b *testing.B, size int) {
	data := make([]byte, size)
	_, err := rand.Read(data)
	if err != nil {
		panic(err)
	}
	dst := make([]byte, lz4.CompressBlockBound(size))
	var c lz4.Compressor
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n, err := c.CompressBlock(data, dst)
		if err != nil {
			b.Fatal(err)
		}
		if n == 0 {
			b.Fatal()
		}
	}
}

func BenchmarkLZ4Compress(b *testing.B) {
	for i := 128; i < 65536; i *= 2 {
		b.Run(fmt.Sprintf("%d", i), func(b *testing.B) {
			benchmarkLZ4Compress(b, i)
		})
	}
}

func benchmarkDatadogZstdCompress(b *testing.B, size int) {
	data := make([]byte, size)
	_, err := rand.Read(data)
	if err != nil {
		panic(err)
	}
	dst := make([]byte, zstd.CompressBound(size))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res, err := zstd.Compress(data, dst)
		if err != nil {
			b.Fatal(err)
		}
		if len(res) == 0 {
			b.Fatal()
		}
	}
}
func BenchmarkDatadogZstdCompress(b *testing.B) {
	for i := 128; i < 65536; i *= 2 {
		b.Run(fmt.Sprintf("%d", i), func(b *testing.B) {
			benchmarkDatadogZstdCompress(b, i)
		})
	}
}
