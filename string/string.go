package string

func Encode(s string,n int) string{
	result:=""
	change:=n%26
	for i:=0;i<len(s);i++{
		if s[i]<=122{
			result+=string(s[i]+byte(change))
		}
	}
	return result
}