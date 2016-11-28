package commands

type userInfo struct {
	AboutMe       me
	Educations    []education
	Organizations []organization
	Projects      []project
}

type education struct {
	Name   string
	Start  string
	End    string
	Degree string
}

type me struct {
	Name        string
	Description string
	Role        string
	Social      social
	Contacts    contacts
}

type social struct {
	Github   string
	Linkedin string
	Medium   string
	Twitter  string
}

type contacts struct {
	Email   string
	Skype   string
	Phone   string
	Country string
	City    string
}

type organization struct {
	Id    string
	Name  string
	Link  string
	Role  string
	Start string
	End   string
}

type project struct {
	Name           string
	Link           string
	Role           string
	Description    string
	OrganizationId string `yaml:"organization_id"`
	Categories     []string
	Challenges     []string
	Technologies   []string
}
