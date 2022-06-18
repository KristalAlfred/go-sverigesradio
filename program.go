package sverigesradio

type ProgramService service

// Program represents
type Program struct {
	Description     string `json:"description"`
	Programcategory struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"programcategory,omitempty"`
	Email                    string `json:"email"`
	Phone                    string `json:"phone"`
	Programurl               string `json:"programurl"`
	Programslug              string `json:"programslug"`
	Programimage             string `json:"programimage"`
	Programimagetemplate     string `json:"programimagetemplate"`
	Programimagewide         string `json:"programimagewide"`
	Programimagetemplatewide string `json:"programimagetemplatewide"`
	Socialimage              string `json:"socialimage"`
	Socialimagetemplate      string `json:"socialimagetemplate"`
	Socialmediaplatforms     []struct {
		Platform    string `json:"platform"`
		Platformurl string `json:"platformurl"`
	} `json:"socialmediaplatforms"`
	Channel struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"channel"`
	Archived          bool   `json:"archived"`
	Hasondemand       bool   `json:"hasondemand"`
	Haspod            bool   `json:"haspod"`
	Responsibleeditor string `json:"responsibleeditor"`
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Broadcastinfo     string `json:"broadcastinfo,omitempty"`
	Payoff            string `json:"payoff,omitempty"`
}
