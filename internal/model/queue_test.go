package model

// var testdata1Str string = `
// {
// 	"from": "from-test",
// 	"to": "test123+abcde@gmail.com",
// 	"kind": "email",
// 	"title": "title test",
// 	"body": "abcde\nfghij\n12345"
// }
// `
// var testdata1Ans Queue = Queue{
// 	From:  "from-test",
// 	To:    "test123+abcde@gmail.com",
// 	Kind:  Queuekind(QueueKindEmail),
// 	Title: "title test",
// 	Body: `abcde
// fghij
// 12345`,
// }

// var testdata2Str string = `
// {
// 	"to": "test123+abcde@gmail.com",
// 	"kind": "email",
// 	"body": "abcde\nfghij\n12345"
// }
// `

// var testdata2Ans Queue = Queue{
// 	To:   "test123+abcde@gmail.com",
// 	Kind: Queuekind(QueueKindEmail),
// 	Body: `abcde
// fghij
// 12345`,
// }

// var testdata3Str string = `
// {
// 	"to": "test123+abcde@gmail.com   ,
// 	"kind": "email",
// 	"body": "abcde\nfghij\n12345"
// }
// `

// func TestQueue_UnmarshalJSON(t *testing.T) {
// 	type fields struct {
// 		From  string
// 		To    string
// 		Kind  Queuekind
// 		Title string
// 		Body  string
// 	}
// 	type args struct {
// 		b []byte
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 		want    *Queue
// 	}{
// 		{
// 			name:    "ok",
// 			fields:  fields{},
// 			args:    args{b: []byte(testdata1Str)},
// 			wantErr: false,
// 			want:    &testdata1Ans,
// 		},
// 		{
// 			name:    "ok partially",
// 			fields:  fields{},
// 			args:    args{b: []byte(testdata2Str)},
// 			wantErr: false,
// 			want:    &testdata2Ans,
// 		},
// 		{
// 			name:    "syntax error",
// 			fields:  fields{},
// 			args:    args{b: []byte(testdata3Str)},
// 			wantErr: true,
// 			want:    nil,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			q := &Queue{
// 				From:  tt.fields.From,
// 				To:    tt.fields.To,
// 				Kind:  tt.fields.Kind,
// 				Title: tt.fields.Title,
// 				Body:  tt.fields.Body,
// 			}
// 			if err := q.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
// 				t.Errorf("Queue.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 			if !reflect.DeepEqual(q, tt.want) {
// 				if !tt.wantErr {
// 					t.Errorf("Queue.UnmarshalJSON() = %+v, want %+v", q, tt.want)
// 				}
// 			}
// 		})
// 	}
// }
