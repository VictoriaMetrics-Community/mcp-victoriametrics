// readmegen generates the generated sections of README.md (client setup, tools and prompts)
// and the client configs of the web UI. Client setup comes from web/src/clients.yaml,
// tools and prompts are taken from the server's own registration code.
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

const (
	readmeStartMarker = "<!-- clients:start -->"
	readmeEndMarker   = "<!-- clients:end -->"
	readmeServerURL   = "http://localhost:8080"
	urlPlaceholder    = "{{url}}"
)

type snippet struct {
	Text string `yaml:"text" json:"-"`
	Lang string `yaml:"lang" json:"lang"`
	Code string `yaml:"code" json:"code"`
}

type client struct {
	ID          string    `yaml:"id" json:"id"`
	Title       string    `yaml:"title" json:"title"`
	Description string    `yaml:"description" json:"description"`
	Badge       string    `yaml:"badge" json:"badge,omitempty"`
	Docs        string    `yaml:"docs" json:"docs,omitempty"`
	Local       []snippet `yaml:"local" json:"-"`
	Remote      snippet   `yaml:"remote" json:"remote"`
}

func main() {
	source := flag.String("source", "web/src/clients.yaml", "path to the clients source file")
	readme := flag.String("readme", "README.md", "path to README.md to update")
	uiConfig := flag.String("ui", "web/src/clients.gen.json", "path to the generated web UI configs")
	flag.Parse()

	clients, err := loadClients(*source)
	if err != nil {
		log.Fatalf("cannot load clients: %v", err)
	}
	var infos []*backendInfo
	for _, b := range backends {
		info, err := collectBackend(b)
		if err != nil {
			log.Fatalf("cannot list tools and prompts: %v", err)
		}
		infos = append(infos, info)
	}
	if err := updateReadme(*readme, clients, infos); err != nil {
		log.Fatalf("cannot update %s: %v", *readme, err)
	}
	if err := writeUIConfig(*uiConfig, clients); err != nil {
		log.Fatalf("cannot write %s: %v", *uiConfig, err)
	}
}

func loadClients(path string) ([]client, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var clients []client
	if err := yaml.UnmarshalStrict(data, &clients); err != nil {
		return nil, err
	}
	ids := make(map[string]bool, len(clients))
	for _, c := range clients {
		switch {
		case c.ID == "" || c.Title == "" || c.Description == "":
			return nil, fmt.Errorf("client %q: id, title and description are required", c.ID)
		case ids[c.ID]:
			return nil, fmt.Errorf("client %q is defined more than once", c.ID)
		case len(c.Local) == 0:
			return nil, fmt.Errorf("client %q: local setup is required", c.ID)
		case c.Remote.Code == "" || !strings.Contains(c.Remote.Code, urlPlaceholder):
			return nil, fmt.Errorf("client %q: remote setup must contain %s", c.ID, urlPlaceholder)
		}
		ids[c.ID] = true
	}
	return clients, nil
}

func renderReadme(clients []client) string {
	var b strings.Builder
	b.WriteString(readmeStartMarker + "\n")
	b.WriteString("<!-- Generated from web/src/clients.yaml by `make update-readme`, do not edit manually. -->\n")
	for _, c := range clients {
		fmt.Fprintf(&b, "\n### %s\n", c.Title)
		for _, s := range c.Local {
			writeSnippet(&b, s.Text, s.Lang, s.Code)
		}
		remoteText := fmt.Sprintf("To connect to a server running in [Streamable HTTP](#modes) mode at `%s`:", readmeServerURL)
		writeSnippet(&b, remoteText, c.Remote.Lang, strings.ReplaceAll(c.Remote.Code, urlPlaceholder, readmeServerURL))
		if c.Docs != "" {
			fmt.Fprintf(&b, "\nSee [%s MCP docs](%s) for more info.\n", c.Title, c.Docs)
		}
	}
	b.WriteString("\n" + readmeEndMarker)
	return b.String()
}

func writeSnippet(b *strings.Builder, text, lang, code string) {
	if text != "" {
		fmt.Fprintf(b, "\n%s\n", strings.TrimRight(text, "\n"))
	}
	fmt.Fprintf(b, "\n```%s\n%s\n```\n", lang, strings.TrimRight(code, "\n"))
}

func updateReadme(path string, clients []client, infos []*backendInfo) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	content := string(data)
	for _, section := range []struct{ start, end, text string }{
		{readmeStartMarker, readmeEndMarker, renderReadme(clients)},
		{toolsStartMarker, toolsEndMarker, renderTools(infos)},
		{promptsStartMarker, promptsEndMarker, renderPrompts(infos)},
	} {
		if content, err = replaceBetween(content, section.start, section.end, section.text); err != nil {
			return err
		}
	}
	return os.WriteFile(path, []byte(content), 0o644)
}

func writeUIConfig(path string, clients []client) error {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	if err := enc.Encode(clients); err != nil {
		return err
	}
	return os.WriteFile(path, buf.Bytes(), 0o644)
}
