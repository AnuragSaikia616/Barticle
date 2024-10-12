package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	gowiki "github.com/trietmn/go-wiki"
)

type model struct {
	width    int
	height   int
	messages []string
	textarea textinput.Model
	viewport viewport.Model
	summary  string
}

func InitModel() model {
	ta := textinput.New()
	ta.Placeholder = "get summary!!!"
	ta.CharLimit = 500
	ta.Focus()
	vp := viewport.New(100, 20)
	return model{
		textarea: ta,
		viewport: vp,
		summary:  "",
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+d":
			return m, tea.Quit
		case "ctrl+r":
			m.textarea.Placeholder = "get summary!!!"
			m.summary = ""
			return m, nil
		case tea.KeyEnter.String():
			v := m.textarea.Value()
			if v == "" {
				return m, nil
			}
			if m.summary == "" {
				summ, _ := gowiki.Summary(v, 5, -1, false, true)
				m.messages = append(m.messages, serverStyle.Render("SUMMARY:"), summ)
				m.summary = summ
				m.textarea.Placeholder = "ask questions!"
			} else {
				m.messages = append(m.messages, questionStyle.Render("QUESTION:"), v)
				m.messages = append(m.messages, answerStyle.Render("ANSWER:"), "answer to the question!!!")
			}
			m.viewport.SetContent(strings.Join(m.messages, "\n"))
			m.textarea.Reset()
			m.viewport.GotoBottom()
			return m, nil
		default:
			var cmd tea.Cmd
			m.textarea, cmd = m.textarea.Update(msg)
			return m, cmd
		}
	}
	return m, nil
}

func (m model) View() string {
	// return lipgloss.JoinVertical(lipgloss.Center, m.textarea.View(), m.viewport.View())
	return fmt.Sprintf("\n%s\n%s\n", m.textarea.View(), m.viewport.View())
}

func (m model) Init() tea.Cmd {
	return nil
}

func main() {
	m := InitModel()
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
