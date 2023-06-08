// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package lib

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	big "math/big"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson6601e8cdDecodeFunction(in *jlexer.Lexer, out *computeEvent) {
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
		case "JobId":
			if in.IsNull() {
				in.Skip()
				out.JobId = nil
			} else {
				if out.JobId == nil {
					out.JobId = new(big.Int)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.JobId).UnmarshalJSON(data))
				}
			}
		case "Gender":
			out.Gender = string(in.String())
		case "Sender":
			out.Sender = string(in.String())
		case "Count":
			if in.IsNull() {
				in.Skip()
				out.Count = nil
			} else {
				if out.Count == nil {
					out.Count = new(big.Int)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.Count).UnmarshalJSON(data))
				}
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
func easyjson6601e8cdEncodeFunction(out *jwriter.Writer, in computeEvent) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"JobId\":"
		out.RawString(prefix[1:])
		if in.JobId == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.JobId).MarshalJSON())
		}
	}
	{
		const prefix string = ",\"Gender\":"
		out.RawString(prefix)
		out.String(string(in.Gender))
	}
	{
		const prefix string = ",\"Sender\":"
		out.RawString(prefix)
		out.String(string(in.Sender))
	}
	{
		const prefix string = ",\"Count\":"
		out.RawString(prefix)
		if in.Count == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.Count).MarshalJSON())
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v computeEvent) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeFunction(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v computeEvent) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeFunction(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *computeEvent) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeFunction(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *computeEvent) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeFunction(l, v)
}
func easyjson6601e8cdDecodeFunction1(in *jlexer.Lexer, out *avatar) {
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
		case "Cid":
			out.Cid = string(in.String())
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
func easyjson6601e8cdEncodeFunction1(out *jwriter.Writer, in avatar) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Cid\":"
		out.RawString(prefix[1:])
		out.String(string(in.Cid))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v avatar) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeFunction1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v avatar) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeFunction1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *avatar) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeFunction1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *avatar) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeFunction1(l, v)
}