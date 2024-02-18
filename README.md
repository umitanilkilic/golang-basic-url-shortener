# Basic URL Shortener

Welcome to the Basic URL Shortener project! This Golang-based application provides a straightforward URL shortening service with a redirection page.

## Usage

### Shorten URL
Endpoint: `POST   /shorten-url`

**Request Body:**
```json
{
    "long_url": "https://www.umitanilkilic.com"
}
```

**Response Body:**
```json
{
    "message": "URL created successfully",
    "short_url": "http://localhost:9808/s/1786799982"
}
```

### Redirect Shortened URL
Endpoint: `GET /s/:shortUrl`


## Run

To run the application, use the following command:
```bash
go run main.go
```


Enjoy using the Basic URL Shortener! If you encounter any issues or have suggestions, feel free to contribute or reach out.

## License

This repository is licensed under the [MIT License](LICENSE). Feel free to explore, use, and modify it as per the license terms.
