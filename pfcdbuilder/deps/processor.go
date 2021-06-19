package deps

func NewProcessor() *Processor {
	return &Processor{
		package_redirects: map[string]string{},
		package_skiped:    map[string]bool{},
	}
}

type Processor struct {
	package_redirects map[string]string
	package_skiped    map[string]bool
}

func (p *Processor) AddRedirect(tag string, to string) {
	p.package_redirects[tag] = to
}

func (p *Processor) Redirects() map[string]string {
	return p.package_redirects
}

func (p *Processor) AddSkippedPackage(tag string) {
	p.package_skiped[tag] = true
}
