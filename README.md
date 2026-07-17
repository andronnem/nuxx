# Markdown Link Checker

A lightweight utility that scans Markdown documentation and validates hyperlinks.

The project searches README files, extracts URLs, checks their status and generates a summary report.

---

## Example

```
Scanning project...

README.md

✔ https://golang.org

✔ https://github.com

✖ https://old.example.com

--------------------------

Markdown files : 3

Links checked : 7

Broken links : 1
```

---

## Project Structure

- scanner.go — finds Markdown files
- markdown.go — represents documents
- extractor.go — extracts URLs
- validator.go — validates links
- report.go — prints results
- sample_docs.go — demo project
- link.go — link model

---

Run

```bash
go run .
```
