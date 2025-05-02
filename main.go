package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	docStyle     = lipgloss.NewStyle().Margin(1, 2).Border(lipgloss.NormalBorder())
	focusedStyle = lipgloss.NewStyle().Border(lipgloss.ThickBorder()).Padding()
	normalStyle  = lipgloss.NewStyle().Border(lipgloss.HiddenBorder()).Padding()
)

type card struct {
	title, desc string
}

func (c card) Title() string       { return c.title }
func (c card) Description() string { return c.desc }
func (c card) FilterValue() string { return c.title }

type board struct {
	lists            []list.Model
	focusIndex       int
	baseDelegate     list.ItemDelegate
	noCursorDelegate list.ItemDelegate
}

func newBoard() *board {
	// prepare our two delegates
	base := list.NewDefaultDelegate()
	noCursor := list.NewDefaultDelegate()
	// make "selected" look just like "normal"
	noCursor.Styles.SelectedTitle = noCursor.Styles.NormalTitle
	noCursor.Styles.SelectedDesc = noCursor.Styles.NormalDesc

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
		l := list.New(items, base, 0, 0)
		l.Title = t
		l.SetShowHelp(false)
		lists[i] = l
	}

	return &board{
		lists:            lists,
		baseDelegate:     base,
		noCursorDelegate: noCursor,
	}
}

func (b *board) Init() tea.Cmd {
	return nil
}

func (b *board) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return b, tea.Quit
		case "h":
			if b.focusIndex > 0 {
				b.focusIndex--
			}
		case "l":
			if b.focusIndex < len(b.lists)-1 {
				b.focusIndex++
			}
		}
	case tea.WindowSizeMsg:
		// size all lists equally
		h, v := docStyle.GetFrameSize()
		width := (msg.Width - h) / len(b.lists)
		height := msg.Height - v
		for i := range b.lists {
			b.lists[i].SetSize(width, height)
		}
	}

	// only update the focused list
	var cmd tea.Cmd
	b.lists[b.focusIndex], cmd = b.lists[b.focusIndex].Update(msg)
	return b, cmd
}

func (b *board) View() string {
	var views []string
	for i := range b.lists {
		m := &b.lists[i]
		if i == b.focusIndex {
			m.SetDelegate(b.baseDelegate)
			views = append(views, focusedStyle.Render(m.View()))
		} else {
			m.SetDelegate(b.noCursorDelegate)
			views = append(views, normalStyle.Render(m.View()))
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
