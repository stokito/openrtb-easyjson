package openrtb

import (
	"encoding/json"
)

// Bid object contains bid information.
// ID, ImpID and Price are required; all other optional.
// If the bidder wins the impression, the exchange calls notice URL (nurl)
// a) to inform the bidder of the win;
// b) to convey certain information using substitution macros.
// Adomain can be used to check advertiser block list compliance.
// Cid can be used to block ads that were previously identified as inappropriate.
// Substitution macros may allow a bidder to use a static notice URL for all of its bids.
type Bid struct {
	ID          string              `json:"id"`
	ImpID       string              `json:"impid"`                    // Required string ID of the impression object to which this bid applies.
	Price       float64             `json:"price"`                    // Bid price in CPM. Suggests using integer math for accounting to avoid rounding errors.
	AdID        string              `json:"adid,omitempty"`           // References the ad to be served if the bid wins.
	NoticeURL   string              `json:"nurl,omitempty"`           // Win notice URL.
	BillingURL  string              `json:"burl,omitempty"`           // Billing notice URL.
	LossURL     string              `json:"lurl,omitempty"`           // Loss notice URL.
	AdMarkup    string              `json:"adm,omitempty"`            // Actual ad markup. XHTML if a response to a banner object, or VAST XML if a response to a video object.
	AdvDomains  []string            `json:"adomain,omitempty"`        // Advertiser’s primary or top-level domain for advertiser checking; or multiple if imp rotating.
	Bundle      string              `json:"bundle,omitempty"`         // A platform-specific application identifier intended to be unique to the app and independent of the exchange.
	ImageURL    string              `json:"iurl,omitempty"`           // Sample image URL.
	CampaignID  string              `json:"cid,omitempty"`            // Campaign ID that appears with the Ad markup.
	CreativeID  string              `json:"crid,omitempty"`           // Creative ID for reporting content issues or defects. This could also be used as a reference to a creative ID that is posted with an exchange.
	Tactic      string              `json:"tactic,omitempty"`         // Tactic ID to enable buyers to label bids for reporting to the exchange the tactic through which their bid was submitted.
	Categories  []ContentCategory   `json:"cat,omitempty"`            // IAB content categories of the creative. Refer to List 5.1
	Attrs       []CreativeAttribute `json:"attr,omitempty"`           // Array of creative attributes.
	API         APIFramework        `json:"api,omitempty"`            // API required by the markup if applicable
	Protocol    Protocol            `json:"protocol,omitempty"`       // Video response protocol of the markup if applicable
	MediaRating IQGRating           `json:"qagmediarating,omitempty"` // Creative media rating per IQG guidelines.
	Language    string              `json:"language,omitempty"`       // Language of the creative using ISO-639-1-alpha-2.
	DealID      string              `json:"dealid,omitempty"`         // DealID extension of private marketplace deals
	Width       int                 `json:"w,omitempty"`              // Width of the ad in pixels.
	Height      int                 `json:"h,omitempty"`              // Height of the ad in pixels.
	WidthRatio  int                 `json:"wratio,omitempty"`         // Relative width of the creative when expressing size as a ratio.
	HeightRatio int                 `json:"hratio,omitempty"`         // Relative height of the creative when expressing size as a ratio.
	Exp         int                 `json:"exp,omitempty"`            // Advisory as to the number of seconds the bidder is willing to wait between the auction and the actual impression.
	Ext         json.RawMessage     `json:"ext,omitempty"`
}
