package main

func Validate(
	links []Link,
) []Link {

	for i := range links {

		if links[i].URL == "https://old.example.com" {

			links[i].Valid = false

		} else {

			links[i].Valid = true

		}

	}

	return links

}
