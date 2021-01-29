package repository

import (
	"fmt"
	"github.com/alex-dev-master/chat-golang/pkg/model"
	"github.com/jmoiron/sqlx"
)

type MessageRepository struct {
	db *sqlx.DB
}

func NewMessage(db *sqlx.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

func (r *MessageRepository) CreateMessage(message model.Message) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int64
	query := fmt.Sprintf("INSERT INTO %s (user_id, chat_rubric_id, content, created, disabled) values (?, ?, ?, ?, ?)", "messages")
	row, err := tx.Exec(query, message.UserId, message.ChatRubric, message.Content, message.Created, 0)

	if err != nil {
		tx.Rollback()
		return 0, err
	}
	id, err = row.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return int(id), tx.Commit()
}

func (r *MessageRepository) GetAllOfRubric(rubric int) ([]model.MessageWithUser, error) {
	var items []model.MessageWithUser
	query := "SELECT * FROM `messages` LEFT JOIN `users` on `messages`.`user_id`=`users`.`id` WHERE `messages`.`chat_rubric_id` = ?"
	if err := r.db.Select(&items, query, rubric); err != nil {
		return nil, err
	}

	return items, nil
}