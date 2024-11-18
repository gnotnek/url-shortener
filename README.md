# URL Shortener

A simple URL shortener service written in Go.

## Features

- Shorten long URLs
- Redirect to original URLs
- Track URL usage statistics

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/gnotnek/url-shortener.git
    ```
2. Navigate to the project directory:
    ```sh
    cd url-shortener
    ```
3. Install dependencies:
    ```sh
    go mod tidy
    ```

## Usage

1. Start the server:
    ```sh
    go run main.go
    ```
2. Open your browser and navigate to `http://localhost:9808`.

## API Endpoints

- `POST /shorten`
    - Request: 
        ```json
        {
            "url": "https://example.com",
            "user_id" : "123e4567-e89b-12d3-a456-426614174000"
        }
        ```
    - Response:
        ```json
        {
            "short_url": "http://localhost:9808/abc123"
        }
        ```

- `GET /{short_url}`
    - Redirects to the original URL.

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Commit your changes (`git commit -am 'Add new feature'`).
4. Push to the branch (`git push origin feature-branch`).
5. Create a new Pull Request.

## Acknowledgements

- [Gin Gonic](https://github.com/gin-gonic/gin) - HTTP web framework for Go
