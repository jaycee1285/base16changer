package ui

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/jaycee1285/base16changer/internal/scheme"
	"github.com/jaycee1285/base16changer/internal/targets"
)

// Layout constraints
const (
	maxWidth  = 70
	maxHeight = 30
)

type tab int

const (
	tabSchemes tab = iota
	tabIcons
	tabWalls
	tabCount
)

var tabNames = []string{"Schemes", "Icons", "Wallpapers"}

type item struct{ title string }

func (i item) Title() string       { return i.title }
func (i item) Description() string { return "" }
func (i item) FilterValue() string { return i.title }

// Compact delegate for items inside expanded panels
type compactDelegate struct {
	normal  lipgloss.Style
	focused lipgloss.Style
}

func newCompactDelegate() compactDelegate {
	return compactDelegate{
		normal:  lipgloss.NewStyle(),
		focused: lipgloss.NewStyle().Bold(true).Reverse(true),
	}
}

func (d compactDelegate) Height() int                               { return 1 }
func (d compactDelegate) Spacing() int                              { return 0 }
func (d compactDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }

func (d compactDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	it, _ := listItem.(item)
	if index == m.Index() {
		line := "  ▸ " + it.title
		fmt.Fprint(w, d.focused.Render(line))
		return
	}
	line := "    " + it.title
	fmt.Fprint(w, d.normal.Render(line))
}

type dataLoadedMsg struct {
	schemes []string
	icons   []string
	walls   []string
}

type applyDoneMsg struct{ err error }

// Selections tracks what the user has chosen
type Selections struct {
	Scheme    string
	IconTheme string
	Wallpaper string
}

type Model struct {
	active   tab // Currently focused panel (title row)
	expanded tab // Which panel is expanded (-1 = none)
	inList   bool
	lists    map[tab]list.Model
	spinner  spinner.Model
	width    int
	height   int

	schemes []string
	icons   []string
	walls   []string

	cfg      *targets.Config
	selected Selections
	status   string
	applying bool
	loaded   bool
}

func New(cfg *targets.Config) Model {
	sp := spinner.New()
	sp.Spinner = spinner.Line
	m := Model{
		active:   tabSchemes,
		expanded: -1,
		inList:   false,
		lists:    map[tab]list.Model{},
		spinner:  sp,
		cfg:      cfg,
		status:   "Loading…",
	}
	del := newCompactDelegate()

	for t := tabSchemes; t < tabCount; t++ {
		l := list.New([]list.Item{}, del, maxWidth-6, 8)
		l.SetShowStatusBar(false)
		l.SetFilteringEnabled(true)
		l.SetShowHelp(false)
		l.SetShowTitle(false)
		l.SetShowPagination(true)
		l.KeyMap.Quit.SetEnabled(false)
		m.lists[t] = l
	}
	return m
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, loadDataCmd(m.cfg))
}

