package users

//
//const (
//	findUserByEmailQuery = `select name, email from users where email like=$1`
//)
//
//type FindUserServer struct {
//	db *sql.DB
//}
//
//func NewFindUserByEmail(db *sql.DB) *FindUserServer {
//	return &FindUserServer{db: db}
//}
//
//func (u *FindUserServer) GetByEmail(ctx context.Context, req *pb.FindUsersByEmailRequest) (*pb.FindUsersByEmailResponse, error) {
//	var user *pb.FindUsersByEmailResponse
//	rows, err := u.db.Query(findUserByEmailQuery, req.Email)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//	getUsers := &pb.FindUsersByEmailResponse{Id: user.GetId(), Name: user.Name, Email: user.Email}
//	//for rows.Next() {
//	//	if err := rows.Scan(&user.Id, &user.Name, &user.Email); err != nil {
//	//		return nil, err
//	//	}
//	//	users = append(users, user)
//	//}
//
//	return getUsers, nil
//}
//
//func (u *FindUserServer) MustEmbedUnimplementedGetUsersServer() {}
