package overlay

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ModelSelectionOverlay represents a model selection dialog overlay
type ModelSelectionOverlay struct {
	// Whether the overlay has been dismissed
	Dismissed bool
	// Width of the overlay
	width int
	// Selected model (empty if none selected)
	selectedModel string
	// Callback function to be called when a model is selected
	OnSelect func(model string)
	// Callback function to be called when cancelled
	OnCancel func()
	// Custom styling options
	borderColor lipgloss.Color
}

// NewModelSelectionOverlay creates a new model selection dialog overlay
func NewModelSelectionOverlay() *ModelSelectionOverlay {
	return &ModelSelectionOverlay{
		Dismissed:   false,
		width:       50, // Default width
		borderColor: lipgloss.Color("#7D56F4"), // Purple color for model selection
	}
}

// HandleKeyPress processes a key press and updates the state
// Returns true if the overlay should be closed
func (m *ModelSelectionOverlay) HandleKeyPress(msg tea.KeyMsg) bool {
	switch msg.String() {
	case "c":
		m.Dismissed = true
		m.selectedModel = "claude"
		if m.OnSelect != nil {
			m.OnSelect("claude")
		}
		return true
	case "esc":
		m.Dismissed = true
		if m.OnCancel != nil {
			m.OnCancel()
		}
		return true
	default:
		// Ignore other keys in model selection state
		return false
	}
}

// Render renders the model selection overlay
func (m *ModelSelectionOverlay) Render(opts ...WhitespaceOption) string {
	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(m.borderColor).
		Padding(1, 2).
		Width(m.width)

	// Add the model selection instructions
	content := lipgloss.NewStyle().Bold(true).Render("Which model to launch?") + "\n\n" +
		lipgloss.NewStyle().Bold(true).Render("c") + " => claude\n\n" +
		"Press " + lipgloss.NewStyle().Bold(true).Render("esc") + " to cancel"

	// Apply the border style and return
	return style.Render(content)
}

// SetWidth sets the width of the model selection overlay
func (m *ModelSelectionOverlay) SetWidth(width int) {
	m.width = width
}

// GetSelectedModel returns the selected model
func (m *ModelSelectionOverlay) GetSelectedModel() string {
	return m.selectedModel
}
