# Email Verifier CLI

A simple CLI tool built with Go to verify email addresses by checking their format, domain validity, and reachability.

<img src="https://api.visitorbadge.io/api/visitors?path=https%3A%2F%2Fgithub.com%2Fayushsarode%2Femail-verifier&label=visitors&countColor=%2337d67a&style=for-the-badge&labelStyle=upper" />

## Features
- Validate email format using regex.
- Check if the email domain has valid MX records.
- Verify the reachability of the email via SMTP.
- Interactive CLI using Bubble Tea.

## Prerequisites
- Go 1.22+
- Docker (optional)

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
go run main.go verify --email=your-email@example.com
```

Example output:
```
[1/3] Checking email format...
[2/3] Checking domain validity...
[3/3] Checking email reachability...
✅ Email your-email@example.com is valid and reachable!
```

## Docker Usage

### Build Docker Image
```bash
docker build -t ayushsarode777/email-verifier .
```

### Run Container
```bash
docker run -it ayushsarode777/email-verifier verify --email=your-email@example.com
```

### Pull from Docker Hub
If you have pushed the image to Docker Hub, you can pull and run it directly:
```bash
docker pull ayushsarode777/email-verifier:latest
docker run -it ayushsarode777/email-verifier ./app verify --email=your-email@example.com
```

## Testing

Run unit tests with:
```bash
go test ./cli/internal
```

## License

This project is licensed under the MIT License.

## Contribution

Feel free to submit issues or pull requests to enhance the project.

