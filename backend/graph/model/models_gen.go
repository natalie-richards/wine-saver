// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AddBookmarkRequest struct {
	Name     *string `json:"name,omitempty"`
	Grape    *string `json:"grape,omitempty"`
	Region   *string `json:"region,omitempty"`
	Location *string `json:"location,omitempty"`
	Notes    *string `json:"notes,omitempty"`
	Image    *string `json:"image,omitempty"`
	Username *string `json:"username,omitempty"`
}

type Bookmark struct {
	Name     *string `json:"name,omitempty"`
	Grape    *string `json:"grape,omitempty"`
	Region   *string `json:"region,omitempty"`
	Location *string `json:"location,omitempty"`
	Notes    *string `json:"notes,omitempty"`
	Image    *string `json:"image,omitempty"`
	Username *string `json:"username,omitempty"`
}

type Mutation struct {
}

type Query struct {
}

type User struct {
	ID        string  `json:"id"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Email     *string `json:"email,omitempty"`
}
