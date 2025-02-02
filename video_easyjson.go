// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package openrtb

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson3c9ce8c3DecodeGithubComStokitoOpenrtbEasyjsonV3(in *jlexer.Lexer, out *Video) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "mimes":
			if in.IsNull() {
				in.Skip()
				out.MIMEs = nil
			} else {
				in.Delim('[')
				if out.MIMEs == nil {
					if !in.IsDelim(']') {
						out.MIMEs = make([]string, 0, 4)
					} else {
						out.MIMEs = []string{}
					}
				} else {
					out.MIMEs = (out.MIMEs)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.MIMEs = append(out.MIMEs, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "minduration":
			out.MinDuration = int(in.Int())
		case "maxduration":
			out.MaxDuration = int(in.Int())
		case "protocols":
			if in.IsNull() {
				in.Skip()
				out.Protocols = nil
			} else {
				in.Delim('[')
				if out.Protocols == nil {
					if !in.IsDelim(']') {
						out.Protocols = make([]Protocol, 0, 8)
					} else {
						out.Protocols = []Protocol{}
					}
				} else {
					out.Protocols = (out.Protocols)[:0]
				}
				for !in.IsDelim(']') {
					var v2 Protocol
					v2 = Protocol(in.Int())
					out.Protocols = append(out.Protocols, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "protocol":
			out.Protocol = Protocol(in.Int())
		case "w":
			out.Width = int(in.Int())
		case "h":
			out.Height = int(in.Int())
		case "startdelay":
			out.StartDelay = StartDelay(in.Int())
		case "linearity":
			out.Linearity = VideoLinearity(in.Int())
		case "skip":
			out.Skip = int(in.Int())
		case "skipmin":
			out.SkipMin = int(in.Int())
		case "skipafter":
			out.SkipAfter = int(in.Int())
		case "sequence":
			out.Sequence = int(in.Int())
		case "battr":
			if in.IsNull() {
				in.Skip()
				out.BlockedAttrs = nil
			} else {
				in.Delim('[')
				if out.BlockedAttrs == nil {
					if !in.IsDelim(']') {
						out.BlockedAttrs = make([]CreativeAttribute, 0, 8)
					} else {
						out.BlockedAttrs = []CreativeAttribute{}
					}
				} else {
					out.BlockedAttrs = (out.BlockedAttrs)[:0]
				}
				for !in.IsDelim(']') {
					var v3 CreativeAttribute
					v3 = CreativeAttribute(in.Int())
					out.BlockedAttrs = append(out.BlockedAttrs, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "maxextended":
			out.MaxExtended = int(in.Int())
		case "minbitrate":
			out.MinBitrate = int(in.Int())
		case "maxbitrate":
			out.MaxBitrate = int(in.Int())
		case "boxingallowed":
			if in.IsNull() {
				in.Skip()
				out.BoxingAllowed = nil
			} else {
				if out.BoxingAllowed == nil {
					out.BoxingAllowed = new(int)
				}
				*out.BoxingAllowed = int(in.Int())
			}
		case "playbackmethod":
			if in.IsNull() {
				in.Skip()
				out.PlaybackMethods = nil
			} else {
				in.Delim('[')
				if out.PlaybackMethods == nil {
					if !in.IsDelim(']') {
						out.PlaybackMethods = make([]VideoPlayback, 0, 8)
					} else {
						out.PlaybackMethods = []VideoPlayback{}
					}
				} else {
					out.PlaybackMethods = (out.PlaybackMethods)[:0]
				}
				for !in.IsDelim(']') {
					var v4 VideoPlayback
					v4 = VideoPlayback(in.Int())
					out.PlaybackMethods = append(out.PlaybackMethods, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "delivery":
			if in.IsNull() {
				in.Skip()
				out.Delivery = nil
			} else {
				in.Delim('[')
				if out.Delivery == nil {
					if !in.IsDelim(']') {
						out.Delivery = make([]ContentDelivery, 0, 8)
					} else {
						out.Delivery = []ContentDelivery{}
					}
				} else {
					out.Delivery = (out.Delivery)[:0]
				}
				for !in.IsDelim(']') {
					var v5 ContentDelivery
					v5 = ContentDelivery(in.Int())
					out.Delivery = append(out.Delivery, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "pos":
			out.Position = AdPosition(in.Int())
		case "companionad":
			if in.IsNull() {
				in.Skip()
				out.CompanionAds = nil
			} else {
				in.Delim('[')
				if out.CompanionAds == nil {
					if !in.IsDelim(']') {
						out.CompanionAds = make([]Banner, 0, 0)
					} else {
						out.CompanionAds = []Banner{}
					}
				} else {
					out.CompanionAds = (out.CompanionAds)[:0]
				}
				for !in.IsDelim(']') {
					var v6 Banner
					(v6).UnmarshalEasyJSON(in)
					out.CompanionAds = append(out.CompanionAds, v6)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "api":
			if in.IsNull() {
				in.Skip()
				out.APIs = nil
			} else {
				in.Delim('[')
				if out.APIs == nil {
					if !in.IsDelim(']') {
						out.APIs = make([]APIFramework, 0, 8)
					} else {
						out.APIs = []APIFramework{}
					}
				} else {
					out.APIs = (out.APIs)[:0]
				}
				for !in.IsDelim(']') {
					var v7 APIFramework
					v7 = APIFramework(in.Int())
					out.APIs = append(out.APIs, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "companiontype":
			if in.IsNull() {
				in.Skip()
				out.CompanionTypes = nil
			} else {
				in.Delim('[')
				if out.CompanionTypes == nil {
					if !in.IsDelim(']') {
						out.CompanionTypes = make([]CompanionType, 0, 8)
					} else {
						out.CompanionTypes = []CompanionType{}
					}
				} else {
					out.CompanionTypes = (out.CompanionTypes)[:0]
				}
				for !in.IsDelim(']') {
					var v8 CompanionType
					v8 = CompanionType(in.Int())
					out.CompanionTypes = append(out.CompanionTypes, v8)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "placement":
			out.Placement = VideoPlacement(in.Int())
		case "ext":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Ext).UnmarshalJSON(data))
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson3c9ce8c3EncodeGithubComStokitoOpenrtbEasyjsonV3(out *jwriter.Writer, in Video) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.MIMEs) != 0 {
		const prefix string = ",\"mimes\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v9, v10 := range in.MIMEs {
				if v9 > 0 {
					out.RawByte(',')
				}
				out.String(string(v10))
			}
			out.RawByte(']')
		}
	}
	if in.MinDuration != 0 {
		const prefix string = ",\"minduration\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.MinDuration))
	}
	if in.MaxDuration != 0 {
		const prefix string = ",\"maxduration\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.MaxDuration))
	}
	if len(in.Protocols) != 0 {
		const prefix string = ",\"protocols\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v11, v12 := range in.Protocols {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.Int(int(v12))
			}
			out.RawByte(']')
		}
	}
	if in.Protocol != 0 {
		const prefix string = ",\"protocol\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Protocol))
	}
	if in.Width != 0 {
		const prefix string = ",\"w\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Width))
	}
	if in.Height != 0 {
		const prefix string = ",\"h\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Height))
	}
	if in.StartDelay != 0 {
		const prefix string = ",\"startdelay\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.StartDelay))
	}
	if in.Linearity != 0 {
		const prefix string = ",\"linearity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Linearity))
	}
	if in.Skip != 0 {
		const prefix string = ",\"skip\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Skip))
	}
	if in.SkipMin != 0 {
		const prefix string = ",\"skipmin\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.SkipMin))
	}
	if in.SkipAfter != 0 {
		const prefix string = ",\"skipafter\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.SkipAfter))
	}
	if in.Sequence != 0 {
		const prefix string = ",\"sequence\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Sequence))
	}
	if len(in.BlockedAttrs) != 0 {
		const prefix string = ",\"battr\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v13, v14 := range in.BlockedAttrs {
				if v13 > 0 {
					out.RawByte(',')
				}
				out.Int(int(v14))
			}
			out.RawByte(']')
		}
	}
	if in.MaxExtended != 0 {
		const prefix string = ",\"maxextended\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.MaxExtended))
	}
	if in.MinBitrate != 0 {
		const prefix string = ",\"minbitrate\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.MinBitrate))
	}
	if in.MaxBitrate != 0 {
		const prefix string = ",\"maxbitrate\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.MaxBitrate))
	}
	if in.BoxingAllowed != nil {
		const prefix string = ",\"boxingallowed\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(*in.BoxingAllowed))
	}
	if len(in.PlaybackMethods) != 0 {
		const prefix string = ",\"playbackmethod\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v15, v16 := range in.PlaybackMethods {
				if v15 > 0 {
					out.RawByte(',')
				}
				out.Int(int(v16))
			}
			out.RawByte(']')
		}
	}
	if len(in.Delivery) != 0 {
		const prefix string = ",\"delivery\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v17, v18 := range in.Delivery {
				if v17 > 0 {
					out.RawByte(',')
				}
				out.Int(int(v18))
			}
			out.RawByte(']')
		}
	}
	if in.Position != 0 {
		const prefix string = ",\"pos\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Position))
	}
	if len(in.CompanionAds) != 0 {
		const prefix string = ",\"companionad\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v19, v20 := range in.CompanionAds {
				if v19 > 0 {
					out.RawByte(',')
				}
				(v20).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if len(in.APIs) != 0 {
		const prefix string = ",\"api\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v21, v22 := range in.APIs {
				if v21 > 0 {
					out.RawByte(',')
				}
				out.Int(int(v22))
			}
			out.RawByte(']')
		}
	}
	if len(in.CompanionTypes) != 0 {
		const prefix string = ",\"companiontype\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v23, v24 := range in.CompanionTypes {
				if v23 > 0 {
					out.RawByte(',')
				}
				out.Int(int(v24))
			}
			out.RawByte(']')
		}
	}
	if in.Placement != 0 {
		const prefix string = ",\"placement\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Placement))
	}
	if len(in.Ext) != 0 {
		const prefix string = ",\"ext\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Ext).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Video) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3c9ce8c3EncodeGithubComStokitoOpenrtbEasyjsonV3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Video) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3c9ce8c3EncodeGithubComStokitoOpenrtbEasyjsonV3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Video) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3c9ce8c3DecodeGithubComStokitoOpenrtbEasyjsonV3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Video) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3c9ce8c3DecodeGithubComStokitoOpenrtbEasyjsonV3(l, v)
}
