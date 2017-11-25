package glbs

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

const (
	DEFAULT_NAMESPACE 	= "default"
)

var space = DEFAULT_NAMESPACE

func basepath(key string) (string, error) {

	if len(key) == 0 {
		return "", errors.New("glbs basepath(): 0 length key")
	}

	return fmt.Sprintf("%s/%s", space, key[0:2]), nil

} // basepath

func path(key string) (string, error) {

	if len(key) == 0 {
		return "", errors.New("glbs path(): 0 length key")
	}
	
	return fmt.Sprintf("%s/%s/%s", space, key[0:2], key), nil

} // path

func hash(data []byte) string {

  digest := sha256.New()

	digest.Write(data)

	return hex.EncodeToString(digest.Sum(nil))

} // hash

func SetNamespace(namespace string) {
	space = namespace
} // SetNamespace

func Put(file io.Reader) *string {
  
	data, err := ioutil.ReadAll(file)

	if err != nil {
		log.Printf("glbs Put(): %s", err)
		return nil
	}

  key := hash(data)

	if len(key) == 0 {
		return nil
	}

	if Exists(key) {
		log.Println("glbs Put(): blob exists")
		return &key
	} else {

		bp, err := basepath(key)

		if err != nil {
			log.Printf("glbs Put(): %s", err)
			return nil
		}

		errMk := os.MkdirAll(bp, 0755)

		if errMk != nil {
			log.Printf("glbs Put(): %s", errMk)
			return nil
		}

		p, err := path(key)

		blob, errCreate := os.Create(p)

		defer blob.Close()

		if errCreate != nil {
			log.Printf("glbs Put(): %s", errCreate)
			return nil
		}

		_, errCopy := blob.Write(data)

		if errCopy != nil {
			log.Printf("glbs Put(): %s", errCopy)
			return nil
		}

		return &key

	}

} // Put

func Get(key string) ([]byte, error) {

  p, err := path(key)

	if err != nil {
		log.Println(err)
		return nil, errors.New("glbs Get(): no file found")
	}

  file, err := os.Open(p)

  defer file.Close()

	if err != nil {
		log.Println(err)
		return nil, errors.New("glbs Get(): unable to open file")
	}

	buf, errRead := ioutil.ReadAll(file)

	if errRead != nil {
		log.Println(errRead)
		return nil, errors.New("glbs Get(): unable to read file")
	}
	
	return buf, nil

} // Get

func GetPath(key string) *string {

  p, err := path(key)

	if err != nil {
		log.Println(err)
		return nil
	}

	return &p

} // GetPath

func Delete(key string) {

} // Delete

func Exists(key string) bool {

	p, err := path(key)

	if err != nil {
		log.Println(err)
		return false
	}

	_, errStat := os.Stat(p)

  if os.IsNotExist(errStat) {
		return false
	} else {
		return true
	}

} // Exists
