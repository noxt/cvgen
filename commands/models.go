package commands

type userInfo struct {
	AboutMe       me
	Educations    []education
	Organizations []organization
	Projects      []project
	Skills        []skill
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
	Languages   []language
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

type language struct {
	Name  string
	Level languageLevel
}

type languageLevel struct {
	Value       string
	Description string
}

type organization struct {
	Name        string
	City        string
	Link        string
	Description string
	Role        string
	Start       string
	End         string
	Projects    []project
}

type project struct {
	Name         string
	Link         string
	Role         string
	Description  string
	Challenges   []string
	Technologies string
}

type skill struct {
	Name     string
	Progress string
}
