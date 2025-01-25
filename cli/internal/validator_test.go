package internal

import (
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		{"valid@example.com", true},
		{"invalid-email", false},
		{"missing@domain", false},
		{"@nodomain.com", false},
		{"user.name+tag@domain.co", true},
		{"user@sub.domain.com", true},
		{"user@domain-with-dash.com", true},
	}

	for _, test := range tests {
		result := IsValidEmail(test.email)
		if result != test.expected {
			t.Errorf("IsValidEmail(%s) = %v; expected %v", test.email, result, test.expected)
		}
	}
}

func TestIsDomainValid(t *testing.T) {
	tests := []struct {
		domain   string
		expected bool
	}{
		{"gmail.com", true},     
		{"example.com", true},    
		{"invalid-domain.xyz", false}, 
		{"nonexistent.com", false},
	}

	for _, test := range tests {
		result := IsDomainValid(test.domain)
		if result != test.expected {
			t.Errorf("IsDomainValid(%s) = %v; expected %v", test.domain, result, test.expected)
		}
	}
}

func TestIsEmailReachable(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		{"test@gmail.com", false}, 
		{"admin@example.com", false}, 
	}

	for _, test := range tests {
		result := IsEmailReachable(test.email)
		if result != test.expected {
			t.Errorf("IsEmailReachable(%s) = %v; expected %v", test.email, result, test.expected)
		}
	}
}
