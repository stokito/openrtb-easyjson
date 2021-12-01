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

func easyjsonDedb5be1DecodeGithubComStokitoOpenrtbEasyjson(in *jlexer.Lexer, out *PMP) {
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
		case "private_auction":
			out.Private = int(in.Int())
		case "deals":
			if in.IsNull() {
				in.Skip()
				out.Deals = nil
			} else {
				in.Delim('[')
				if out.Deals == nil {
					if !in.IsDelim(']') {
						out.Deals = make([]Deal, 0, 0)
					} else {
						out.Deals = []Deal{}
					}
				} else {
					out.Deals = (out.Deals)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Deal
					(v1).UnmarshalEasyJSON(in)
					out.Deals = append(out.Deals, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonDedb5be1EncodeGithubComStokitoOpenrtbEasyjson(out *jwriter.Writer, in PMP) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Private != 0 {
		const prefix string = ",\"private_auction\":"
		first = false
		out.RawString(prefix[1:])
		out.Int(int(in.Private))
	}
	if len(in.Deals) != 0 {
		const prefix string = ",\"deals\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v2, v3 := range in.Deals {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
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
func (v PMP) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDedb5be1EncodeGithubComStokitoOpenrtbEasyjson(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PMP) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDedb5be1EncodeGithubComStokitoOpenrtbEasyjson(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PMP) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDedb5be1DecodeGithubComStokitoOpenrtbEasyjson(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PMP) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDedb5be1DecodeGithubComStokitoOpenrtbEasyjson(l, v)
}
func easyjsonDedb5be1DecodeGithubComStokitoOpenrtbEasyjson1(in *jlexer.Lexer, out *Deal) {
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
		case "id":
			out.ID = string(in.String())
		case "bidfloor":
			out.BidFloor = float64(in.Float64())
		case "bidfloorcur":
			out.BidFloorCurrency = string(in.String())
		case "wseat":
			if in.IsNull() {
				in.Skip()
				out.Seats = nil
			} else {
				in.Delim('[')
				if out.Seats == nil {
					if !in.IsDelim(']') {
						out.Seats = make([]string, 0, 4)
					} else {
						out.Seats = []string{}
					}
				} else {
					out.Seats = (out.Seats)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.Seats = append(out.Seats, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "wadomain":
			if in.IsNull() {
				in.Skip()
				out.AdvDomains = nil
			} else {
				in.Delim('[')
				if out.AdvDomains == nil {
					if !in.IsDelim(']') {
						out.AdvDomains = make([]string, 0, 4)
					} else {
						out.AdvDomains = []string{}
					}
				} else {
					out.AdvDomains = (out.AdvDomains)[:0]
				}
				for !in.IsDelim(']') {
					var v5 string
					v5 = string(in.String())
					out.AdvDomains = append(out.AdvDomains, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "at":
			out.AuctionType = int(in.Int())
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
func easyjsonDedb5be1EncodeGithubComStokitoOpenrtbEasyjson1(out *jwriter.Writer, in Deal) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != "" {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	if in.BidFloor != 0 {
		const prefix string = ",\"bidfloor\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.BidFloor))
	}
	if in.BidFloorCurrency != "" {
		const prefix string = ",\"bidfloorcur\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.BidFloorCurrency))
	}
	if len(in.Seats) != 0 {
		const prefix string = ",\"wseat\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v6, v7 := range in.Seats {
				if v6 > 0 {
					out.RawByte(',')
				}
				out.String(string(v7))
			}
			out.RawByte(']')
		}
	}
	if len(in.AdvDomains) != 0 {
		const prefix string = ",\"wadomain\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v8, v9 := range in.AdvDomains {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.String(string(v9))
			}
			out.RawByte(']')
		}
	}
	if in.AuctionType != 0 {
		const prefix string = ",\"at\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.AuctionType))
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
func (v Deal) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDedb5be1EncodeGithubComStokitoOpenrtbEasyjson1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Deal) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDedb5be1EncodeGithubComStokitoOpenrtbEasyjson1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Deal) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDedb5be1DecodeGithubComStokitoOpenrtbEasyjson1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Deal) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDedb5be1DecodeGithubComStokitoOpenrtbEasyjson1(l, v)
}
