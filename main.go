package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type card struct {
	title, desc string
}

func (c card) Title() string       { return c.title }
func (c card) Description() string { return c.desc }
func (c card) FilterValue() string { return c.title }

type data struct {
	lists []list.Model
}

type state struct {
	listIdx int
}

type board struct {
	Data  data
	State state
}

func newBoard() board {
	// sample items
	items := []list.Item{
		card{"Raspberry Pi’s", "I have ’em all over my house"},
		card{"Nutella", "It's good on toast"},
		card{"Bitter melon", "It cools you down"},
		card{"Nice socks", "And by that I mean socks without holes"},
		card{"Eight hours of sleep", "I had this once"},
		card{"Cats", "Usually"},
		card{"Plantasia, the album", "My plants love it too"},
		card{"Pour over coffee", "It takes forever to make though"},
		card{"VR", "Virtual reality...what is there to say?"},
		card{"Noguchi Lamps", "Such pleasing organic forms"},
		card{"Linux", "Pretty much the best OS"},
	}

	// build three lists
	titles := []string{"To Do", "Doing", "Done"}
	lists := make([]list.Model, len(titles))
	for i, t := range titles {
		l := list.New(items, BaseDelegate, 0, 0)
		l.Title = t
		l.SetShowHelp(false)
		lists[i] = l
	}

	return board{
		Data: data{
			lists: lists,
		},
		State: state{
			listIdx: 0,
		},
	}
}

func (b board) Init() tea.Cmd {
	return nil
}

func (b board) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return b, tea.Quit
		case "h":
			if b.State.listIdx > 0 {
				b.State.listIdx--
			}
		case "l":
			if b.State.listIdx < len(b.Data.lists)-1 {
				b.State.listIdx++
			}
		}
	case tea.WindowSizeMsg:
		// size all lists equally
		h, v := DocStyle.GetFrameSize()
		width := (msg.Width - h) / len(b.Data.lists)
		height := msg.Height - v
		for i := range b.Data.lists {
			b.Data.lists[i].SetSize(width, height)
		}
	}

	// only update the focused list
	var cmd tea.Cmd
	b.Data.lists[b.State.listIdx], cmd = b.Data.lists[b.State.listIdx].Update(msg)
	return b, cmd
}

func (b board) View() string {
	var views []string
	for i := range b.Data.lists {
		m := b.Data.lists[i]
		if i == b.State.listIdx {
			m.SetDelegate(BaseDelegate)
			views = append(views, FocusedStyle.Render(m.View()))
		} else {
			m.SetDelegate(NoCursorDelegate)
			views = append(views, NormalStyle.Render(m.View()))
		}
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, views...)
}

func main() {
	p := tea.NewProgram(newBoard(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v\n", err)
		os.Exit(1)
	}
}
