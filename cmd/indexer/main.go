package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/BurntSushi/toml"
)

const contentPath = "content"
const outputPath = "static/js/lunr/PagesIndex.json"

type FrontMatter struct {
	Title string   `toml:"title"`
	Tags  []string `toml:"tags"`
}

type PageIndex struct {
	Title   string   `json:"title"`
	Tags    []string `json:"tags,omitempty"`
	Href    string   `json:"href"`
	Content string   `json:"content"`
}

func main() {
	fmt.Println("Building pages index...")

	var pagesIndex []PageIndex

	err := filepath.Walk(contentPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Ignore directories and non-markdown files
		if info.IsDir() || !strings.HasSuffix(path, ".md") {
			return nil
		}

		// Ignore images and system files
		if matched, _ := regexp.MatchString(`\.(png|jpg|DS_Store)$`, path); matched {
			return nil
		}

		fmt.Printf("Processing: %s\n", path)
		pageIndex, err := processMarkdownFile(path)
		if err != nil {
			fmt.Printf("Warning: failed to process %s: %v\n", path, err)
			return nil
		}

		pagesIndex = append(pagesIndex, pageIndex)
		return nil
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error walking directory: %v\n", err)
		os.Exit(1)
	}

	// Write JSON output
	if err := writeJSON(outputPath, pagesIndex); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing output: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Index built successfully: %d pages indexed\n", len(pagesIndex))
}

func processMarkdownFile(path string) (PageIndex, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return PageIndex{}, err
	}

	// Split front matter and content
	parts := strings.Split(string(content), "+++")
	if len(parts) < 3 {
		return PageIndex{}, fmt.Errorf("invalid markdown format")
	}

	// Parse front matter
	var frontMatter FrontMatter
	if err := toml.Unmarshal([]byte(strings.TrimSpace(parts[1])), &frontMatter); err != nil {
		return PageIndex{}, fmt.Errorf("failed to parse front matter: %w", err)
	}

	// Generate href
	relPath := strings.TrimPrefix(path, contentPath)
	href := strings.TrimSuffix(relPath, ".md") + "/"
	
	// Special handling for index.md
	if strings.HasSuffix(path, "index.md") {
		href = strings.TrimSuffix(relPath, "index.md")
	}

	// Clean content: remove markdown syntax
	cleanContent := stripMarkdown(parts[2])

	return PageIndex{
		Title:   frontMatter.Title,
		Tags:    frontMatter.Tags,
		Href:    href,
		Content: cleanContent,
	}, nil
}

func stripMarkdown(content string) string {
	// Remove code blocks
	re := regexp.MustCompile("```[\\s\\S]*?```")
	content = re.ReplaceAllString(content, "")

	// Remove inline code
	content = strings.ReplaceAll(content, "`", "")

	// Remove links [text](url) -> text
	re = regexp.MustCompile(`\[([^\]]+)\]\([^\)]+\)`)
	content = re.ReplaceAllString(content, "$1")

	// Remove headers
	re = regexp.MustCompile(`#+\s*`)
	content = re.ReplaceAllString(content, "")

	// Remove emphasis
	content = strings.ReplaceAll(content, "**", "")
	content = strings.ReplaceAll(content, "*", "")
	content = strings.ReplaceAll(content, "__", "")
	content = strings.ReplaceAll(content, "_", "")

	// Remove extra whitespace
	re = regexp.MustCompile(`\s+`)
	content = re.ReplaceAllString(content, " ")

	return strings.TrimSpace(content)
}

func writeJSON(path string, data interface{}) error {
	// Create directory if it doesn't exist
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Marshal to JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// Write to file
	return os.WriteFile(path, jsonData, 0644)
}
