package flashbench

// import (
// 	"fmt"
// 	"os"
// 	"testing"

// 	"github.com/boltdb/bolt"
// )

// var boltDB *bolt.DB

// func init() {
// 	dir := "testdata/boltDB"
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

// 	boltDB, err = bolt.Open("testdata/boltDB/boltDB.db", 0666, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func BenchmarkBoltDBPutValue64B(b *testing.B) {
// 	b.ReportAllocs()
// 	b.ResetTimer()

// 	for n := 0; n < b.N; n++ {
// 		key := getKey(n)
// 		val := geyValue64B()
// 		if err = boltDB.Update(func(tx *bolt.Tx) error {
// 			bucket, err := tx.CreateBucketIfNotExists([]byte("bucket1"))
// 			if err != nil {
// 				return err
// 			}

// 			return bucket.Put([]byte(key), val)
// 		}); err != nil {
// 			b.Fatal(err)
// 		}
// 	}
// }

// func BenchmarkBoltDBPutValue128B(b *testing.B) {
// 	b.ReportAllocs()
// 	b.ResetTimer()

// 	for n := 0; n < b.N; n++ {
// 		key := getKey(n)
// 		val := geyValue128B()
// 		if err = boltDB.Update(func(tx *bolt.Tx) error {
// 			bucket, err := tx.CreateBucketIfNotExists([]byte("bucket1"))
// 			if err != nil {
// 				return err
// 			}

// 			return bucket.Put([]byte(key), val)
// 		}); err != nil {
// 			b.Fatal(err)
// 		}
// 	}
// }

// func BenchmarkBoltDBPutValue256B(b *testing.B) {
// 	b.ReportAllocs()
// 	b.ResetTimer()

// 	for n := 0; n < b.N; n++ {
// 		key := getKey(n)
// 		val := geyValue256B()
// 		if err = boltDB.Update(func(tx *bolt.Tx) error {
// 			bucket, err := tx.CreateBucketIfNotExists([]byte("bucket1"))
// 			if err != nil {
// 				return err
// 			}

// 			return bucket.Put([]byte(key), val)
// 		}); err != nil {
// 			b.Fatal(err)
// 		}
// 	}
// }

// func BenchmarkBoltDBPutValue512B(b *testing.B) {
// 	b.ReportAllocs()
// 	b.ResetTimer()

// 	for n := 0; n < b.N; n++ {
// 		key := getKey(n)
// 		val := geyValue512B()
// 		if err = boltDB.Update(func(tx *bolt.Tx) error {
// 			bucket, err := tx.CreateBucketIfNotExists([]byte("bucket1"))
// 			if err != nil {
// 				return err
// 			}

// 			return bucket.Put([]byte(key), val)
// 		}); err != nil {
// 			b.Fatal(err)
// 		}
// 	}
// }

// func BenchmarkBoltDBGet(b *testing.B) {
// 	InitBoltDBData()

// 	b.ReportAllocs()
// 	b.ResetTimer()

// 	for n := 0; n < b.N; n++ {
// 		if err := boltDB.View(func(tx *bolt.Tx) error {
// 			key := []byte("key_" + fmt.Sprintf("%07d", 99))
// 			tx.Bucket([]byte("bucket1")).Get([]byte(key))
// 			return nil
// 		}); err != nil {
// 			b.Fatal(err)
// 		}
// 	}
// }

// func InitBoltDBData() {
// 	for n := 0; n < 10000; n++ {
// 		key := getKey(n)
// 		val := geyValue64B()

// 		if err = boltDB.Update(func(tx *bolt.Tx) error {
// 			bucket, err := tx.CreateBucketIfNotExists([]byte("bucket1"))
// 			if err != nil {
// 				return err
// 			}
// 			return bucket.Put([]byte(key), val)
// 		}); err != nil {
// 			panic(err)
// 		}
// 	}
// }
