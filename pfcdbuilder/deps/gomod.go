package deps

type Dependency struct {
	Import string
	Version string
}

type GoModHandler struct {
	Tag string
	Dependencies []Dependency
}
