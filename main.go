package kaleid
// package main
//
// import (
// 	"encoding/hex"
// 	"flag"
// 	"fmt"
// 	"io/ioutil"
//
// 	"zeptotenshi/kaleid/actions"
// )
//
// var (
// 	pw   *string
// 	path *string
// )
//
// func main() {
// 	pw = flag.String("pw", "", "passphrase to use to generate encryption key")
// 	path = flag.String("path", "", "path to data file to encrypt")
// 	flag.Parse()
//
// 	if *pw == "" {
// 		panic("empty passphrase")
// 	}
//
// 	if *path == "" {
// 		panic("empty path")
// 	}
//
// 	srcBytes, err := ioutil.ReadFile(*path)
// 	if err != nil {
// 		panic(err.Error())
// 	}
//
// 	fmt.Printf("passphrase: %s\n", *pw)
//
// 	bytes := actions.GenerateKey(*pw)
//
// 	key := hex.EncodeToString(bytes) //encode key in bytes to string and keep as secret, put in a vault
// 	fmt.Printf("key to encrypt/decrypt : %s\n", key)
//
// 	src := fmt.Sprintf("%s", srcBytes)
//
// 	encrypted, err := actions.Encrypt(src, key)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Printf("encrypted : %s\n", encrypted)
//
// 	decrypted, err := actions.Decrypt(encrypted, key)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Printf("decrypted : %s\n", decrypted)
// }
