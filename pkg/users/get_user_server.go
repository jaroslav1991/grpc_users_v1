package users

const (
// findUserByIdQuery = `select name, email from users where id=$1`
)

//type GetUserServer struct {
//	db *sql.DB
//}
//
//func NewGetUserServer(db *sql.DB) *GetUserServer {
//	return &GetUserServer{db: db}
//}
////
//func (u *GetUserServer) GetById(ctx context.Context, req *pb.FindUserByIdRequest) (*pb.FindUserByIdResponse, error) {
//	var user pb.FindUsersByEmailResponse
//	rows, err := u.db.Query(findUserByIdQuery, req.Id)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//	getUser := &pb.FindUserByIdResponse{Id: req.GetId(), Name: user.Name, Email: user.Email}
//	for rows.Next() {
//		if err := rows.Scan(&user.Id, &user.Name, &user.Email); err != nil {
//			return nil, err
//		}
//	}
//	return getUser, nil
//}
//
//func (u *GetUserServer) MustEmbedUnimplementedGetUserServer() {}
