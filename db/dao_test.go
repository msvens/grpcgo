package db

import (
	"github.com/msvens/grpcgo/service"
	"testing"
)

const (
	email = "some@emial.com"
	name = "first last"
)

var req = service.AddUserRequest{Email:email, Name:name}

func createAndClear(t *testing.T) (*PgUserDao){
	dao, err := NewPgDao()
	if err != nil {
		t.Errorf("could not connect %v", err)
		return nil
	}
	err = dao.DropTables()
	if err != nil {
		t.Errorf("could not drop tables %v", err)
		return nil
	}
	err = dao.CreateTables()
	if err != nil {
		t.Errorf("could not create tables %v", err)
	}
	return dao
}

func createAndAdd(t *testing.T) (*PgUserDao) {
	dao := createAndClear(t)
	_, err := dao.Create(&req)
	if err != nil {
		t.Errorf("could not create user: %v", err)
	}
	return dao
}

func TestAddUser(t *testing.T) {
	dao := createAndClear(t)
	id, err := dao.Create(&req)
	if err != nil {
		t.Errorf("could not create user: %v", err)
	}
	if id < 1 {
		t.Errorf("id cannot be less than 1: %v", id)
	}
	//test that we cannot add the same user twice
	_, err = dao.Create(&req)
	if err == nil {
		t.Errorf("adding the same useer should yield an error")
	}
}

func TestList(t *testing.T) {
	dao := createAndClear(t)
	ul, err := dao.List()
	if err != nil {
		t.Errorf("could not list: %v", err)
	}
	length := len(ul.Users)
	if length != 0 {
		t.Errorf("number of users should be 0 got: %v", length)
	}
	_, _ = dao.Create(&req)
	ul, err = dao.List()
	length = len(ul.Users)
	if length != 1 {
		t.Errorf("number of users should be 1 got: %v", length)
	}
}

func TestDelete(t *testing.T) {
	dao := createAndClear(t)
	id, err := dao.Create(&req)
	if err != nil {
		t.Errorf("could not create user: %v", err)
	}
	delReq := service.DeleteUserRequest{Id: id}
	rows, err := dao.Delete(&delReq)
	if err != nil {
		t.Errorf("error when deleting: %v", err)
	}
	if rows != 1 {
		t.Errorf("expected 1 row to be affected got: %v", rows)
	}
	//trying to delete the same user again
	rows, err = dao.Delete(&delReq)
	if err != nil {
		t.Errorf("error when deleting: %v", err)
	}
	if rows > 0 {
		t.Errorf("expected 0 rows to be affected got: %v", rows)
	}
}

func TestDeleteAll(t *testing.T) {
	dao := createAndAdd(t)
	_, err := dao.Create(&service.AddUserRequest{Email:"some1@email", Name:"First1 Last1"})
	if err != nil {
		t.Errorf("could not add user: %v", err)
	}
	rows, err := dao.DeleteAll()
	if err != nil {
		t.Errorf("could not delete all: %v", err)
	}
	if rows != 2 {
		t.Errorf("expected 2 rows to be deleted got: %v", rows)
	}
	rows, _ = dao.DeleteAll()
	if rows != 0 {
		t.Errorf("expected 0 rows to be deleted got: %v", rows)
	}
}

func TestGet(t *testing.T) {
	dao := createAndClear(t)
	_, err := dao.Get(1)
	switch err {
	case nil:
		t.Errorf("expected NoSuchUserError got no error")
	case NoSuchUser:
	default:
		t.Errorf("expected NoSuchUserError got: %v", err)
	}

}


