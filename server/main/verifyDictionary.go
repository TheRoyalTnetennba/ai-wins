package main

import (
    "fmt"
    "sort"
    "os"
    "bufio"
    "strings"
    "cloud.google.com/go/datastore"
)

func (x *DictWord) Sorted () string {
    return x.Key.Name
}

func readDict(path string) ([]string) {
    file, err := os.Open(path)
    if err != nil {
        fmt.Println("error reading file")
        return nil
    }
    defer file.Close()
    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines
}

func getLocalDict() (map[string][]string) {
    dict := make(map[string][]string)
    words := readDict("./server/main/assets/sowpods.txt")
    for _, word := range words {
        arr := strings.Split(word, "")
        sort.Strings(arr)
        sorted := strings.Join(arr, "")
        if val, ok := dict[sorted]; ok {
            newVal := append(val, word)
            dict[sorted] = newVal
        } else {
            newVal := []string{word}
            dict[sorted] = newVal
        }
    }
    return dict
}

func VerifyDictionary() {
    dict := getLocalDict()

    var words []*DictWord

    _, err := Client.GetAll(Ctx, datastore.NewQuery("DictWord"), &words)
    if err != nil {
      fmt.Println("Error fetching data")
    }

    for i := 0; i < len(words); i++ {
        // fmt.Println(words[i])
        delete(dict, words[i].Key.Name)
    }

    for k, v := range dict {
        newKey := datastore.NameKey("DictWord", k, nil)
        newEntity := &DictWord{Anagrams: v}
        newKey, err = Client.Put(Ctx, newKey, newEntity)
        if err != nil {
            fmt.Println("Could not save")
            fmt.Println(newEntity)
            fmt.Println(err)
        }
    }
}
