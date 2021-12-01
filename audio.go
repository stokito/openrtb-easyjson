package openrtb

import "encoding/json"

// Audio object must be included directly in the impression object
type Audio struct {
	MIMEs          []string            `json:"mimes"`                 // Content MIME types supported.
	MinDuration    int                 `json:"minduration,omitempty"` // Minimum video ad duration in seconds
	MaxDuration    int                 `json:"maxduration,omitempty"` // Maximum video ad duration in seconds
	Protocols      []Protocol          `json:"protocols,omitempty"`   // Video bid response protocols
	StartDelay     StartDelay          `json:"startdelay,omitempty"`  // Indicates the start delay in seconds
	Sequence       int                 `json:"sequence,omitempty"`    // Default: 1
	BlockedAttrs   []CreativeAttribute `json:"battr,omitempty"`       // Blocked creative attributes
	MaxExtended    int                 `json:"maxextended,omitempty"` // Maximum extended video ad duration
	MinBitrate     int                 `json:"minbitrate,omitempty"`  // Minimum bit rate in Kbps
	MaxBitrate     int                 `json:"maxbitrate,omitempty"`  // Maximum bit rate in Kbps
	Delivery       []ContentDelivery   `json:"delivery,omitempty"`    // List of supported delivery methods
	CompanionAds   []Banner            `json:"companionad,omitempty"`
	APIs           []APIFramework      `json:"api,omitempty"`
	CompanionTypes []CompanionType     `json:"companiontype,omitempty"`
	MaxSequence    int                 `json:"maxseq,omitempty"`   // The maximumnumber of ads that canbe played in an ad pod.
	Feed           FeedType            `json:"feed,omitempty"`     // Type of audio feed.
	Stitched       int                 `json:"stitched,omitempty"` // Indicates if the ad is stitched with audio content or delivered independently
	VolumeNorm     VolumeNorm          `json:"nvol,omitempty"`     // Volume normalization mode.
	Ext            json.RawMessage     `json:"ext,omitempty"`
}
