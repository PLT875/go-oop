package person

import "strings"

type Person struct {
	name string
	job  string
}

func (p *Person) SetPersonName(name string) {
	p.name = name
}

func (p *Person) SetPersonJob(job string) {
	p.job = job
}

func (p *Person) NewPerson(name string, job string) {
	p.name = name
	p.job = job
}

func (p *Person) GetName() string {
	return p.name
}

func (p *Person) ToString() string {
	p_attr := []string{"Name:", p.name, "Job:", p.job}
	return strings.Join(p_attr, " ")
}
