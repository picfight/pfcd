package policy

const YES = "YES"
const NO = "NO"

type PackagePolicy struct {
	PackageName       string `json: "PackageName"`
	SkipProcessing    string `json: "SkipProcessing"`
	RedirectPackageTo string `json: "RedirectPackageTo"`

	ConvertFiles string `json: "ConvertFiles"`
	UseInjectors string `json: "UseInjectors"`

	Files []FilePolicy `json: "Files"`
}

func (p *PackagePolicy) IsRedirectPackage() bool {
	return p.RedirectPackageTo != ""
}

func (p *PackagePolicy) IsSkipProcessing() bool {
	return p.SkipProcessing == "YES"
}

func (p *PackagePolicy) IsConvertFiles() bool {
	return p.ConvertFiles == "YES"
}

type FilePolicy struct {
	FileName        string `json: "FileName"`
	CopyAsIs        string `json: "CopyAsIs"`
	UseAutoReplacer string `json: "UseAutoReplacer"`
}

func (p *FilePolicy) IsCopyAsIs() bool {
	return p.CopyAsIs == YES
}

func (p *FilePolicy) IsValid() bool {
	if p.FileName == "" {
		return false
	}
	if p.CopyAsIs == "" {
		return false
	}
	if p.CopyAsIs == NO {
		if p.UseAutoReplacer == "YES" {
			return true
		}
		return false
	}
	return true
}

func (p *FilePolicy) IsUseAutoReplacer() bool {
	return p.UseAutoReplacer == YES
}
