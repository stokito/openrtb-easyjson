// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package response

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

func easyjson3b94576aDecodeGithubComStokitoOpenrtbEasyjsonNativeResponse(in *jlexer.Lexer, out *Asset) {
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
			out.ID = int(in.Int())
		case "required":
			out.Required = int(in.Int())
		case "title":
			if in.IsNull() {
				in.Skip()
				out.Title = nil
			} else {
				if out.Title == nil {
					out.Title = new(Title)
				}
				easyjson3b94576aDecodeGithubComStokitoOpenrtbEasyjsonNativeResponse1(in, out.Title)
			}
		case "img":
			if in.IsNull() {
				in.Skip()
				out.Image = nil
			} else {
				if out.Image == nil {
					out.Image = new(Image)
				}
				easyjson3b94576aDecodeGithubComStokitoOpenrtbEasyjsonNativeResponse2(in, out.Image)
			}
		case "video":
			if in.IsNull() {
				in.Skip()
				out.Video = nil
			} else {
				if out.Video == nil {
					out.Video = new(Video)
				}
				easyjson3b94576aDecodeGithubComStokitoOpenrtbEasyjsonNativeResponse3(in, out.Video)
			}
		case "data":
			if in.IsNull() {
				in.Skip()
				out.Data = nil
			} else {
				if out.Data == nil {
					out.Data = new(Data)
				}
				easyjson3b94576aDecodeGithubComStokitoOpenrtbEasyjsonNativeResponse4(in, out.Data)
			}
		case "link":
			if in.IsNull() {
				in.Skip()
				out.Link = nil
			} else {
				if out.Link == nil {
					out.Link = new(Link)
				}
				easyjson3b94576aDecodeGithubComStokitoOpenrtbEasyjsonNativeResponse5(in, out.Link)
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
func easyjson3b94576aEncodeGithubComStokitoOpenrtbEasyjsonNativeResponse(out *jwriter.Writer, in Asset) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	if in.Required != 0 {
		const prefix string = ",\"required\":"
		out.RawString(prefix)
		out.Int(int(in.Required))
	}
	if in.Title != nil {
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		easyjson3b94576aEncodeGithubComStokitoOpenrtbEasyjsonNativeResponse1(out, *in.Title)
	}
	if in.Image != nil {
		const prefix string = ",\"img\":"
		out.RawString(prefix)
		easyjson3b94576aEncodeGithubComStokitoOpenrtbEasyjsonNativeResponse2(out, *in.Image)
	}
	if in.Video != nil {
		const prefix string = ",\"video\":"
		out.RawString(prefix)
		easyjson3b94576aEncodeGithubComStokitoOpenrtbEasyjsonNativeResponse3(out, *in.Video)
	}
	if in.Data != nil {
		const prefix string = ",\"data\":"
		out.RawString(prefix)
		easyjson3b94576aEncodeGithubComStokitoOpenrtbEasyjsonNativeResponse4(out, *in.Data)
	}
	if in.Link != nil {
		const prefix string = ",\"link\":"
		out.RawString(prefix)
		easyjson3b94576aEncodeGithubComStokitoOpenrtbEasyjsonNativeResponse5(out, *in.Link)
	}
	if len(in.Ext) != 0 {
		const prefix string = ",\"ext\":"
		out.RawString(prefix)
		out.Raw((in.Ext).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Asset) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3b94576aEncodeGithubComStokitoOpenrtbEasyjsonNativeResponse(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Asset) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3b94576aEncodeGithubComStokitoOpenrtbEasyjsonNativeResponse(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Asset) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3b94576aDecodeGithubComStokitoOpenrtbEasyjsonNativeResponse(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Asset) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3b94576aDecodeGithubComStokitoOpenrtbEasyjsonNativeResponse(l, v)
}
func easyjson3b94576aDecodeGithubComStokitoOpenrtbEasyjsonNativeResponse5(in *jlexer.Lexer, out *Link) {
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
		case "url":
			out.URL = string(in.String())
		case "clicktrackers":
			if in.IsNull() {
				in.Skip()
				out.ClickTrackers = nil
			} else {
				in.Delim('[')
				if out.ClickTrackers == nil {
					if !in.IsDelim(']') {
						out.ClickTrackers = make([]string, 0, 4)
					} else {
						out.ClickTrackers = []string{}
					}
				} else {
					out.ClickTrackers = (out.ClickTrackers)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.ClickTrackers = append(out.ClickTrackers, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "fallback":
			out.FallbackURL = string(in.String())
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
func easyjson3b94576aEncodeGithubComStokitoOpenrtbEasyjsonNativeResponse5(out *jwriter.Writer, in Link) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"url\":"
		out.RawString(prefix[1:])
		out.String(string(in.URL))
	}
	if len(in.ClickTrackers) != 0 {
		const prefix string = ",\"clicktrackers\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v2, v3 := range in.ClickTrackers {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	if in.FallbackURL != "" {
		const prefix string = ",\"fallback\":"
		out.RawString(prefix)
		out.String(string(in.FallbackURL))
	}
	if len(in.Ext) != 0 {
		const prefix string = ",\"ext\":"
		out.RawString(prefix)
		out.Raw((in.Ext).MarshalJSON())
	}
	out.RawByte('}')
}
func easyjson3b94576aDecodeGithubComStokitoOpenrtbEasyjsonNativeResponse4(in *jlexer.Lexer, out *Data) {
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
		case "label":
			out.Label = string(in.String())
		case "value":
			out.Value = string(in.String())
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
func easyjson3b94576aEncodeGithubComStokitoOpenrtbEasyjsonNativeResponse4(out *jwriter.Writer, in Data) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Label != "" {
		const prefix string = ",\"label\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Label))
	}
	{
		const prefix string = ",\"value\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Value))
	}
	if len(in.Ext) != 0 {
		const prefix string = ",\"ext\":"
		out.RawString(prefix)
		out.Raw((in.Ext).MarshalJSON())
	}
	out.RawByte('}')
}
func easyjson3b94576aDecodeGithubComStokitoOpenrtbEasyjsonNativeResponse3(in *jlexer.Lexer, out *Video) {
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
		case "vasttag":
			out.VASTTag = string(in.String())
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
func easyjson3b94576aEncodeGithubComStokitoOpenrtbEasyjsonNativeResponse3(out *jwriter.Writer, in Video) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"vasttag\":"
		out.RawString(prefix[1:])
		out.String(string(in.VASTTag))
	}
	out.RawByte('}')
}
func easyjson3b94576aDecodeGithubComStokitoOpenrtbEasyjsonNativeResponse2(in *jlexer.Lexer, out *Image) {
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
		case "url":
			out.URL = string(in.String())
		case "w":
			out.Width = int(in.Int())
		case "h":
			out.Height = int(in.Int())
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
func easyjson3b94576aEncodeGithubComStokitoOpenrtbEasyjsonNativeResponse2(out *jwriter.Writer, in Image) {
	out.RawByte('{')
	first := true
	_ = first
	if in.URL != "" {
		const prefix string = ",\"url\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.URL))
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
func easyjson3b94576aDecodeGithubComStokitoOpenrtbEasyjsonNativeResponse1(in *jlexer.Lexer, out *Title) {
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
		case "text":
			out.Text = string(in.String())
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
func easyjson3b94576aEncodeGithubComStokitoOpenrtbEasyjsonNativeResponse1(out *jwriter.Writer, in Title) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix[1:])
		out.String(string(in.Text))
	}
	if len(in.Ext) != 0 {
		const prefix string = ",\"ext\":"
		out.RawString(prefix)
		out.Raw((in.Ext).MarshalJSON())
	}
	out.RawByte('}')
}
