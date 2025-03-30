package main

import "fmt"

type User struct {
	// ID       int64
	Username string
	Email    string
	Password string
}

var usernames = []string{
	"alice", "bob", "charlie", "dave", "eve", "frank", "grace", "heidi",
	"ivan", "judy", "karl", "laura", "mallory", "nina", "oscar", "peggy",
	"quinn", "rachel", "steve", "trent", "ursula", "victor", "wendy", "xander",
	"yvonne", "zack", "amber", "brian", "carol", "doug", "eric", "fiona",
	"george", "hannah", "ian", "jessica", "kevin", "lisa", "mike", "natalie",
	"oliver", "peter", "queen", "ron", "susan", "tim", "uma", "vicky",
	"walter", "xenia", "yasmin", "zoe",
}

func generateUsers(num int) []*User {
	// Create a slice of user pointers
	users := make([]*User, num)

	for i := 0; i < num; i++ {
		users[i] = &User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			Password: "123456",
		}
	}
	return users
}

func main() {
	// Generate 10 users
	users := generateUsers(10)

	// Print user details
	for _, user := range users {
		fmt.Printf("Username: %s, Email: %s, Password: %s\n", user.Username, user.Email, user.Password)
	}
}
