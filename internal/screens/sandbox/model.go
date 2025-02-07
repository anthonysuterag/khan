package sandbox

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/evertras/khan/internal/components/datatree"
	"github.com/hashicorp/nomad/api"
)

type SampleStruct struct {
	Inner struct {
		ID         string
		Another    int
		unexported float64
	}

	Name string
	shh  int

	Job api.Job

	Nums []int

	MyMap map[string]string
}

type Model struct {
	tree datatree.Model
}

func NewModel() Model {
	sample := SampleStruct{
		Name: "Hello",
		Nums: []int{3, 4, 1},
		MyMap: map[string]string{
			"bb":       "b",
			"hi":       "ok",
			"aardvark": "highest",
		},
	}

	sample.Inner.Another = 3

	sample.Inner.ID = "some-id"

	return Model{
		tree: datatree.New(&sample),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd tea.Cmd
	)

	m.tree, cmd = m.tree.Update(msg)

	return m, cmd
}

func (m Model) View() string {
	body := strings.Builder{}

	body.WriteString(m.tree.View())

	return body.String()
}
