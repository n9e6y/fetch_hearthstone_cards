# Hearthstone API Fetch

This Go project fetches Hearthstone card data from an API source and saves it to both CSV and JSON formats.

## Features

- Fetches Hearthstone card data from a specified API
- Saves card data in CSV format
- Saves card data in JSON format
- Written in Go for efficient performance

## Prerequisites

- Go 1.15 or higher
- Internet connection to access the Hearthstone API

## Installation

1. Clone this repository:
   ```
   git clone https://github.com/n9e6y/fetch_hearthstone_cards.git
   ```
2. Navigate to the project directory:
   ```
   cd fetch_hearthstone_cards/cmd
   ```
3. Install dependencies (if any):
   ```
   go mod tidy
   ```

## Usage

Run the program with the following command:

```
go run main.go
```

The program will fetch the Hearthstone card data and save it to:
- `cards.csv`: CSV format
- `cards.json`: JSON format

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.


## Acknowledgments

- Hearthstone API provider (https://hearthstonejson.com/)
- Go community

## Contact

If you have any questions or feedback, please open an issue in this repository.