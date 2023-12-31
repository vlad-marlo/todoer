// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package model

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

func easyjson11d1a9baDecodeGithubComVladMarloTodoerInternalModel(in *jlexer.Lexer, out *UpdateTaskRequest) {
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
		case "task":
			out.Task = string(in.String())
		case "status":
			out.Status = string(in.String())
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
func easyjson11d1a9baEncodeGithubComVladMarloTodoerInternalModel(out *jwriter.Writer, in UpdateTaskRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"task\":"
		out.RawString(prefix[1:])
		out.String(string(in.Task))
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.String(string(in.Status))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UpdateTaskRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncodeGithubComVladMarloTodoerInternalModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UpdateTaskRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncodeGithubComVladMarloTodoerInternalModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UpdateTaskRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecodeGithubComVladMarloTodoerInternalModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UpdateTaskRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecodeGithubComVladMarloTodoerInternalModel(l, v)
}
func easyjson11d1a9baDecodeGithubComVladMarloTodoerInternalModel1(in *jlexer.Lexer, out *GetManyTasksRequest) {
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
func easyjson11d1a9baEncodeGithubComVladMarloTodoerInternalModel1(out *jwriter.Writer, in GetManyTasksRequest) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetManyTasksRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncodeGithubComVladMarloTodoerInternalModel1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetManyTasksRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncodeGithubComVladMarloTodoerInternalModel1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetManyTasksRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecodeGithubComVladMarloTodoerInternalModel1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetManyTasksRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecodeGithubComVladMarloTodoerInternalModel1(l, v)
}
func easyjson11d1a9baDecodeGithubComVladMarloTodoerInternalModel2(in *jlexer.Lexer, out *CreateTaskRequest) {
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
		case "task":
			out.Value = string(in.String())
		case "status":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Status).UnmarshalJSON(data))
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
func easyjson11d1a9baEncodeGithubComVladMarloTodoerInternalModel2(out *jwriter.Writer, in CreateTaskRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"task\":"
		out.RawString(prefix[1:])
		out.String(string(in.Value))
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.String(string(in.Status))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CreateTaskRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncodeGithubComVladMarloTodoerInternalModel2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CreateTaskRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncodeGithubComVladMarloTodoerInternalModel2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CreateTaskRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecodeGithubComVladMarloTodoerInternalModel2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CreateTaskRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecodeGithubComVladMarloTodoerInternalModel2(l, v)
}
