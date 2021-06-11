package policy

type PackagePolicy struct {
	PackageName      string   `json: "PackageName"`
	Files      []FilePolicy   `json: "Files"`
}

type FilePolicy struct {
	FileName      string   `json: "FileName"`
}
