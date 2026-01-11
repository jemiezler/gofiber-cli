package {{.Name}}

type {{.Pascal}}Service struct {
	repo *{{.Pascal}}Repository
}

func New{{.Pascal}}Service(repo *{{.Pascal}}Repository) *{{.Pascal}}Service {
	return &{{.Pascal}}Service{repo: repo}
}
