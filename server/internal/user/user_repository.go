package user

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	var lastInsertId int
	query := "INSERT INTO users(username, password, email) VALUES ($1, $2, $3) returning id"
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Password, user.Email).Scan(&lastInsertId)
	if err != nil {
		return &User{}, err
	}

	user.ID = int64(lastInsertId)
	return user, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	u := User{}
	query := "SELECT id, email, username, password FROM users WHERE email = $1"
	err := r.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.Email, &u.Username, &u.Password)
	if err != nil {
		return &User{}, nil
	}

	return &u, nil
}

func (r *repository)CheckUsernameExist(ctx context.Context , username string) (bool,error){
	query:="SELECT username FROM users WHERE username = $1"
	rows,err:=r.db.QueryContext(ctx , query , username)
	 
	 //check if there is error first
	 if err != nil {
        return false, err
    }
    defer rows.Close()
	//it's mean that the username is already exist
	if rows.Next() {
        return false, nil
    }
	//it's mean that the username doesn't  exist
	return true,nil
}

func (r *repository)CheckEmailExist(ctx context.Context , email string)(bool,error){
	query:="SELECT email FROM users WHERE email = $1"
	rows, err := r.db.QueryContext(ctx, query, email)
	 //check if there is error first 
	 if err != nil {
        return false, err
    }
    defer rows.Close()
	//mean that the email does  exist
	if rows.Next() {
        return false, nil
    }
	//mean that the email doesn't  exist 
	return true,nil
}

