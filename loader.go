package envloader

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

const (
	tagName = "env"
)

//ParseEnv takes a pointer to a struct and an env file name and populates that struct with the values
func ParseEnv(e interface{}, filename string) {
	envMap := createEnvMap(filename)
	envVal := reflect.ValueOf(e)
	if envVal.Kind() != reflect.Pointer || envVal.IsNil() {
		log.Fatal("env struct is not a pointer or has nil value")
	}
	envEl := envVal.Elem()
	envType := reflect.TypeOf(e).Elem()
	for i := 0; i < envType.NumField(); i++ {
		t := envType.Field(i)
		v := envEl.Field(i)
		tag := strings.Split(t.Tag.Get(tagName), ",")
		tagId := tag[0]
		required := strings.Split(tag[1], "=")[1]
		envVar := envMap[tagId]
		if envVar == "" && required == "true" {
			log.Fatalf("env variable %v is required got empty string", tagId)
		} else if envVar == "" && required == "false" {
			fmt.Printf("%v is empty var in map and is not required '%v' not set\n", tagId, t.Name)
		} else {
			v.SetString(envVar)
		}
	}
}

func createEnvMap(f string) map[string]string {
	em := make(map[string]string)
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(f, " file not found")
	}
	defer file.Close()
	s := bufio.NewScanner(file)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		line := s.Text()
		if string(line[0]) != "#" {
			e := strings.Split(line, "=")
			em[e[0]] = e[1]
		}
	}
	return em
}
