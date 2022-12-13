package utils

type Directory struct {
	Name           string
	Size           int
	ChildFile      []File
	ChildDirectory map[string]*Directory
	Parent         *Directory
}

type File struct {
	Name string
	Size int
}

func MakeDir(name string, parent *Directory) *Directory {
	return &Directory{
		Name:           name,
		Parent:         parent,
		ChildDirectory: map[string]*Directory{},
	}
}
