# Email Verifier CLI

A simple CLI tool built with Go to verify email addresses by checking their format, domain validity, and reachability.


<img src="https://api.visitorbadge.io/api/visitors?path=https%3A%2F%2Fgithub.com%2Fayushsarode%2Femail-verifier&label=visitors&countColor=%2337d67a&style=for-the-badge&labelStyle=upper" />

## Features
- Validate email format using regex.
- Check if the email domain has valid MX records.
- Verify the reachability of the email via SMTP.
- Interactive CLI using Bubble Tea.

## Prerequisites
- Go 1.18+

## Installation

Clone the repository:
```bash
git clone https://github.com/ayushsarode/email-verifier.git
cd email-verifier/cli
```

Install dependencies:
```bash
go mod tidy
```

## Usage

Run the CLI tool:
```bash
go run main.go verify --email your-email@example.com
```

Example output:
```
[1/3] Checking email format...
[2/3] Checking domain validity...
[3/3] Checking email reachability...
✅ Email your-email@example.com is valid and reachable!
```

## Project Structure
```
email-verifier/
  └── cli/
      ├── cmd/
      │   ├── root.go
      │   ├── verify.go
      ├── internal/
      │   ├── validator_test.go
      │   ├── validator.go
      ├── main.go
      ├── go.mod
      ├── go.sum
```

## Commands

### Verify an email
```bash
go run main.go verify --email=your-email@example.com
```

### Flags
- `--email` or `-e`: The email address to be verified.

## Error Handling

Common errors handled:
- Invalid email format.
- Invalid domain without MX records.
- Unreachable email via SMTP.

## Testing

Run unit tests with:
```bash
go test ./...
```

## License

This project is licensed under the MIT License.

## Contribution

Feel free to submit issues or pull requests to enhance the project.


