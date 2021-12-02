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

func easyjson10eb023eDecodeGithubComStokitoOpenrtbEasyjsonV3(in *jlexer.Lexer, out *BidResponse) {
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
		case "seatbid":
			if in.IsNull() {
				in.Skip()
				out.SeatBids = nil
			} else {
				in.Delim('[')
				if out.SeatBids == nil {
					if !in.IsDelim(']') {
						out.SeatBids = make([]SeatBid, 0, 0)
					} else {
						out.SeatBids = []SeatBid{}
					}
				} else {
					out.SeatBids = (out.SeatBids)[:0]
				}
				for !in.IsDelim(']') {
					var v1 SeatBid
					easyjson10eb023eDecodeGithubComStokitoOpenrtbEasyjsonV31(in, &v1)
					out.SeatBids = append(out.SeatBids, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "bidid":
			out.BidID = string(in.String())
		case "cur":
			out.Currency = string(in.String())
		case "customdata":
			out.CustomData = string(in.String())
		case "nbr":
			out.NBR = NBR(in.Int())
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
func easyjson10eb023eEncodeGithubComStokitoOpenrtbEasyjsonV3(out *jwriter.Writer, in BidResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"seatbid\":"
		out.RawString(prefix)
		if in.SeatBids == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.SeatBids {
				if v2 > 0 {
					out.RawByte(',')
				}
				easyjson10eb023eEncodeGithubComStokitoOpenrtbEasyjsonV31(out, v3)
			}
			out.RawByte(']')
		}
	}
	if in.BidID != "" {
		const prefix string = ",\"bidid\":"
		out.RawString(prefix)
		out.String(string(in.BidID))
	}
	if in.Currency != "" {
		const prefix string = ",\"cur\":"
		out.RawString(prefix)
		out.String(string(in.Currency))
	}
	if in.CustomData != "" {
		const prefix string = ",\"customdata\":"
		out.RawString(prefix)
		out.String(string(in.CustomData))
	}
	if in.NBR != 0 {
		const prefix string = ",\"nbr\":"
		out.RawString(prefix)
		out.Int(int(in.NBR))
	}
	if len(in.Ext) != 0 {
		const prefix string = ",\"ext\":"
		out.RawString(prefix)
		out.Raw((in.Ext).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BidResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson10eb023eEncodeGithubComStokitoOpenrtbEasyjsonV3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BidResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson10eb023eEncodeGithubComStokitoOpenrtbEasyjsonV3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BidResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson10eb023eDecodeGithubComStokitoOpenrtbEasyjsonV3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BidResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson10eb023eDecodeGithubComStokitoOpenrtbEasyjsonV3(l, v)
}
func easyjson10eb023eDecodeGithubComStokitoOpenrtbEasyjsonV31(in *jlexer.Lexer, out *SeatBid) {
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
		case "bid":
			if in.IsNull() {
				in.Skip()
				out.Bids = nil
			} else {
				in.Delim('[')
				if out.Bids == nil {
					if !in.IsDelim(']') {
						out.Bids = make([]Bid, 0, 0)
					} else {
						out.Bids = []Bid{}
					}
				} else {
					out.Bids = (out.Bids)[:0]
				}
				for !in.IsDelim(']') {
					var v4 Bid
					(v4).UnmarshalEasyJSON(in)
					out.Bids = append(out.Bids, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "seat":
			out.Seat = string(in.String())
		case "group":
			out.Group = int(in.Int())
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
func easyjson10eb023eEncodeGithubComStokitoOpenrtbEasyjsonV31(out *jwriter.Writer, in SeatBid) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"bid\":"
		out.RawString(prefix[1:])
		if in.Bids == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Bids {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if in.Seat != "" {
		const prefix string = ",\"seat\":"
		out.RawString(prefix)
		out.String(string(in.Seat))
	}
	if in.Group != 0 {
		const prefix string = ",\"group\":"
		out.RawString(prefix)
		out.Int(int(in.Group))
	}
	if len(in.Ext) != 0 {
		const prefix string = ",\"ext\":"
		out.RawString(prefix)
		out.Raw((in.Ext).MarshalJSON())
	}
	out.RawByte('}')
}
