package services

import (
	"errors"
	"github.com/saleamlakw/A2SV_backend_track/models"
	"fmt"
)
type LibraryManager interface{
	AddBook(book models.Book)
	RemoveBook(bookID int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) ([]models.Book,error)
}

type Libray struct{
	books map[int]models.Book
	members map[int]models.Member
} 
func NewLibrary() *Libray {
    return &Libray{
        books:   make(map[int]models.Book),
        members: make(map[int]models.Member),
    }
}

func (l Libray)AddBook(book models.Book){
	l.books[book.ID]=book
	fmt.Println("----",l.books)

}

func (l Libray)RemoveBook(bookID int)error{
	_,exists:=l.books[bookID]
	if exists{
		delete(l.books,bookID)
		return nil
	}else{
		return errors.New("the book doesn't exist in the library")
	}
	
}
func (l Libray) BorrowBook(bookID int, memberID int) error{
	book,bookExists:=l.books[bookID]
	if bookExists{
		if book.Status!="Available"{
			return errors.New("book not Available")
		}

		member,memberExists:=l.members[memberID]
		if memberExists{
			member.BorrowedBooks=append(member.BorrowedBooks,book)
			book.Status="Borrowed"
		}else{
			return errors.New("member not found")
		}
		}else{
			return errors.New("book not found")

	}
	return nil
}

func (l Libray) ReturnBook(bookID int, memberID int) error{
	member,memberExists:=l.members[memberID]
	if !memberExists{
		return errors.New("member not found")
	}
	bookindex:=-1
	for ind,val :=range member.BorrowedBooks{
		if bookID==val.ID{
			bookindex=ind
			break
		}
	}
	if bookindex==-1{
		return errors.New("book doesnt exist")
	}
	book:=member.BorrowedBooks[bookindex]
	member.BorrowedBooks=append(member.BorrowedBooks[:bookindex],member.BorrowedBooks[bookindex+1:]...)
	book.Status = "Available"
	return nil
}

func (l Libray) ListAvailableBooks() []models.Book{
	AvailableBooks:=[]models.Book{}

	for _,book := range l.books{
		if book.Status=="Available"{
			AvailableBooks=append(AvailableBooks, book)
		}
	}
	return AvailableBooks
}

func (l Libray) ListBorrowedBooks(memberID int) ([]models.Book,error){
	member,meberExists:=l.members[memberID]
	if !meberExists{
		fmt.Println("member not found")
	}
	return member.BorrowedBooks,nil
}