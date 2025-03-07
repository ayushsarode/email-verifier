# Email Verifier CLI

A command-line tool built with Go to verify email addresses.

![Visitors](https://api.visitorbadge.io/api/visitors?path=https%3A%2F%2Fgithub.com%2Fayushsarode%2Femail-verifier&label=visitors&countColor=%2337d67a&style=for-the-badge&labelStyle=upper)

## Features

- Validate email format
- Check domain MX records
- Verify email reachability via SMTP
- Interactive CLI interface

## Installation

```bash
# Clone the repository
git clone https://github.com/ayushsarode/email-verifier.git
cd email-verifier/cli

# Install dependencies
go mod tidy
```

## Usage

```bash
go run main.go verify --email=your-email@example.com
```

## Using Docker

```bash
# Pull the pre-built image
docker pull ayushsarode777/email-verifier:latest

# Run the container
docker run -it ayushsarode777/email-verifier ./app verify --email=your-email@example.com
```

## Example output:
```bash
[1/3] Checking email format...
[2/3] Checking domain validity...
[3/3] Checking email reachability...
✅ Email your-email@example.com is valid and reachable!
```

## Testing

```bash
go test ./cli/internal
```

## License

MIT License

## Contributing

Contributions welcome via issues and pull requests.
