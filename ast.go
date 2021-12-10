package main

type File struct {
	QualifiedPackageName string
	Imports              []string
	Class                *Class
}

type Class struct {
	Name    string
	Members []*Member
}

type Member struct {
	Name      string
	Static    bool
	Output    Type
	Arguments []*Argument
	Body      []*CodeLine
}

type Type struct {
}

type Argument struct {
}

type CodeLine struct {
}
