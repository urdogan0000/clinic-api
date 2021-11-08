package auth

import (
	"database/sql"
	"errors"
	"github.com/jackc/pgconn"
	pg "main/adapters/database/postgres"
	ent "main/internals/interfaces"
)

var (
	pgErr *pgconn.PgError
)

func Register(dat ent.RegisterReq) (Record int64) {
	query := "INSERT INTO clinic.auth(user_title,user_name,user_pass,phone,city,town) VALUES ($1, $2, $3, $4, $5, $6)"
	res, errDB := pg.DB.Exec(pg.Ctx, query, dat.UserTitle, dat.UserName, dat.UserPass, dat.Phone, dat.City, dat.Town)

	if errDB != nil && errDB != sql.ErrNoRows {
		if errors.As(errDB, &pgErr) {
			if pgErr.Code == "23505" {

			}
		}
		Record = -1
	} else {
		Record = 0
	}

	return res.RowsAffected()
}

func Login(dat ent.LoginReq) (Record int, auth ent.Auth) {
	Record = 1
	query := "SELECT id, user_title, user_name, city, town, is_active FROM clinic.auth WHERE user_name=$1 AND user_pass=$2"
	err := pg.DB.QueryRow(pg.Ctx, query, dat.UserName, dat.UserPass).Scan(&auth.Id, &auth.UserTitle, &auth.UserName, &auth.City, &auth.Town, &auth.IsActive)

	if err != nil && err != sql.ErrNoRows {
		Record = -1
		println(err.Error())
	} /*else if err.Error() == "no rows in result set" {
		Record = 0
	}*/
	return
}

func ForgetPassword() {

}

func Verification() {

}
