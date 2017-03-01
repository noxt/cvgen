package models

type Config struct {
	Template  TemplateRepo
	OutputDir string `yaml:"output_dir"`
}

type TemplateRepo struct {
	RepoURL string `yaml:"repo_url"`
	Path    string
	Files   []string
}

type UserInfo struct {
	AboutMe       Me
	Educations    []Education
	Organizations []Organization
	Projects      []Project
	Skills        []Skill
}

type Education struct {
	Name   string
	Start  string
	End    string
	Degree string
}

type Me struct {
	Name        string
	Description string
	Role        string
	Social      Social
	Contacts    Contacts
	Languages   []Language
}

type Social struct {
	Github   string
	Linkedin string
	Medium   string
	Twitter  string
}

type Contacts struct {
	Email   string
	Skype   string
	Phone   string
	Country string
	City    string
}

type Language struct {
	Name  string
	Level LanguageLevel
}

type LanguageLevel struct {
	Value       string
	Description string
}

type Organization struct {
	Name        string
	City        string
	Link        string
	Description string
	Role        string
	Start       string
	End         string
	Projects    []Project
}

type Project struct {
	Name         string
	Link         string
	Role         string
	Description  string
	Challenges   []string
	Technologies string
}

type Skill struct {
	Name     string
	Progress string
}
