package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gomarkdown/markdown"
	"golang.org/x/exp/slices"
)

var (
	MaleName   []string = ReadFromResource("../male-first-names.txt")
	FemaleName []string = ReadFromResource("../female-first-names.txt")
	FamilyName []string = ReadFromResource("../last-names.txt")
)

var MiddleChar []string = []string{".", "_"}

type Gender int

const (
	Male Gender = iota
	Female
	Both
)

func main() {
	var port string = ":8000"
	// Handlers
	http.HandleFunc("/", idx)
	http.HandleFunc("/get", GenerateRandomNameHandler)

	fmt.Println(fmt.Sprintf("Web server running on %v...", port))
	http.ListenAndServe(port, nil)
}

func idx(w http.ResponseWriter, req *http.Request) {
	dat, err := os.ReadFile("../README.md")
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  0,
			"message": "Oops, sorry. Something does not go as we expected.",
		})
		panic(err)
	}
	md := []byte(dat)
	doc := markdown.ToHTML(md, nil, nil)
	w.Write([]byte(doc))
}

func GenerateRandomNameHandler(w http.ResponseWriter, req *http.Request) {
	limitParam, ok := req.URL.Query()["limit"]

	if !ok || len(limitParam) < 1 {
		limitParam = append(limitParam, "1")
	}
	limit, err := strconv.Atoi(limitParam[0])
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  0,
			"message": "Oops, sorry. Something does not go as we expected.",
		})
		return
	}

	genderParam, ok := req.URL.Query()["gender"]
	if !ok || len(genderParam) < 1 {
		genderParam = append(genderParam, "both")
	}
	gender := genderParam[0]

	w.Header().Add("Content-Type", "application/json")
	res := map[string]any{
		"status": 1,
		"names":  GenerateRandomNames(limit, gender),
	}
	json.NewEncoder(w).Encode(&res)
}

func GenerateRandomNames(limit int, gender string) []string {
	var nameList []string
	switch GenderFromText(gender) {
	case Male:
		nameList = MaleName
	case Female:
		nameList = FemaleName
	case Both:
		nameList = append(MaleName, FemaleName...)
	}

	var result []string

	rand.Seed(time.Now().Unix())
	for len(result) < limit {
		newName := GenerateName(nameList)
		if slices.Contains(result, newName) {
			continue
		}
		result = append(result, strings.Trim(newName, "\r"))
	}
	return result
}

func GenerateName(nameList []string) string {
	var name string
	name += nameList[rand.Intn(len(nameList))]
	name += GetRandomChars()
	name += FamilyName[rand.Intn(len(FamilyName))]
	name += GenerateRandomNumber()
	return name
}

func GenderFromText(gender string) Gender {
	switch gender {
	case "male":
		return Male
	case "female":
		return Female
	}
	return Both
}

func GenerateRandomNumber() string {
	if !GenerateRandomBool() {
		return ""
	}
	return strconv.Itoa(rand.Intn(999))
}

func GetRandomChars() string {
	if !GenerateRandomBool() {
		return ""
	}
	return MiddleChar[rand.Intn(len(MiddleChar))]
}

func GenerateRandomBool() bool {
	return rand.Intn(2) == 1
}

func ReadFromResource(fileName string) []string {
	dat, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(dat), "\n")
}
