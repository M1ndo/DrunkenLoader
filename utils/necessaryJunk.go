package utils

import (
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"time"
)

// This file contains a list of actions are dummy actions. Why? Because by
// default Defender flags Go applications mainly because there's so much go malware
// that go apps are flagged through a runtime signature.
// Defender looks for legitimate actions in app to consider it safe.
// That's why this file exists, altho its probably not going to do much, why let it there.

type model struct {
	timer timer.Model
	quit  bool
	help  help.Model
}

func (m model) Init() tea.Cmd {
	return m.timer.Init()
}

func (m model) View() string {
	s := m.timer.View() + "\n"
	if m.timer.Timedout() {
		s = "All done!"
	}
	if !m.quit {
		s = "Evading starts In: " + s
	}
	return s
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case timer.TickMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd

	case timer.StartStopMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd

	case timer.TimeoutMsg:
		m.quit = true
		return m, tea.Quit
	}

	return m, nil
}

func RunApp() {
	m := model{
		timer: timer.NewWithInterval(time.Second*2, time.Millisecond),
		help:  help.New(),
	}
	newProgram := tea.NewProgram(m)
	//timer := time.AfterFunc(10*time.Second, func() {
	//	newProgram.Quit()
	//})

	if _, err := newProgram.Run(); err != nil {
		fmt.Println("Oh no, it didn't work:", err)
		os.Exit(1)
	}
	//timer.Stop()
}
