package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func sumSha1(str string) string {
	t := sha1.New()
	io.WriteString(t, str)
	return fmt.Sprintf("%x", t.Sum(nil))
}

// now to sql DATETIME string
// refer to http://ichon.me/post/998.html
func nowToDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func makeJson(v interface{}) string {
	bin, err := json.Marshal(v)
	if err != nil {
		log.Print(err)
		return "Marshal json fail!"
	}
	return string(bin)
}

func isExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

// return [start, end)
func randInRange(start int, end int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn(end)
	return num + start
}

func itoa(i int) string {
	return strconv.Itoa(i)
}
