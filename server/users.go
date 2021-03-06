package main;

import (
	"net/http"
	"encoding/json"
	"strings"
	"math/rand"
)

type UserResponse struct {
	Name string
	Email string
	City string
	Image string
	Interests string
	Reviews []ReviewResponse
	Rating_Average float64
	Profession string
	Company string
	Bio string
	State string
}


func handleUsers(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		getNearbyUsers(w, req);
	}
}

func getUser(u User) UserResponse {
	return UserResponse{ u.Name, u.Email, getCity(u), getUserImage(u), strings.Join(getInterests(u), "-"), getReviews(u), getAverageRating(u), u.Profession, u.Company, u.Bio, u.State,}
}

func getNearbyUsers(w http.ResponseWriter, req *http.Request ) {
	var UserResponses []UserResponse;
	var cityId Cities;
	var users []User;
	var u User;
	city := req.Header.Get("City");
	email := req.Header.Get("Email");
	db.Where(&User{Email: email}).First(&u);
	if (len(city) > 0 && email == u.Email) {
		db.Where(&Cities{City_Name : city}).First(&cityId);
		db.Where(&User{CitiesID: cityId.ID}).Find(&users);
		users = sortUsers(u, filterSaves(u, filterFriends(u, filterRejects(u, users))));
		for _, v := range users {
			res := getUser(v);
			UserResponses = append(UserResponses, res);
		}
		r, _ := json.Marshal(UserResponses);
		w.Write(r);
	} else {
		badRequest(w, "bad get request", http.StatusBadRequest);
	}
}

func handleTarget(w http.ResponseWriter, req *http.Request) {
	var u User;
	if req.Method == http.MethodGet {
		email := req.Header.Get("Email");
		db.Where(&User{Email: email}).First(&u);
		if u.Email == email {
			ur := getUser(u);
			r, _ := json.Marshal(ur);
			w.Write(r);
		} else {
			badRequest(w, "user not found", http.StatusBadRequest);
		}
	}
}

func filterRejects(u User, users []User) []User {
	var filtered []User;
	for _,v := range users {
		if v.Email != u.Email && checkReject(u, v) {
			filtered = append(filtered, v)
		}
	}
	return filtered;
}


func sortUsers(u User, users []User) []User {
	ui := getInterests(u);
	var sortedAndShuffled []User;
	var sorted = make([][]User, len(interests));
	for _, v := range users {
		mi := getInterests(v);
		matches := countMatchingInterests(ui, mi);
		sorted[matches] = append(sorted[matches], v);
	}
	for i, a := range sorted {
		sorted[i] = shuffleUsers(a);
	}	
	for i := len(sorted) - 1; i >= 0; i-- {
		sortedAndShuffled = append(sortedAndShuffled, sorted[i]...);
	}
	return sortedAndShuffled;
}

func shuffleUsers(users []User) []User {
	shuffled := make([]User, len(users));
	perm := rand.Perm(len(users))
	for i, v := range perm {
		shuffled[v] = users[i];
	}
	return shuffled;
}

func countMatchingInterests(ui []string, mi []string) int {
	count := 0;
	uHash := make(map[string]bool, len(ui));
	for _, v := range ui {
		uHash[v] = true;
	}
	for _, m := range mi {
		if uHash[m] == true {
			count ++;
		}
	}
	return count;
}