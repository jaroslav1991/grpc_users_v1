package users

//
//const (
//	createUserQuery = `insert into users (name, email, password_hash) values ($1, $2, $3)`
//)
//
//type CreateUserServer struct {
//	pb.UnimplementedUsersServer
//	db *sql.DB
//}
//
//func NewCreateUserService(db *sql.DB) *CreateUserServer {
//	return &CreateUserServer{db: db}
//}

//func (u *CreateUserServer) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
//	var user pb.CreateUserResponse
//	rows, err := u.db.Query(createUserQuery, req.Name, req.Email, req.PasswordHash)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	createUser := &pb.CreateUserResponse{Name: req.GetName(), Email: req.GetEmail(), Id: user.GetId()}
//	for rows.Next() {
//		if err := rows.Scan(&user.Id); err != nil {
//			return nil, err
//		}
//	}
//	return createUser, nil
//}
