package openrtb

import (
	"encoding/json"
)

// Video object must be included directly in the impression object if the impression offered
// for auction is an in-stream video ad opportunity.
type Video struct {
	MIMEs           []string            `json:"mimes,omitempty"`          // Content MIME types supported.
	MinDuration     int                 `json:"minduration,omitempty"`    // Minimum video ad duration in seconds
	MaxDuration     int                 `json:"maxduration,omitempty"`    // Maximum video ad duration in seconds
	Protocols       []Protocol          `json:"protocols,omitempty"`      // Video bid response protocols
	Protocol        Protocol            `json:"protocol,omitempty"`       // Video bid response protocols DEPRECATED
	Width           int                 `json:"w,omitempty"`              // Width of the player in pixels
	Height          int                 `json:"h,omitempty"`              // Height of the player in pixels
	StartDelay      StartDelay          `json:"startdelay,omitempty"`     // Indicates the start delay in seconds
	Linearity       VideoLinearity      `json:"linearity,omitempty"`      // Indicates whether the ad impression is linear or non-linear
	Skip            int                 `json:"skip,omitempty"`           // Indicates if the player will allow the video to be skipped, where 0 = no, 1 = yes.
	SkipMin         int                 `json:"skipmin,omitempty"`        // Videos of total duration greater than this number of seconds can be skippable
	SkipAfter       int                 `json:"skipafter,omitempty"`      // Number of seconds a video must play before skipping is enabled
	Sequence        int                 `json:"sequence,omitempty"`       // Default: 1
	BlockedAttrs    []CreativeAttribute `json:"battr,omitempty"`          // Blocked creative attributes
	MaxExtended     int                 `json:"maxextended,omitempty"`    // Maximum extended video ad duration
	MinBitrate      int                 `json:"minbitrate,omitempty"`     // Minimum bit rate in Kbps
	MaxBitrate      int                 `json:"maxbitrate,omitempty"`     // Maximum bit rate in Kbps
	BoxingAllowed   *int                `json:"boxingallowed,omitempty"`  // If exchange publisher has rules preventing letter boxing
	PlaybackMethods []VideoPlayback     `json:"playbackmethod,omitempty"` // List of allowed playback methods
	Delivery        []ContentDelivery   `json:"delivery,omitempty"`       // List of supported delivery methods
	Position        AdPosition          `json:"pos,omitempty"`            // Ad Position
	CompanionAds    []Banner            `json:"companionad,omitempty"`
	APIs            []APIFramework      `json:"api,omitempty"` // List of supported API frameworks
	CompanionTypes  []CompanionType     `json:"companiontype,omitempty"`
	Placement       VideoPlacement      `json:"placement,omitempty"` // Video placement type
	Ext             json.RawMessage     `json:"ext,omitempty"`
}

// GetBoxingAllowed returns the boxing-allowed indicator
func (v *Video) GetBoxingAllowed() int {
	if v.BoxingAllowed != nil {
		return *v.BoxingAllowed
	}
	return 1
}
