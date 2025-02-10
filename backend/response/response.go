package response

import (
	"github.com/tejiriaustin/lema/models"
)

func SingleUserResponse(account *models.User) map[string]interface{} {
	return map[string]interface{}{
		"id":       account.ID.String(),
		"email":    account.Email,
		"fullName": account.FullName,
		"address":  SingleAddressResponse(account.Address),
	}
}

func SingleAddressResponse(address *models.Address) map[string]interface{} {
	return map[string]interface{}{
		"id":      address.ID.String(),
		"street":  address.Street,
		"city":    address.City,
		"state":   address.State,
		"zipCode": address.ZipCode,
	}
}

func MultipleUserResponse(users []*models.User) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, len(users))
	for _, a := range users {
		m = append(m, SingleUserResponse(a))
	}
	return m
}

func SinglePostResponse(post *models.Post) map[string]interface{} {
	return map[string]interface{}{
		"id":    post.ID.String(),
		"title": post.Title,
		"body":  post.Body,
	}
}

func MultiplePostResponse(posts []*models.Post) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, len(posts))
	for _, a := range posts {
		m = append(m, SinglePostResponse(a))
	}
	return m
}
