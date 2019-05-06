package models

type User struct {
	Handle                  string `form:"username"`
	Email                   string `json:"email"`
	VkId                    string `json:"vkid"`
	OpenId                  string `json:"openid"`
	FirstName               string `json:"firstname"`
	LastName                string `json:"lastname"`
	Country                 string `json:"country"`
	City                    string `json:"city"`
	Organization            string `json:"organization"`
	Contribution            int    `json:"contribution"`
	Rank                    string `json:"rank"`
	Rating                  int    `json:"rating"`
	MaxRank                 string `json:"maxrank"`
	MaxRating               int    `json:"maxrating"`
	LastOnlineTimeSeconds   int    `json:"lastonlinetimeseconds"`
	RegistrationTimeSeconds int    `json:"registrationtimeseconds"`
	FriendOfCount           int    `json:"friendofcount"`
	Avatar                  string `json:"avatar"`
	TitlePhoto              string `json:"titlephoto"`
}
