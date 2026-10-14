package store  //ทุกไฟล์ที่อยุ่ใน store ใช้ package store 

import (
    "errors"

    "github.com/jackc/pgx/v5/pgxpool"
)

var ErrNotFound = errors.New("not found")

type Store struct {
    Pool *pgxpool.Pool
}