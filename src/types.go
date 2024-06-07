package src

type Rickroll struct {
	ID              uint   `gorm:"primaryKey"`
	URLEnding       string `gorm:"unique"`
	Clicks          uint
	SiteTitle       string
	SiteName        string
	ImgLink         string
	SiteDescription string
}

type RickrollRequest struct {
	URLEnding       string `json:"urlEnding"`
	SiteTitle       string `json:"siteTitle"`
	SiteName        string `json:"siteName"`
	ImgLink         string `json:"imgLink"`
	SiteDescription string `json:"siteDescription"`
}

type AdminPasswordRequest struct {
	AdminPassword string `json:"adminPassword"`
}

type DeleteRequest struct {
	ID            uint   `json:"id"`
	AdminPassword string `json:"adminPassword"`
}

// used for returning all links in the admin panel
type RickrollInfo struct {
	ID        uint   `json:"id"`
	URLEnding string `json:"urlEnding"`
	Clicks    uint   `json:"clicks"`
}
