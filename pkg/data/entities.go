package data

type Translation struct {
	Texts []Text
}

type Text struct {
	Language string
	Text     string
}
