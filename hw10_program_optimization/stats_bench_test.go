//go:build bench
// +build bench

package hw10programoptimization

import (
	"archive/zip"
	"testing"

	"github.com/stretchr/testify/require"
)

// go test -bench=. -benchmem  > name.txt -tags=bench -run=BenchmarkGetDomainStat
// benchcmp old.txt new.txt

func Benchmark_GetDomainStat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r, err := zip.OpenReader("testdata/users.dat.zip")
		require.NoError(b, err)
		defer r.Close()

		require.Equal(b, 1, len(r.File))

		data, err := r.File[0].Open()
		require.NoError(b, err)

		if _, err := GetDomainStat(data, "com"); err != nil {
			b.Fatal(err)
		}
	}
}
