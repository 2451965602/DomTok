namespace go api.user
include "../model.thrift"

//
struct RegisterRequest {
    1: required string name
    2: required string password
    3: required string email
}

struct RegisterResponse {
}

struct LoginRequest {
    1: required string name
    2: required string password
}

struct LoginResponse {
    1: model.BaseResp base,
    2: model.UserInfo user,
}


service UserService {
    RegisterResponse Register(1: RegisterRequest req)(api.post = "api/v1/user/register"),
    LoginResponse Login(1: LoginRequest req)(api.post = "api/v1/user/login")
}
