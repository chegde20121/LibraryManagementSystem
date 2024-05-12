#!/bin/bash

# Function to build the project with unit tests and display coverage
build_with_tests() {
    # Run unit tests and generate coverage profile
    go test -coverprofile=coverage.out ./...

    # Display coverage
    go tool cover -func=coverage.out

    # Display HTML coverage report
    go tool cover -html=coverage.out

    # Build the binary
    go build  cmd/library_management.go
}

# Function to build the project and run the application
build_and_run() {
    # Build the binary
   go build  cmd/library_management.go

    # Check if the binary exists
    if [ -f "./library_management" ]; then
        # Run the binary
        ./library_management
    else
        echo "Binary not found. Please build the project first."
    fi
}

# Function to build the project without running unit tests
build_only() {
    # Build the binary
     go build  cmd/library_management.go
}

# Main function
main() {
    # Display options
    echo "Select an option:"
    echo "1. Build with unit tests and show coverage"
    echo "2. Build and run the application"
    echo "3. Build the application only"
    echo "4. Exit"

    # Read user choice
    read -p "Enter your choice: " choice

    # Process user choice
    case $choice in
        1) build_with_tests ;;
        2) build_and_run ;;
        3) build_only ;;
        4) echo "Exiting..." ;;
        *) echo "Invalid choice. Please enter a valid option." ;;
    esac
}

# Run main function
main
