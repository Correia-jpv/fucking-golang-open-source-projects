package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type catalog struct {
	Title          string     `json:"title"`
	Description    string     `json:"description"`
	UpdatedAt      string     `json:"updated_at"`
	Highlights     []string   `json:"highlights"`
	SelectionRules []string   `json:"selection_rules"`
	Categories     []category `json:"categories"`
}

type category struct {
	Slug     string    `json:"slug"`
	Name     string    `json:"name"`
	Summary  string    `json:"summary"`
	Projects []project `json:"projects"`
}

type project struct {
	Name string `json:"name"`
	Repo string `json:"repo"`
	Desc string `json:"desc"`
}

func main() {
	raw, err := os.ReadFile("projects.json")
	if err != nil {
		fail(err)
	}

	var data catalog
	if err := json.Unmarshal(raw, &data); err != nil {
		fail(err)
	}

	var total int
	for _, category := range data.Categories {
		total += len(category.Projects)
	}

	var b strings.Builder
	b.WriteString("# ")
	b.WriteString(data.Title)
	b.WriteString("\n\n")
	b.WriteString(data.Description)
	b.WriteString("\n\n")
	b.WriteString(fmt.Sprintf("当前版本收录 **%d** 个项目，分成 **%d** 个主题；最近一次维护状态审阅时间为 **%s**。\n\n", total, len(data.Categories), data.UpdatedAt))
	b.WriteString("- [分类与维护策略](docs/分类与维护策略.md)\n")
	b.WriteString("- [移除与迁移记录](docs/移除与迁移记录.md)\n\n")

	b.WriteString("## 这次整理做了什么\n\n")
	for _, item := range data.Highlights {
		b.WriteString("- ")
		b.WriteString(item)
		b.WriteString("\n")
	}

	b.WriteString("\n## 收录原则\n\n")
	for _, rule := range data.SelectionRules {
		b.WriteString("- ")
		b.WriteString(rule)
		b.WriteString("\n")
	}

	b.WriteString("\n## 分类导航\n\n")
	b.WriteString("| 分类 | 关注点 | 项目数 |\n")
	b.WriteString("| --- | --- | --- |\n")
	for _, category := range data.Categories {
		b.WriteString("| ")
		b.WriteString(category.Name)
		b.WriteString(" | ")
		b.WriteString(category.Summary)
		b.WriteString(" | ")
		b.WriteString(fmt.Sprintf("%d", len(category.Projects)))
		b.WriteString(" |\n")
	}

	for _, category := range data.Categories {
		b.WriteString("\n## ")
		b.WriteString(category.Name)
		b.WriteString("\n\n")
		b.WriteString(category.Summary)
		b.WriteString("\n\n")
		b.WriteString("| 项目 | 简介 |\n")
		b.WriteString("| --- | --- |\n")
		for _, project := range category.Projects {
			b.WriteString("| [")
			b.WriteString(project.Repo)
			b.WriteString("](https://github.com/")
			b.WriteString(project.Repo)
			b.WriteString(") | ")
			b.WriteString(project.Desc)
			b.WriteString(" |\n")
		}
	}

	b.WriteString("\n## 维护说明\n\n")
	b.WriteString("目录已经去掉旧版 README 中的重复收录、过时仓库和“其它”大杂烩分类。后续如果继续扩展，建议优先更新 [projects.json](projects.json)，再运行 `go run ./tools/generate_readme.go` 同步生成 README。\n")

	if err := os.WriteFile("README.md", []byte(b.String()), 0o644); err != nil {
		fail(err)
	}
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