func loadDataCmd(cfg *targets.Config) tea.Cmd {
	return func() tea.Msg {
		var schemes []string
		if cfg.SchemesDir != "" {
			// CLI override - use single directory
			schemes, _ = targets.ScanSchemesDir(cfg.SchemesDir)
		} else {
			// Default - scan all standard locations
			schemes = targets.ScanAllSchemes()
		}
		return dataLoadedMsg{
			schemes: schemes,
			icons:   targets.ScanIconThemes(),
			walls:   targets.ScanWallpapers(),
		}
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = min(msg.Width, maxWidth)
		m.height = min(msg.Height, maxHeight)
		m = m.resizeLists()
		return m, nil

	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case dataLoadedMsg:
		m.schemes, m.icons, m.walls = msg.schemes, msg.icons, msg.walls
		m.status = "Ready"
		m.loaded = true

		m.lists[tabSchemes] = rebuildList(m.lists[tabSchemes], msg.schemes)
		m.lists[tabIcons] = rebuildList(m.lists[tabIcons], msg.icons)
		m.lists[tabWalls] = rebuildList(m.lists[tabWalls], msg.walls)
		return m, nil

	case applyDoneMsg:
		m.applying = false
		if msg.err != nil {
			m.status = "Apply failed: " + firstLine(msg.err.Error())
		} else {
			m.status = "Applied successfully!"
		}
		return m, nil

	case tea.KeyMsg:
		k := msg.String()

		// Global keys
		switch k {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "a":
			if m.applying || m.selected.Scheme == "" {
				if m.selected.Scheme == "" {
					m.status = "Select a scheme first"
				}
				return m, nil
			}
			m.applying = true
			m.status = "Applying…"
			return m, tea.Batch(m.spinner.Tick, applyCmd(m.cfg, m.selected))
		}

		// Navigation depends on whether we're in a list or at panel titles
		if m.inList && m.expanded >= 0 {
			switch k {
			case "left", "esc":
				m.inList = false
				return m, nil
			case "enter":
				m = m.selectCurrentItem()
				return m, nil
			case "up", "down", "j", "k", "pgup", "pgdown", "home", "end":
				l := m.lists[m.expanded]
				l, cmd = l.Update(msg)
				m.lists[m.expanded] = l
				return m, cmd
			default:
				l := m.lists[m.expanded]
				l, cmd = l.Update(msg)
				m.lists[m.expanded] = l
				return m, cmd
			}
		} else {
			switch k {
			case "up", "k":
				if m.active > 0 {
					m.active--
				}
				return m, nil
			case "down", "j":
				if m.active < tabCount-1 {
					m.active++
				}
				return m, nil
			case "right", "enter", "l":
				m.expanded = m.active
				m.inList = true
				return m, nil
			case "left", "h":
				if m.expanded == m.active {
					m.expanded = -1
				}
				return m, nil
			}
		}
	}

	return m, cmd
}

func rebuildList(l list.Model, items []string) list.Model {
	lis := make([]list.Item, 0, len(items))
	for _, it := range items {
		lis = append(lis, item{title: it})
	}
	l.SetItems(lis)
	return l
}

func applyCmd(cfg *targets.Config, sel Selections) tea.Cmd {
	return func() tea.Msg {
		// Find scheme path
		var schemePath string
		var err error
		if cfg.SchemesDir != "" {
			schemePath = cfg.SchemesDir + "/" + sel.Scheme + ".yaml"
		} else {
			schemePath, err = targets.FindScheme(sel.Scheme)
			if err != nil {
				return applyDoneMsg{err: err}
			}
		}

		// Parse scheme
		s, err := scheme.Parse(schemePath)
		if err != nil {
			return applyDoneMsg{err: err}
		}

		// Set optional selections
		cfg.IconTheme = sel.IconTheme
		cfg.Wallpaper = sel.Wallpaper

		// Apply
		err = targets.Apply(cfg, s)
		return applyDoneMsg{err: err}
	}
}

func (m Model) selectCurrentItem() Model {
	if m.expanded < 0 {
		return m
	}
	l := m.lists[m.expanded]
	it, ok := l.SelectedItem().(item)
	if !ok {
		return m
	}

	switch m.expanded {
	case tabSchemes:
		m.selected.Scheme = it.title
		m.status = "Scheme: " + it.title
	case tabIcons:
		m.selected.IconTheme = it.title
		m.status = "Icons: " + it.title
	case tabWalls:
		m.selected.Wallpaper = it.title
		m.status = "Wallpaper: " + it.title
	}
	return m
}

func (m Model) resizeLists() Model {
	if m.width <= 0 || m.height <= 0 {
		return m
	}
	listHeight := m.height - 16
	if listHeight < 5 {
		listHeight = 5
	}
	if listHeight > 12 {
		listHeight = 12
	}

	listWidth := m.width - 6
	if listWidth < 30 {
		listWidth = 30
	}

	for t, l := range m.lists {
		l.SetSize(listWidth, listHeight)
		m.lists[t] = l
	}
	return m
}

