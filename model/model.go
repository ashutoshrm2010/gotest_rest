package model


//User
type User struct {
    ID        int                     `json:"id"`
    CreatedOn string                  `json:"createdOn"`
    Name      string                  `json:"name"`
    UserName  string                  `json:"userName"`
    Email     string                  `json:"email"`
    Password  string                  `json:"password"`
}

//Contacts
type Contacts struct {
    ID        int                     `json:"id"`
    CreatedOn string                  `json:"createdOn"`
    FirstName string                  `json:"first_name"`
    LastName  string                  `json:"last_name"`
    Organization     string           `json:"organization"`
    PhoneNumber  string               `json:"phoneNumber"`
    Email  string                     `json:"email"`
    Website  string                   `json:"website"`
    CreatedBy  string                   `json:"createdBy"`
}