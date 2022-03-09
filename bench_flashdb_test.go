package flashbench

import (
	"fmt"
	"os"
	"testing"

	"github.com/arriqaaq/flashdb"
)

var (
	fdb      *flashdb.FlashDB
	flashDir string
)

func init() {

	var err error

	flashDir = "testdata/flashDB"
	if err := os.RemoveAll(flashDir); err != nil {
		panic(err)
	}

	if ok := isPathOk("testdata"); !ok {
		if err := os.Mkdir("testdata", os.ModePerm); err != nil {
			panic(err)
		}

	}

	if ok := isPathOk(flashDir); !ok {
		if err := os.Mkdir(flashDir, os.ModePerm); err != nil {
			panic(err)
		}
	}

	opts := flashdb.DefaultConfig()
	opts.EvictionInterval = 0
	opts.NoSync = false

	fdb, err = flashdb.New(opts)
	if err != nil {
		panic(err)
	}
}

func initFlashDB() {
	for n := 0; n < 1000; n++ {
		key := string(getKey(n))
		val := string(geyValue64B())

		if err := fdb.Update(
			func(tx *flashdb.Tx) error {
				err := tx.Set(key, val)
				return err
			}); err != nil {
			panic(err)
		}
	}
}

func BenchmarkFlashDBPutValue64B(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		key := string(getKey(n))
		val := string(geyValue64B())
		if err := fdb.Update(
			func(tx *flashdb.Tx) error {
				return tx.Set(key, val)
			}); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkFlashDBPutValue128B(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		key := string(getKey(n))
		val := string(geyValue128B())
		if err := fdb.Update(
			func(tx *flashdb.Tx) error {
				return tx.Set(key, val)
			}); err != nil {
			b.Fatal(err)
		}
	}

}

func BenchmarkFlashDBPutValue256B(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		key := string(getKey(n))
		val := string(geyValue256B())
		if err := fdb.Update(
			func(tx *flashdb.Tx) error {
				return tx.Set(key, val)
			}); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkFlashDBPutValue512B(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		key := string(getKey(n))
		val := string(geyValue512B())
		if err := fdb.Update(
			func(tx *flashdb.Tx) error {
				return tx.Set(key, val)
			}); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkFlashDBGet(b *testing.B) {
	initFlashDB()

	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		if err := fdb.View(
			func(tx *flashdb.Tx) error {
				key := "key_" + fmt.Sprintf("%07d", 99)
				// member := "member_" + fmt.Sprintf("%07d", 99)
				if _, err := tx.Get(key); err != nil {
					return err
				}
				return nil
			}); err != nil {
			b.Fatal(err)
		}
	}
}
