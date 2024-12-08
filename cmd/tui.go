package cmd

import (
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type MainModel struct {
	textarea  textarea.Model
	textinput textinput.Model
	context   string
	messages  []string
	winHeight int
	winWidth  int
	viewport  viewport.Model
	timetaken float64
}

func InitMainModel() MainModel {
	ta := textarea.New()
	ta.CharLimit = 5000
	ta.Placeholder = "Enter your text here..."
	ta.Focus()
	ta.ShowLineNumbers = true

	ti := textinput.New()
	ti.CharLimit = 30
	ti.Placeholder = "Enter your question here..."

	vp := viewport.New(100, 20)
	vp.GotoBottom()

	return MainModel{
		textarea:  ta,
		textinput: ti,
		context:   "",
		winHeight: 0,
		winWidth:  0,
		viewport:  vp,
	}
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.winHeight = msg.Height
		m.winWidth = msg.Width
		m.textarea.SetWidth(m.winWidth)
		m.viewport.Width = m.winWidth
		m.viewport.Height = m.winHeight - 10
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyCtrlD.String():
			return m, tea.Quit
		case tea.KeyCtrlR.String():
			m.context = ""
			m.textinput.Reset()
		case tea.KeyDown.String():
			m.viewport.LineDown(1)
		case tea.KeyUp.String():
			m.viewport.LineUp(1)
		case tea.KeyEnter.String():
			start := time.Now()
			if m.context == "" {
				m.handleSummary()
			} else {
				m.handleQuestion()
			}
			end := time.Now()
			timediff := end.Sub(start).Seconds()
			m.timetaken = timediff
			m.viewport.SetContent(strings.Join(m.messages, "\n\n"))
			m.viewport.GotoBottom()
		default:
			var cmd tea.Cmd
			if m.context == "" {
				m.textarea, cmd = m.textarea.Update(msg)
			} else {
				m.textinput, cmd = m.textinput.Update(msg)
			}
			return m, cmd
		}
	}
	return m, nil
}

func (m *MainModel) handleSummary() {
	m.context = m.textarea.Value()
	outputSummary, _ := getSummary(m.context)
	m.messages = append(m.messages, SummaryStyle.Render(outputSummary))
	m.textarea.Reset()
	m.textinput.Focus()
}

func (m *MainModel) handleQuestion() {
	inputQuestion := m.textinput.Value()
	outputAnswer, _ := getAnswer(inputQuestion, m.context)
	m.messages = append(m.messages, "?: "+QuestionStyle.Render(inputQuestion))
	m.messages = append(m.messages, "A: "+AnswerStyle.Render(outputAnswer))
	m.textinput.Reset()
}

func (m MainModel) View() string {
	header := LogoStyle.Render(logo) + desc
	footer := strconv.FormatFloat(m.timetaken, 'f', 2, 64) + " sec"
	if m.context == "" {
		return header + "\n" + m.textarea.View() + "\n" + m.viewport.View()
	} else {
		return header + footer + "\n" + m.textinput.View() + "\n" + m.viewport.View()
	}
}
