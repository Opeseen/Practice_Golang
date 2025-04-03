package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound          = errors.New("resource not found")
	QueryTimeOutDuration = time.Second * 5
	ErrConflict          = errors.New("resource already exists")
)

type Storage struct {
	Posts interface {
		GetByID(context.Context, int64) (*Post, error) // will return a pointer to a post and error
		Create(context.Context, *Post) error           // will return an error
		Delete(context.Context, int64) error           // will return an error
		Update(context.Context, *Post) error           // will return an error
		GetUserFeed(context.Context, int64, PaginatedFeedQuery) ([]PostWithMetadata, error)
	}
	Users interface {
		Create(context.Context, *sql.Tx, *User) error
		CreateAndInvite(ctx context.Context, user *User, token string, exp time.Duration) error
		GetByID(context.Context, int64) (*User, error)
		Activate(context.Context, string) error
		Delete(context.Context, int64) error
	}
	Comments interface {
		GetPostByID(context.Context, int64) ([]Comment, error)
		Create(context.Context, *Comment) error
	}
	Followers interface {
		Follow(ctx context.Context, followerID, userID int64) error
		UnFollow(ctx context.Context, followerID, userID int64) error
	}
}

// Initialize the store
func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:     &PostStore{db},
		Users:     &UserStore{db},
		Comments:  &CommentStore{db},
		Followers: &FollowerStore{db},
	}
}

// wrapper for SQL transactions
func withTx(db *sql.DB, ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	if err := fn(tx); err != nil {
		// rollback when an error occurred
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}