// ─────────────────────────────────────────────────────────────────────────────
// View rendering
// ─────────────────────────────────────────────────────────────────────────────

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("15")).
			Background(lipgloss.Color("62")).
			Padding(0, 1).
			MarginBottom(1)

	panelTitleFocused = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("0")).
				Background(lipgloss.Color("7"))

	panelTitleNormal = lipgloss.NewStyle().
				Foreground(lipgloss.Color("7"))

	panelTitleExpanded = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("10"))

	selLabelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("8")).
			Width(12)

	selValueStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("15"))

	statusStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("11")).
			Italic(true)

	dimStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
)

func (m Model) View() string {
	var b strings.Builder

	// Title
	title := titleStyle.Render("Base16 Theme Changer")
	b.WriteString(title + "\n")

	// Current selections
	b.WriteString(m.renderSelections())
	b.WriteString("\n")

	// Panel list
	b.WriteString(m.renderPanels())
	b.WriteString("\n")

	// Help commands
	b.WriteString(m.renderHelp())
	b.WriteString("\n")

	// Status line
	status := m.status
	if m.applying {
		status = m.spinner.View() + " " + status
	}
	b.WriteString(statusStyle.Render(status))

	content := b.String()
	return lipgloss.NewStyle().
		Width(m.width).
		MaxWidth(maxWidth).
		Render(content)
}

func (m Model) renderSelections() string {
	var lines []string
	lines = append(lines, dimStyle.Render("─── Current Selection ───"))

	selections := []struct {
		label string
		value string
	}{
		{"Scheme", m.selected.Scheme},
		{"Icons", m.selected.IconTheme},
		{"Wallpaper", m.selected.Wallpaper},
	}

	for _, sel := range selections {
		label := selLabelStyle.Render(sel.label + ":")
		value := selValueStyle.Render(emptyDash(sel.value))
		lines = append(lines, label+value)
	}

	return strings.Join(lines, "\n")
}

func (m Model) renderPanels() string {
	var lines []string
	lines = append(lines, dimStyle.Render("─── Theme Panels ───"))

	for t := tabSchemes; t < tabCount; t++ {
		var indicator string
		var style lipgloss.Style

		isExpanded := m.expanded == t
		isFocused := m.active == t

		if isExpanded {
			indicator = "▼ "
			style = panelTitleExpanded
		} else {
			indicator = "▶ "
			style = panelTitleNormal
		}

		prefix := "  "
		if isFocused && !m.inList {
			prefix = "› "
			style = panelTitleFocused
		}

		count := len(m.lists[t].Items())
		countStr := dimStyle.Render(fmt.Sprintf(" (%d)", count))

		line := prefix + indicator + style.Render(tabNames[t]) + countStr
		lines = append(lines, line)

		if isExpanded {
			listView := m.lists[t].View()
			indented := indentLines(listView, "  ")
			lines = append(lines, indented)
		}
	}

	return strings.Join(lines, "\n")
}

var helpKeyStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("6")).Bold(true)
var helpDescStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("7"))

func (m Model) renderHelp() string {
	var lines []string
	lines = append(lines, dimStyle.Render("─── Commands ───"))

	commands := []struct {
		key  string
		desc string
	}{
		{"↑ ↓", "Navigate panels"},
		{"→ / Enter", "Expand panel"},
		{"← / Esc", "Collapse panel"},
		{"/", "Filter items"},
		{"A", "Apply changes"},
		{"Q", "Quit"},
	}

	for _, cmd := range commands {
		line := helpKeyStyle.Render(fmt.Sprintf("%-10s", cmd.key)) + helpDescStyle.Render(cmd.desc)
		lines = append(lines, line)
	}

	return strings.Join(lines, "\n")
}

func indentLines(s string, prefix string) string {
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		lines[i] = prefix + line
	}
	return strings.Join(lines, "\n")
}

func emptyDash(s string) string {
	if strings.TrimSpace(s) == "" {
		return "—"
	}
	return s
}

func firstLine(s string) string {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	if i := strings.IndexByte(s, '\n'); i >= 0 {
		return s[:i]
	}
	return s
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
