package main

func SampleProject() []Markdown {

	return []Markdown{

		{
			Name: "README.md",

			Links: []Link{

				{URL: "https://golang.org"},

				{URL: "https://github.com"},

				{URL: "https://old.example.com"},
			},
		},

		{
			Name: "docs/api.md",

			Links: []Link{

				{URL: "https://go.dev"},
			},
		},

		{
			Name: "docs/install.md",

			Links: []Link{

				{URL: "https://pkg.go.dev"},
			},
		},
	}

}
