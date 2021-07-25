package storage

import (
	"fmt"
	"github.com/srm-developer/standardWebServer/internal/app/models"
	"log"
)

type ArticleRepository struct {
	storage *Storage
}

var (
	tableArticle string = "articles"
)

func (ar *ArticleRepository) Create(a *models.Article) (*models.Article, error) {
	query := fmt.Sprintf("INSERT INTO %s (title, author, content) VALUES ($1, $2, $3) RETURNING id", tableArticle)
	if err := ar.storage.db.QueryRow(query, a.Title, a.Author, a.Content).Scan(&a.ID); err != nil {
		return nil, err
	}
	return a, nil
}

func (ar *ArticleRepository) Delete(id int) (*models.Article, error) {
	article, ok, err := ar.FindArticleById(id)
	if err != nil {
		return nil, err
	}

	if ok {
		query := fmt.Sprintf("DELETE FROM  %s WHERE id = $1", tableArticle)
		_, err = ar.storage.db.Exec(query, id)
		if err != nil {
			return nil, err
		}
	}

	return article, nil
}

func (ar *ArticleRepository) FindArticleById(id int) (*models.Article, bool, error) {
	users, err := ar.SelectAll()
	var founded bool
	if err != nil {
		return nil, founded, err
	}

	var articleFinded *models.Article

	for _, a := range users {
		if a.ID == id {
			articleFinded = a
			founded = true
			break
		}
	}
	return articleFinded, founded, nil
}
func (ar *ArticleRepository) SelectAll() ([]*models.Article, error) {
	query := fmt.Sprintf("SELECT * from  %s", tableArticle)
	rows, err := ar.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articles := make([]*models.Article, 0)
	for rows.Next() {
		a := models.Article{}
		err := rows.Scan(&a.ID, &a.Title, &a.Author, &a.Content)
		if err != nil {
			log.Println(err)
			continue
		}
		articles = append(articles, &a)
	}
	return articles, nil
}
