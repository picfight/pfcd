package policy

const YES = "YES"
const NO = "NO"

type PackagePolicy struct {
	PackageName string       `json: "PackageName"`
	Files       []FilePolicy `json: "Files"`
}

type FilePolicy struct {
	FileName string `json: "FileName"`
	CopyAsIs string `json: "CopyAsIs"`
}

func (p *FilePolicy) IsCopyAsIs() bool {
	return p.CopyAsIs == "YES"
}

func (p *FilePolicy) IsValid() bool {
	if p.FileName == "" {
		return false
	}
	if p.CopyAsIs == "" {
		return false
	}
	return true
}
