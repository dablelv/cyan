package strings

import "testing"

func TestFormatThousand(t *testing.T) {
	type args struct {
		s   string
		sep []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "no fmt",
			args: args{
				s: "11.10",
			},
			want:    "11.10",
			wantErr: false,
		},
		{
			name: "need fmt for decimal",
			args: args{
				s: "1111.10",
			},
			want:    "1,111.10",
			wantErr: false,
		},
		{
			name: "need fmt for decimal with specified separator",
			args: args{
				s:   "1111.10",
				sep: []string{"_"},
			},
			want:    "1_111.10",
			wantErr: false,
		},
		{
			name: "need fmt for integer",
			args: args{
				s: "1111111111",
			},
			want:    "1,111,111,111",
			wantErr: false,
		},
		{
			name: "need fmt for integer with specified separator",
			args: args{
				s:   "1111111111",
				sep: []string{"_"},
			},
			want:    "1_111_111_111",
			wantErr: false,
		},
		{
			name: "input is invalid",
			args: args{
				s: "11111.11111.123456789",
			},
			want:    "11111.11111.123456789",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FormatThousand(tt.args.s, tt.args.sep...)
			if (err != nil) != tt.wantErr {
				t.Errorf("FormatThousand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FormatThousand() = %v, want %v", got, tt.want)
			}
		})
	}
}
