package model

type User struct {
    ID        int                     `json:"id"`
    CreatedOn string                  `json:"createdOn"`
    Name      string                  `json:"name"`
    UserName  string                  `json:"userName"`
    Email     string                  `json:"email"`
    Password  string                  `json:"password"`
}
