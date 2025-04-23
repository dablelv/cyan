package strings

import (
	"testing"

	"github.com/dablelv/cyan/enum"
)

func TestEmailDesensitization(t *testing.T) {
	tests := []struct {
		name  string
		email string
		want  string
	}{
		{
			name:  "empty email",
			email: "",
			want:  "",
		},
		{
			name:  "invalid email",
			email: "invalid",
			want:  "invalid",
		},
		{
			name:  "one char local part",
			email: "a@x.com",
			want:  "a***@x.com",
		},
		{
			name:  "two chars local part",
			email: "ab@x.com",
			want:  "ab***@x.com",
		},
		{
			name:  "three chars local part",
			email: "abc@x.com",
			want:  "abc***@x.com",
		},
		{
			name:  "more than three chars local part",
			email: "abcd@x.com",
			want:  "abc***@x.com",
		},
		{
			name:  "chinese local part",
			email: "中文邮箱@x.com",
			want:  "中文邮***@x.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EmailDesensitization(tt.email); got != tt.want {
				t.Errorf("EmailDesensitization() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPhoneDesensitization(t *testing.T) {
	type args struct {
		region string
		number string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "mainland phone",
			args: args{enum.PhoneCountryCodeMainland, "13800138000"},
			want: "138****8000",
		},
		{
			name: "hong kong phone",
			args: args{enum.PhoneCountryCodeHongKong, "52123456"},
			want: "52****56",
		},
		{
			name: "macao phone",
			args: args{enum.PhoneCountryCodeMacao, "52123456"},
			want: "52****56",
		},
		{
			name: "taiwan phone",
			args: args{enum.PhoneCountryCodeTaiwan, "0912345678"},
			want: "09****678",
		},
		{
			name: "other region phone",
			args: args{enum.PhoneCountryCodeJapan, "07012345678"},
			want: "07****78",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PhoneDesensitization(tt.args.region, tt.args.number); got != tt.want {
				t.Errorf("PhoneDesensitization() = %v, want %v", got, tt.want)
			}
		})
	}
}
