package main

import "fmt"

func PrintReport(
	files []Markdown,
) {

	total := 0

	broken := 0

	fmt.Println("Scanning project...\n")

	for _, file := range files {

		fmt.Println(file.Name)

		fmt.Println()

		checked := Validate(

			ExtractLinks(file),

		)

		for _, link := range checked {

			if link.Valid {

				fmt.Println("✔", link.URL)

			} else {

				fmt.Println("✖", link.URL)

				broken++

			}

			total++

		}

		fmt.Println()

	}

	fmt.Println("--------------------------")

	fmt.Println()

	fmt.Println("Markdown files :", len(files))

	fmt.Println("Links checked :", total)

	fmt.Println("Broken links :", broken)

}
