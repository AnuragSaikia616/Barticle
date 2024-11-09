package main

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	width    int
	height   int
	messages []string
	textarea textarea.Model
	viewport viewport.Model
	context  string
	numchars int
	timediff float64
}

func InitModel() model {
	ta := textarea.New()
	ta.CharLimit = 0
	ta.Placeholder = "get summary!!!"
	ta.SetHeight(10)
	ta.SetWidth(100)
	ta.Focus()
	vp := viewport.New(100, 10)
	return model{
		textarea: ta,
		viewport: vp,
		context:  "",
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.viewport.Height = msg.Height - 20
		m.viewport.Width = msg.Width
		m.textarea.SetWidth(msg.Width)
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+d":
			return m, tea.Quit
		case "ctrl+r":
			m.textarea.Placeholder = "get summary!!!"
			m.context = ""
			return m, nil
		case tea.KeyEnter.String():
			v := m.textarea.Value()
			if v == "" {
				return m, nil
			}
			stime := time.Now()
			summ, err := getSummary(v)
			etime := time.Now()
			m.timediff = stime.Sub(etime).Seconds()
			if err != nil {
				m.viewport.SetContent(errorStyle.Render(err.Error()))
			}
			m.viewport.SetContent(summ)
			m.context = v
			m.textarea.Reset()
			m.viewport.GotoBottom()
			return m, nil
		default:
			var cmd tea.Cmd
			m.numchars = m.textarea.Length()
			m.textarea, cmd = m.textarea.Update(msg)
			return m, cmd
		}
	}
	return m, nil
}

func (m model) View() string {
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
