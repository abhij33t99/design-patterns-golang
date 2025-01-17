package creational

import "fmt"

type INode interface {
	print(string)
	clone() INode
}

type File struct {
	name string
}

func (f *File) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *File) clone() INode {
	return &File{
		name: f.name + "_clone",
	}
}

type Folder struct {
	children []INode
	name     string
}

func (f *Folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.children {
		i.print(indentation + indentation)
	}
}

func (f *Folder) clone() INode {
	cloneFolder := &Folder{
		name:     f.name + "_clone",
		children: make([]INode, len(f.children)),
	}
	for _, i := range f.children {
		copy := i.clone()
		cloneFolder.children = append(cloneFolder.children, copy)
	}
	return cloneFolder
}
