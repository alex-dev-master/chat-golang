package repository

import (
	"fmt"
	"github.com/alex-dev-master/chat-golang/pkg/model"
	"github.com/jmoiron/sqlx"
)

type ChatRubricRepository struct {
	db *sqlx.DB
}

func NewChatRubric(db *sqlx.DB) *ChatRubricRepository {
	return &ChatRubricRepository{db: db}
}

func (r *ChatRubricRepository) CreateRubric(rubric model.ChatRubric) (int64, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int64
	query := fmt.Sprintf("INSERT INTO %s (user_id, caption) values (?, ?)", "chat_rubrics")
	row, err := tx.Exec(query, rubric.UserId, rubric.Caption)

	if err != nil {
		tx.Rollback()
		return 0, err
	}
	id, err = row.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()

}

func (r *ChatRubricRepository) GetRubrics() ([]model.ChatRubricUser, error) {
	var items []model.ChatRubricUser
	query := "SELECT * FROM `chat_rubrics` LEFT JOIN `users` on `chat_rubrics`.`user_id`=`users`.`id`"
	if err := r.db.Select(&items, query); err != nil {
		return nil, err
	}

	return items, nil
}
