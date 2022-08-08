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
	if envVal.Kind() != reflect.Pointer {
		log.Fatal("error you must pass a pointer to the env struct")
	}
	if envVal.IsNil() {
		log.Fatal("error env struct was empty")
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
		comment := strings.ContainsAny(line, "#")
		if !comment {
			xs := splitKeyValue(line)
			em[xs[0]] = xs[1]
		}
		if comment {
			skipLine := strings.HasPrefix(line, "#")
			if !skipLine {
				xs, err := splitKeyValueComment(line)
				if err == nil {
					em[xs[0]] = xs[1]
				}
			}
		}
	}
	return em
}

//splitKeyValue splits an env file line into a slice of strings
func splitKeyValue(line string) []string {
	return strings.SplitN(line, "=", 2)
}

//splitKeyValueComment takes an env file line and removes any inline comment and passes it to splitKeyValue
func splitKeyValueComment(line string) ([]string, error) {
	before, _, found := strings.Cut(line, "#")
	if !found {
		err := fmt.Errorf("error ")
		fmt.Println(err)
		return []string{}, err
	}
	trim := strings.Trim(before, " ")
	return splitKeyValue(trim), nil
}
