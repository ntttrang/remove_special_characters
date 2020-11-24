package handlerString

import "testing"

// TestRemoveSpecialCharacters Remove special characters
// Generate non-printable characters: https://onlineasciitools.com/generate-ascii-characters
// See non-printerable characters: https://www.soscisurvey.de/tools/view-chars.php
func TestRemoveSpecialCharacters(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1. Correct Input: Input does not contain special characters (non-printable and printable)",
			args: args{input: "MUANGPRUAN INTANIT"},
			want: "MUANGPRUAN INTANIT",
		},
		{
			name:"2. Correct Input: Input does not contain special characters (non-printable and printable)",
			args: args{input: "TUNLAYA-ANUKIT RATTANAPENG"},
			want: "TUNLAYA-ANUKIT RATTANAPENG",
		},
		{
			name:"3. Input contains `#`",
			args: args{input: "MUANG#PRUAN INTAN##IT###"},
			want: "MUANGPRUAN INTANIT",
		},
		{
			name:"4. Input contains `~`",
			args: args{input: "MUANG~~PRUAN INTAN~IT~~~~~~"},
			want: "MUANGPRUAN INTANIT",
		},
		{
			name:"5. Input contains `|`",
			args: args{input: "MUANG|PRUAN INTAN||IT|||||||"},
			want: "MUANGPRUAN INTANIT",
		},
		{
			name:"6. Input contains `:`",
			args: args{input: "MUANG:PRUAN INTAN::IT::::"},
			want: "MUANGPRUAN INTANIT",
		},
		{
			name:"7. Input contains `;`",
			args: args{input: "MUANG;PRUAN INTAN;;IT;;;;;"},
			want: "MUANGPRUAN INTANIT",
		},
		{
			name:"8. Input contains `@`",
			args: args{input: "MUANG@PRUAN INTAN@@@IT@@@@@"},
			want: "MUANGPRUAN INTANIT",
		},
		{
			name:"9. Input contains `^`",
			args: args{input: "MUANG^PRUAN INTAN^^^IT^^^^^^"},
			want: "MUANGPRUAN INTANIT",
		},
		{
			name:"10. Input contains all special characters(printable)",
			args: args{input: "MUANG#^|:;PRUAN INTAN@@IT#;~"},
			want: "MUANGPRUAN INTANIT",
		},
		{
			name:"11. Input contains non-printable characters (ASCII: 0-31)",
			args: args{input: "MUANGPRUAN\u0000\u0001\u0002\u0003\u0004\u0005\u0006\a\b\t\n INTANIT\v\f\n\u000E\u000F\u0010\u0011\u0012\u0013\u0014\u0015\u0016\u0017\u0018\u0019\u001A\u001B\u001C\u001D\u001E\u001F"},
			want: "MUANGPRUAN INTANIT",
		},
		{
			name:"12. Input contains non-printable characters (ASCII: 127-160)",
			args: args{input: "\u007F\u0080\u0081\u0082\u0083\u0084MUANGPRUAN \u0085\u0086\u0087\u0088\u0089\u008A\u008B\u008C\u008D\u008E\u008F\u0090\u0091\u0092\u0093\u0094\u0095\u0096\u0097\u0098\u0099INTANIT\u009A\u009B\u009C\u009D\u009E\u009F "},
			want: "MUANGPRUAN INTANIT",
		},
		{
			name:"13. Input contains non-printable characters (ASCII: 219-222)",
			args: args{input: "\u007F\u0080\u0081\u0082\u0083\u0084MUANGPRUAN \u0085\u0086\u0087\u0088\u0089\u008A\u008B\u008C\u008D\u008E\u008F\u0090\u0091\u0092\u0093\u0094\u0095\u0096\u0097\u0098\u0099INTANIT\u009A\u009B\u009C\u009D\u009E\u009F "},
			want: "MUANGPRUAN INTANIT",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveSpecialCharacters(tt.args.input); got != tt.want {
				t.Errorf("RemoveSpecialCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
