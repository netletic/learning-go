package kv

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
)

type store struct {
	path string
	data map[string]string
}

func OpenStore(path string) (*store, error) {
	s := &store{
		path: path,
		data: map[string]string{},
	}
	f, err := os.Open(path)
	if errors.Is(err
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(s.data)
}

func (s *store) All() map[string]string {
	return s.data
}

const Usage = `Usage: kv COMMAND [key] [value]

kv is a tool to manage a simple key-value store of strings. It
understands the following commands:

	kv all
		Lists all key-value pairs in the store file

	kv get KEY
		Prints the value associated with KEY, if one exists

	kv set KEY VALUE
		Sets KEY to be associated with VALUE, overwriting
		any previous associated value.

	The default store file is 'kv.store'. This file will be
	created automatically the first time a value is set using
	'kv set'.`

func Main() int {
	if len(os.Args) < 2 {
		fmt.Println(Usage)
		return 0
	}
	command := os.Args[1]
	switch command {
	case "all":
		return MainAll()
	case "get":
		return MainGet()
	case "set":
		return MainSet()
	}
	fmt.Fprint(os.Stderr, Usage)
	return 1
}

func MainAll() int {
	s, err := OpenStore("kv.store")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	for k, v := range s.All() {
		fmt.Printf("%s=%s\n", k, v)
	}
	// solution below is if key=value should be printed out lexicograhically sorted by key
	// data := s.All()
	// keys := make([]string, 0, len(data))
	// for k := range data {
	// 	keys = append(keys, k)
	// }
	// sort.Strings(keys)
	// for _, k := range keys {
	// 	v := data[k]
	// 	fmt.Printf("%s=%s\n", k, v)
	// }
	return 0
}

func MainGet() int {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, Usage)
		return 1
	}
	s, err := OpenStore("kv.store")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	value, ok := s.Get(os.Args[2])
	if !ok {
		fmt.Fprintf(os.Stderr, "key %q not found\n", os.Args[2])
		return 1
	}
	fmt.Println(value)
	return 0
}

func MainSet() int {
	if len(os.Args) < 4 {
		fmt.Fprintln(os.Stderr, Usage)
		return 1
	}
	s, err := OpenStore("kv.store")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	s.Set(os.Args[2], os.Args[3])
	err = s.Save()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	return 0
}
