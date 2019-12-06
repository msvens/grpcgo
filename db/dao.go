package db

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	pb "github.com/msvens/grpcgo/service"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "test"
	password = "test123"
	dbname   = "testdb"
)

type UserDao interface {

	Create(u *pb.AddUserRequest) (int64, error)
	Delete(u *pb.DeleteUserRequest) (int64, error)
	DeleteAll() (int64, error)
	Get(id int64) (*pb.GetUserResponse, error)
	List() (*pb.ListUserResponse, error)
	Close() error
	Connect() error
	CreateTables() error
	DropTables() error
}


type PgUserDao struct {
	Db *sql.DB
}

var NoSuchUser = errors.New("No such user Id")
var DuplicateEmail = errors.New("Duplicate email")

func NewPgDao() (*PgUserDao, error) {
	var dao PgUserDao
	err := dao.Connect()
	if err != nil {
		return nil, err
	}
	return &dao, nil
}

func (dao *PgUserDao) Connect() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	dao.Db = db
	return nil
}

func (dao *PgUserDao) CreateTables() error {
	const stmt = `
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	name TEXT,
	email TEXT UNIQUE NOT NULL
);`
	_,err := dao.Db.Exec(stmt)
	return err
}

func (dao *PgUserDao) DropTables() error {
	var stmt = "DROP TABLE IF EXISTS USERS;"
	_, err := dao.Db.Exec(stmt)
	return err
}

func (dao *PgUserDao) Close() error {
	return dao.Db.Close()
}

func (dao *PgUserDao) Delete(u *pb.DeleteUserRequest) (int64, error) {
	const stmt = "DELETE FROM users WHERE id = $1"
	res, err := dao.Db.Exec(stmt, u.Id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (dao *PgUserDao) DeleteAll() (int64, error) {
	res, err := dao.Db.Exec("DELETE FROM USERS")
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (dao *PgUserDao) Create(u *pb.AddUserRequest) (int64, error) {
	const stmt = "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING ID;"
	var id int64
	err := dao.Db.QueryRow(stmt,u.Name,u.Email).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id,nil
}

func (dao *PgUserDao) Get(id int64) (*pb.GetUserResponse, error) {
	const stmt = "SELECT * FROM users WHERE id = $1"
	resp := pb.GetUserResponse{User: new(pb.UserResponse)}
	err := dao.Db.QueryRow(stmt, id).Scan(&resp.User.Id, &resp.User.Name, &resp.User.Email)
	switch err {
	case sql.ErrNoRows:
		return nil, NoSuchUser
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

func (dao *PgUserDao) List() (*pb.ListUserResponse, error) {
	const stmt = "SELECT * FROM users"
	rows, err := dao.Db.Query(stmt)
	if err != nil {
		return nil, err
	}
	users := make([]*pb.UserResponse, 10)
	cnt := 0
	ulist := pb.ListUserResponse{}
	for ; rows.Next(); cnt++ {
		if cap(users) <= cnt {
			n := make([]*pb.UserResponse, len(users)*2)
			copy(n, users)
			users = n
		}
		users[cnt] = new(pb.UserResponse)
		err = rows.Scan(&users[cnt].Id, &users[cnt].Name, &users[cnt].Email)
		if err != nil {
			return nil, err
		}
	}
	ulist.Users = users[:cnt]
	return &ulist, nil
}









