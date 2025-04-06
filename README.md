# Go URL Shortener

An efficient URL shortener service built with Go, featuring MongoDB integration for URL storage and straightforward HTTP endpoints for URL shortening and redirection.

## Overview

- Create shortened URLs from long URLs
- Redirect shortened URLs to their original destinations
- Persistent storage using MongoDB
- Environment variable configuration
- Base64 encoded short URL generation

## Installation

1. Clone the repository:
```bash
git clone https://github.com/sametpalitci/go-url-shortener.git
cd go-url-shortener
```

2. Install dependencies:
```bash
go mod download
```

3. Set up your environment variables by copying the example file:
```bash
cp .env.example .env
```

4. Update the `.env` file with your configuration.

## Usage

1. Start the server:
```bash
go run main.go
```

2. Create a shortened URL:
```bash
curl -X POST -d "url=https://example.com" http://localhost:8080/create
```

3. Use the shortened URL:
Simply visit the shortened URL in your browser, and you will be redirected to the original URL.

## API Endpoints

- `POST /create`
  - Creates a shortened URL
  - Parameters:
    - `url`: The original URL to be shortened
  - Returns: The shortened URL

- `GET /{shortURL}`
  - Redirects to the original URL
  - Parameters:
    - `shortURL`: The shortened URL code

## Contact

A.Samet Palitci - [@asametpalitci](https://twitter.com/asametpalitci) - abdulsametpalitci@gmail.com

Project Link: [https://github.com/sametpalitci/go-url-shortener](https://github.com/sametpalitci/go-url-shortener)