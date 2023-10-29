package storage

import (
	"os"
	"os/user"
	"path"
	"sort"

	"github.com/Grimmore/Sentient3/modelos"
)

func Explore(p string) ([]*modelos.Element, error) {
	if p == "" {
		return GetStorages(), nil
	}
	elements := []*modelos.Element{}
	res, e := os.ReadDir(p)
	if e != nil {
		return nil, e
	}

	var ex bool
	var wat bool
	var file *modelos.Element
	var er error

	for _, r := range res {
		ex = false
		wat = false
		file, er = modelos.GetElementByPath(path.Join(p, r.Name()))
		if er == nil {
			ex = file.Excluded
			wat = file.Watching
		}
		elements = append(elements, &modelos.Element{
			IsDir:    r.IsDir(),
			Name:     r.Name(),
			Path:     path.Join(p, r.Name()),
			Excluded: ex,
			Watching: wat,
		})
	}
	sort.Slice(elements, func(i, j int) bool {
		b := elements[i].Name < elements[j].Name
		if elements[i].IsDir && !elements[j].IsDir {
			b = true
		} else if elements[j].IsDir && !elements[i].IsDir {
			b = false
		}
		return b
	})
	return elements, nil
}

func GetStorages() []*modelos.Element {
	elements := []*modelos.Element{}
	places := getStorages()
	for _, r := range places {
		_, n := path.Split(r)
		elements = append(elements, &modelos.Element{
			IsDir: true,
			Name:  n,
			Path:  r,
		})
	}
	return elements
}

func getUsername() string {
	user, err := user.Current()
	if err == nil {
		return user.Username
	}
	return ""
}

func checkStorage(p string) bool {
	//_, e := os.ReadDir(p)
	//return e == nil
	inf, err := os.Stat(p)
	return err == nil && inf.IsDir()
}
