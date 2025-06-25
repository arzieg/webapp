package api

type EmailResponse struct {
	Address string `json:"address"`
	Primary bool   `json:"primary"`
}

// type CreateUserRequest struct {
// 	FirstName string `json:"first_name" validate:"required_without=LastName"`
// 	LastName  string `json:"last_name" validate:"required_without=FirstName"`
// 	Email     string `json:"email" validate:"required,email"`
// }

// type UpdateUserRequest struct {
// 	FirstName *string `json:"first_name" validate:"required_without=LastName"`
// 	LastName  *string `json:"last_name" validate:"required_without=FirstName"`
// }

// PatchUserRequest defines model for PatchUserRequest.
// type PatchUserRequest struct {
// 	// First name
// 	FirstName *string `json:"first_name,omitempty"`

// 	// Last name
// 	LastName *string `json:"last_name,omitempty"`
// }

// PostUserRequest defines model for PostUserRequest.
// type PostUserRequest struct {
// 	// E-mail
// 	Email string `json:"email"`

// 	// First name
// 	FirstName string `json:"first_name"`

// 	// Last name
// 	LastName string `json:"last_name"`
// }

// UserResponse defines model for UserResponse.
type UserResponse struct {
	Id        uint            `json:"id"`
	FirstName string          `json:"first_name"`
	LastName  string          `json:"last_name"`
	Emails    []EmailResponse `json:"emails"`
}

// UsersResponse defines model for UsersResponse.
type UsersResponse []UserResponse

// UserID defines model for userID.
//type UserID string

// PostUserJSONBody defines parameters for PostUser.
//type PostUserJSONBody PostUserRequest

// PatchUserJSONBody defines parameters for PatchUser.
//type PatchUserJSONBody PatchUserRequest

// PostUserJSONRequestBody defines body for PostUser for application/json ContentType.
//type PostUserJSONRequestBody PostUserJSONBody

// PatchUserJSONRequestBody defines body for PatchUser for application/json ContentType.
//type PatchUserJSONRequestBody PatchUserJSONBody
