package repositories

import (
	"backend_article/models"
	"math"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	FindArticle(page int, limit int, search string) ([]models.Article, int, error)
	GetArticle(ID int) (models.Article, error)
	CreateArticle(article models.Article) (models.Article, error)
	UpdateArticle(article models.Article) (models.Article, error)
	DeleteArticle(article models.Article, ID int) (models.Article, error)
}

func RepositoryArticle(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindArticle(page int, limit int, search string) ([]models.Article, int, error) {
	var articles []models.Article
	var count int64
	offset := (page - 1) * limit

	query := r.db.Model(&models.Article{}).Offset(offset).Limit(limit)

	if search != "" {
		query = query.Where("title LIKE ?", "%"+search+"%")
	}

	// Menambahkan urutan berdasarkan tanggal, dari yang terbaru ke yang terlama
	query = query.Order("posted_at DESC")

	result := query.Find(&articles)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	// Menghitung jumlah total artikel
	countQuery := r.db.Model(&models.Article{})

	if search != "" {
		countQuery = countQuery.Where("title LIKE ?", "%"+search+"%")
	}
	countQuery.Count(&count)

	totalPages := int(math.Ceil(float64(count) / float64(limit)))

	return articles, totalPages, nil
}

func (r *repository) GetArticle(ID int) (models.Article, error) {
	var article models.Article
	err := r.db.First(&article, ID).Error

	return article, err
}

func (r *repository) CreateArticle(article models.Article) (models.Article, error) {
	err := r.db.Create(&article).Error

	return article, err
}

func (r *repository) UpdateArticle(article models.Article) (models.Article, error) {
	err := r.db.Save(&article).Error

	return article, err
}

func (r *repository) DeleteArticle(article models.Article, ID int) (models.Article, error) {
	err := r.db.Delete(&article, ID).Scan(&article).Error

	return article, err
}
