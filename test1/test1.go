package test1

import (
	"fmt"

	ti "github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type textinput struct {
	textInput ti.Model
	err       error
	done      bool
	prompt    string
}

func newTextinput(prompt, placeholder, value string) textinput {
	ti := ti.New()
	ti.SetValue(value)
	ti.Placeholder = placeholder
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return textinput{
		textInput: ti,
		err:       nil,
		prompt:    prompt,
	}
}

func (m textinput) Init() tea.Cmd {
	return ti.Blink
}

func (m textinput) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.err = fmt.Errorf("cancelled by user")
			fallthrough
		case "enter", "esc":
			m.done = true
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case error:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m textinput) View() string {
	if m.done {
		return ""
	}

	return fmt.Sprintf(
		"%s\n\n%s\n\n%s\n",
		m.prompt,
		m.textInput.View(),
		"(press <enter> to submit)",
	)
}

func TextInput(prompt, placeholder, value string) (string, error) {
	ti := newTextinput(prompt, placeholder, value)
	p := tea.NewProgram(ti)
	m, err := p.Run()
	if err != nil {
		return "", err
	}

	model, ok := m.(textinput)
	if !ok {
		return "", fmt.Errorf("unexpected model type")
	}

	return model.textInput.Value(), model.err
}
