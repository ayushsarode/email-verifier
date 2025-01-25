package cmd

import (
	"cli/internal"
	"fmt"
	"strings"
	"time"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	email string
	successStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("2")).Bold(true)
	errorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("1")).Bold(true)
	infoStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("4")).Italic(true)
)

type model struct {
	email             string
	domain            string
	err               error
	done              bool
	formatChecked     bool
	domainChecked     bool
	reachableChecked  bool
	currentStep       int
	verificationSteps []string
}

func (m *model) Init() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return t
	})
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		return m, nil

	case time.Time:
		if m.done {
			return m, tea.Quit
		}
		if m.err != nil {
			return m, tea.Quit
		}

		
		if !m.formatChecked {
			if !internal.IsValidEmail(m.email) {
				m.err = fmt.Errorf("invalid email format: %s", m.email)
				return m, tea.Quit
			}
			m.formatChecked = true
			m.currentStep = 1
			m.verificationSteps = append(m.verificationSteps,
				successStyle.Render("\n[✓] Email format validated"))
			return m, tea.Tick(time.Second, func(t time.Time) tea.Msg { return t })
		}

		
		if !m.domainChecked {
			parts := strings.Split(m.email, "@")
			if len(parts) != 2 {
				m.err = fmt.Errorf("invalid email format")
				return m, tea.Quit
			}
			m.domain = parts[1]
			if !internal.IsDomainValid(m.domain) {
				m.err = fmt.Errorf("invalid domain: %s", m.domain)
				return m, tea.Quit
			}
			m.domainChecked = true
			m.currentStep = 2
			m.verificationSteps = append(m.verificationSteps,
				successStyle.Render("[✓] Domain validated"))
			return m, tea.Tick(time.Second, func(t time.Time) tea.Msg { return t })
		}


		if !m.reachableChecked {
			if !internal.IsEmailReachable(m.email) {
				m.err = fmt.Errorf("email is unreachable: %s", m.email)
				return m, tea.Quit
			}
			m.reachableChecked = true
			m.currentStep = 3
			m.verificationSteps = append(m.verificationSteps,
				successStyle.Render("[✓] Email reachability confirmed"))
			return m, tea.Tick(time.Second, func(t time.Time) tea.Msg { return t })
		}

		m.done = true
		return m, tea.Quit

	default:
		return m, nil
	}
}

func (m *model) View() string {
	if m.err != nil {
		return errorStyle.Render(fmt.Sprintf("❌ Error: %v\n", m.err))
	}
	if m.done {
		output := infoStyle.Render("Verifying email...\n")
		output += ("Steps: [1/3] [2/3] [3/3]\n") 
		output += strings.Join(m.verificationSteps, "\n") + "\n"
		output += successStyle.Render(fmt.Sprintf("✅ Email %s is valid and reachable!\n", m.email))
		return output
	}
	
	currentStepView := "Verifying: "
	for i := 1; i <= 3; i++ {
		if i <= m.currentStep {
			currentStepView += "✓ "
		} else {
			currentStepView += "- "
		}
	}

	return infoStyle.Render("Verifying email...\n") +
		currentStepView + "\n" +
		strings.Join(m.verificationSteps, "\n") + "\n"
}

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify an email address",
	Run: func(cmd *cobra.Command, args []string) {
		if email == "" {
			fmt.Println(errorStyle.Render("❌ Please provide an email using --email flag"))
			return
		}
		p := tea.NewProgram(&model{
			email:            email,
			verificationSteps: []string{},
		}, tea.WithAltScreen())

		finalModel, err := p.Run()
		if err != nil {
			fmt.Println(errorStyle.Render("❌ An error occurred: " + err.Error()))
			return
		}
		if m, ok := finalModel.(*model); ok {
			fmt.Print(m.View())
		} else {
			fmt.Println(errorStyle.Render("❌ Unexpected error occurred while verifying the email."))
		}
	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)
	verifyCmd.Flags().StringVarP(&email, "email", "e", "", "Email to verify")
}