package books

import (
	"database/sql"
	"fmt"
	dto "learn-microservices-server/src/app/dto/books"
)

type openLibraryService struct {
	db *sql.DB
}

type OpenLibraryServices interface {
	GetBooksBySubject(subject string) (*dto.GetBooksRespDTO, error)
}

func NewIntegOpenLibrary(db *sql.DB) OpenLibraryServices {
	return &openLibraryService{
		db: db,
	}
}

func (s *openLibraryService) GetBooksBySubject(subject string) (*dto.GetBooksRespDTO, error) {
	var response dto.GetBooksRespDTO

	// Set response
	response.Name = subject
	response.SubjectType = "subject"

	// Query for books
	query := `SELECT books.id, books.title, books.cover_id, books.edition_count, authors.name 
          FROM books 
          INNER JOIN authors ON books.author_id = authors.id 
          WHERE books.subject = $1`

	rows, err := s.db.Query(query, subject)
	if err != nil {
		return nil, fmt.Errorf("failed to query books from database: %v", err)
	}
	defer rows.Close()

	// Iterate over rows for each book
	for rows.Next() {
		var book dto.WorkDTO
		var authorName string

		if err := rows.Scan(&book.ID, &book.Title, &book.CoverID, &book.EditionCount, &authorName); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		// Add authors
		book.Authors = append(book.Authors, &dto.AuthorDTO{Name: authorName})

		// Add book
		response.Works = append(response.Works, &book)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %v", err)
	}

	return &response, nil
}
