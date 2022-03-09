package flashbench

// import (
// 	"fmt"
// 	"os"
// 	"testing"

// 	"github.com/xujiajun/nutsdb"
// )

// var (
// 	nutsDB *nutsdb.DB
// 	err    error
// 	dir    string
// )

// func init() {
// 	dir = "testdata/nutsDB"
// 	if err := os.RemoveAll(dir); err != nil {
// 		panic(err)
// 	}

// 	if ok := isPathOk("testdata"); !ok {
// 		if err := os.Mkdir("testdata", os.ModePerm); err != nil {
// 			panic(err)
// 		}

// 	}

// 	if ok := isPathOk(dir); !ok {
// 		if err := os.Mkdir(dir, os.ModePerm); err != nil {
// 			panic(err)
// 		}
// 	}
// }

// func InitNutsDBByDefaultOpt() {
// 	opt := nutsdb.DefaultOptions
// 	opt.Dir = dir
// 	opt.SegmentSize = 64 * 1024 * 1024
// 	nutsDB, err = nutsdb.Open(opt)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func InitNutsDBByByHintKeyOpt() {
// 	opt := nutsdb.DefaultOptions
// 	opt.Dir = "testdata/nutsDB"
// 	opt.SegmentSize = 64 * 1024 * 1024
// 	opt.EntryIdxMode = nutsdb.HintKeyAndRAMIdxMode
// 	nutsDB, err = nutsdb.Open(opt)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func InitNutsDBData() {
// 	for n := 0; n < 10000; n++ {
// 		key := getKey(n)
// 		val := geyValue64B()

// 		if err = nutsDB.Update(
// 			func(tx *nutsdb.Tx) error {
// 				return tx.Put("bucket1", key, val, 0)
// 			}); err != nil {
// 			panic(err)
// 		}
// 	}
// }

// func BenchmarkNutsDBPutValue64B(b *testing.B) {
// 	InitNutsDBByDefaultOpt()

// 	b.ReportAllocs()
// 	b.ResetTimer()

// 	for n := 0; n < b.N; n++ {
// 		key := getKey(n)
// 		val := geyValue64B()
// 		if err = nutsDB.Update(
// 			func(tx *nutsdb.Tx) error {
// 				return tx.Put("bucket1", key, val, 0)
// 			}); err != nil {
// 			b.Fatal(err)
// 		}
// 	}

// }

// func BenchmarkNutsDBPutValue128B(b *testing.B) {
// 	InitNutsDBByDefaultOpt()

// 	b.ReportAllocs()
// 	b.ResetTimer()

// 	for n := 0; n < b.N; n++ {
// 		key := getKey(n)
// 		val := geyValue128B()
// 		if err = nutsDB.Update(
// 			func(tx *nutsdb.Tx) error {
// 				return tx.Put("bucket1", key, val, 0)
// 			}); err != nil {
// 			b.Fatal(err)
// 		}
// 	}

// }

// func BenchmarkNutsDBPutValue256B(b *testing.B) {
// 	InitNutsDBByDefaultOpt()

// 	b.ReportAllocs()
// 	b.ResetTimer()

// 	for n := 0; n < b.N; n++ {
// 		key := getKey(n)
// 		val := geyValue256B()
// 		if err = nutsDB.Update(
// 			func(tx *nutsdb.Tx) error {
// 				return tx.Put("bucket1", key, val, 0)
// 			}); err != nil {
// 			b.Fatal(err)
// 		}
// 	}
// }

// func BenchmarkNutsDBPutValue512B(b *testing.B) {
// 	InitNutsDBByDefaultOpt()

// 	b.ReportAllocs()
// 	b.ResetTimer()

// 	for n := 0; n < b.N; n++ {
// 		key := getKey(n)
// 		val := geyValue512B()
// 		if err = nutsDB.Update(
// 			func(tx *nutsdb.Tx) error {
// 				return tx.Put("bucket1", key, val, 0)
// 			}); err != nil {
// 			b.Fatal(err)
// 		}
// 	}
// }

// func BenchmarkNutsDBGet(b *testing.B) {
// 	InitNutsDBByDefaultOpt()
// 	InitNutsDBData()

// 	b.ReportAllocs()
// 	b.ResetTimer()

// 	for n := 0; n < b.N; n++ {
// 		if err := nutsDB.View(
// 			func(tx *nutsdb.Tx) error {
// 				key := []byte("key_" + fmt.Sprintf("%07d", 99))
// 				if _, err := tx.Get("bucket1", key); err != nil {
// 					return err
// 				}
// 				return nil
// 			}); err != nil {
// 			b.Fatal(err)
// 		}
// 	}
// }

// func BenchmarkNutsDBGetByHintKey(b *testing.B) {
// 	InitNutsDBByByHintKeyOpt()
// 	InitNutsDBData()

// 	b.ReportAllocs()
// 	b.ResetTimer()

// 	for n := 0; n < b.N; n++ {
// 		if err := nutsDB.View(
// 			func(tx *nutsdb.Tx) error {
// 				key := []byte("key_" + fmt.Sprintf("%07d", 99))
// 				if _, err := tx.Get("bucket1", key); err != nil {
// 					return err
// 				}
// 				return nil
// 			}); err != nil {
// 			b.Fatal(err)
// 		}
// 	}
// }
