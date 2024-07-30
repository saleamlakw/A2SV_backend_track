package controllers

import (
    "bufio"
    "fmt"
    "example/LibraryManagmnetSystem/models"
    "example/LibraryManagmnetSystem/services"
    "os"
    "strconv"
)

var library services.LibraryManager

func init() {
    library = services.NewLibrary()
}

func RunLibrary() {
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Println("\nLibrary Management System")
        fmt.Println("1. Add a new book")
        fmt.Println("2. Remove an existing book")
        fmt.Println("3. Borrow a book")
        fmt.Println("4. Return a book")
        fmt.Println("5. List all available books")
        fmt.Println("6. List all borrowed books by a member")
        fmt.Println("7. Exit")
        fmt.Print("Enter your choice: ")

        scanner.Scan()
        choice := scanner.Text()

        switch choice {
        case "1":
            addBook(scanner)
        case "2":
            removeBook(scanner)
        case "3":
            borrowBook(scanner)
        case "4":
            returnBook(scanner)
        case "5":
            listAvailableBooks()
        case "6":
            listBorrowedBooks(scanner)
        case "7":
            fmt.Println("Exiting the Library Management System.")
            return
        default:
            fmt.Println("Invalid choice. Please try again.")
        }
    }
}

func addBook(scanner *bufio.Scanner) {
    fmt.Print("Enter book ID: ")
    scanner.Scan()
    id, _ := strconv.Atoi(scanner.Text())

    fmt.Print("Enter book title: ")
    scanner.Scan()
    title := scanner.Text()

    fmt.Print("Enter book author: ")
    scanner.Scan()
    author := scanner.Text()

    book := models.Book{
        ID:     id,
        Title:  title,
        Author: author,
        Status: "Available",
    }

    library.AddBook(book)
    fmt.Println("Book added successfully!")
}

func removeBook(scanner *bufio.Scanner) {
    fmt.Print("Enter book ID to remove: ")
    scanner.Scan()
    id, _ := strconv.Atoi(scanner.Text())

    err := library.RemoveBook(id)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Book removed successfully!")
    }
}

func borrowBook(scanner *bufio.Scanner) {
    fmt.Print("Enter book ID to borrow: ")
    scanner.Scan()
    bookID, _ := strconv.Atoi(scanner.Text())

    fmt.Print("Enter member ID: ")
    scanner.Scan()
    memberID, _ := strconv.Atoi(scanner.Text())

    err := library.BorrowBook(bookID, memberID)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Book borrowed successfully!")
    }
}

func returnBook(scanner *bufio.Scanner) {
    fmt.Print("Enter book ID to return: ")
    scanner.Scan()
    bookID, _ := strconv.Atoi(scanner.Text())

    fmt.Print("Enter member ID: ")
    scanner.Scan()
    memberID, _ := strconv.Atoi(scanner.Text())

    err := library.ReturnBook(bookID, memberID)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Book returned successfully!")
    }
}

func listAvailableBooks() {
    books := library.ListAvailableBooks()
    if len(books) == 0 {
        fmt.Println("No available books.")
    } else {
        fmt.Println("Available books:")
        for _, book := range books {
            fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
        }
    }
}

func listBorrowedBooks(scanner *bufio.Scanner) {
    fmt.Print("Enter member ID: ")
    scanner.Scan()
    memberID, _ := strconv.Atoi(scanner.Text())

    books, err := library.ListBorrowedBooks(memberID)
    if err != nil {
        fmt.Println("Error:", err)
    } else if len(books) == 0 {
        fmt.Println("No borrowed books for this member.")
    } else {
        fmt.Println("Borrowed books:")
        for _, book := range books {
            fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
        }
    }
}
