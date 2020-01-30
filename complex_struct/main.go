package main

import (
	"html/template"
	"log"
	"os"
)

func main() {

	const tpl = `
	{{.Name}}{{ range $child := .Children }} {{ $child.Name }} {{ range $grandchild := $child.Children }} {{ $grandchild.Name }} {{ end }} {{ end }}
`

	root := mine{
		"Steven", []mine{
			{
				"Stanley", []mine{
					{
						"Gozie",
						 nil,
					},
					{
						"Mark",
						nil,
					},
				},
			},
		},
	}
	t := template.Must(template.New("tree").Parse(tpl))
	err := t.Execute(os.Stdout, root)
	if err != nil {
		log.Fatalf("executing template: ", err)
	}

	//fmt.Println(ans)

}

type mine struct {
	Name  string
	Children []mine
}
