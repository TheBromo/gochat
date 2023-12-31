package terminalview

import (
	"fmt"
	"strings"
	"time"

	pb "github.com/TheBromo/gochat/common/chat"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TickMsg time.Time

type (
	errMsg error
)

type model struct {
	viewport    viewport.Model
	messages    []string
	textarea    textarea.Model
	senderStyle lipgloss.Style
	ready       bool
	err         error
	input       chan pb.Message
	output      chan pb.Message
	userName    string
}

// Send a message every second.
func tickEvery() tea.Cmd {
	return tea.Every(time.Millisecond*100, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func InitialModel(input chan pb.Message, output chan pb.Message, userName string) model {
	ta := textarea.New()
	ta.Placeholder = "Send a message..."
	ta.Focus()

	ta.Prompt = "┃ "
	ta.CharLimit = 280

	ta.SetWidth(50)
	ta.SetHeight(2)

	// Remove cursor line styling
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()
	ta.FocusedStyle.Base.Align(lipgloss.Bottom)

	ta.ShowLineNumbers = false

	vp := viewport.New(0, 0)
	ta.KeyMap.InsertNewline.SetEnabled(false)

	return model{
		textarea:    ta,
		messages:    []string{},
		viewport:    vp,
		senderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		err:         nil,
		input:       input,
		output:      output,
		userName:    userName,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(tickEvery(), tea.EnterAltScreen, textarea.Blink)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd tea.Cmd
		vpCmd tea.Cmd
	)

	m.textarea, tiCmd = m.textarea.Update(msg)
	m.viewport, vpCmd = m.viewport.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			m.input <- pb.Message{
				Sender:    m.userName,
				Timestamp: timestamppb.New(time.Now()),
				Content:   m.textarea.Value(),
			}
			m.textarea.SetValue("")
			m.textarea.Reset()
		}
	case TickMsg:
		select {
		case value, ok := <-m.output:
			if !ok {
				m.messages = append(m.messages, "error with output channel")
			}
			m.messages = append(m.messages, m.senderStyle.Render(value.Sender+": ")+value.Content)
			m.viewport.SetContent("\n\n\n" + strings.Join(m.messages, "\n"))
			m.viewport.GotoBottom()

			return m, tea.Batch(vpCmd, tickEvery())

		default:
			return m, tea.Batch(tickEvery())
		}

	case tea.WindowSizeMsg:
		m.handleResize(msg)
		m.textarea, tiCmd = m.textarea.Update(msg)
		m.viewport, vpCmd = m.viewport.Update(msg)

	case errMsg:
		m.err = msg
		return m, nil
	}

	return m, tea.Batch(tiCmd, vpCmd)
}

func (m *model) handleResize(msg tea.WindowSizeMsg) {
	chatHeight := lipgloss.Height(m.textarea.View())
	if !m.ready {
		m.viewport = viewport.New(msg.Width, msg.Height-chatHeight)
		m.textarea.SetWidth(msg.Width)
		m.ready = true
	} else {
		m.viewport.Width = msg.Width
		m.textarea.SetWidth(msg.Width)
		m.viewport.Height = msg.Height - chatHeight
	}
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s",
		m.viewport.View(),
		m.textarea.View(),
	) + "\n\n"
}
