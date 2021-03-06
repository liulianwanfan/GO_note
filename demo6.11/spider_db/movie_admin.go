package spider_db

import "DoubanNewMovie/entity"

//{ 0 00}
func QueryAdmin(name,pwd string)(*entity.Admin,error){
	//1、执行sql语句
	row :=MovieSpiderDb.QueryRow("select admin_name, admin_phone, admin_addr from movie_admin where admin_name = ? and admin_pwd = ?",name,pwd)

	var admin entity.Admin
	err := row.Scan(&admin.Name,&admin.Phone,&admin.Address)
	if err != nil {
		return nil,err
	}
	return &admin,nil
}
