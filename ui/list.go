package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github/mirislomovmirjalol/DotEM/internal/service"
	"log"
	"strings"
)

type Model struct {
	cursor     int
	projects   []string
	selected   int
	showDetail bool
	detail     string
}

func NewModel(projects []string) Model {
	return Model{
		cursor:     0,
		projects:   projects,
		selected:   0,
		showDetail: false,
		detail:     "",
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if !m.showDetail {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.Type {
			case tea.KeyUp:
				if m.cursor > 0 {
					m.cursor--
				}
			case tea.KeyDown:
				if m.cursor < len(m.projects)-1 {
					m.cursor++
				}
			case tea.KeyEnter:
				m.showDetail = true
				project, err := service.GetProject(m.projects[m.cursor])
				if err != nil {
					log.Fatal(err)
				}
				m.detail = string(project)
				return m, tea.Batch(tea.ClearScreen, func() tea.Msg {
					return nil
				})
			case tea.KeyCtrlC:
				return m, tea.Quit
			default:
			}
		}
	} else {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.Type {
			case tea.KeyEsc:
				m.showDetail = false
			case tea.KeyCtrlC:
				return m, tea.Quit
			default:
			}
		}
	}
	return m, nil
}

var unfocused = lipgloss.NewStyle()
var focused = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#a78bfa"))
var navigationHelper = "" +
	"Up / Down: Navigate | Enter: Select | Ctrl+C: Quit"

func (m Model) View() string {
	var b strings.Builder

	if m.showDetail {
		b.WriteString(m.detail)
		return b.String()
	}
	b.WriteString(lipgloss.NewStyle().Padding(1, 8).Background(lipgloss.Color("#7e22ce")).Render("Projects"))
	b.WriteString("\n\n")

	for i, project := range m.projects {
		cursor := "  "
		if i == m.cursor {
			cursor = "| "
			b.WriteString(focused.Render(cursor + project))
		} else {
			b.WriteString(unfocused.Render(cursor + project))
		}
		b.WriteString("\n\n")
	}

	b.WriteString(navigationHelper)

	return b.String()
}
