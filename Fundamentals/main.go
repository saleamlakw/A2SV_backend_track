package main
import (
	"unicode"
	"strings"
	"fmt"
)
func WordFrequency(s string) map[string]int{
	words:=strings.Fields(s)
	for ind,_ := range words{
		var finalString string
		for _,char :=range words[ind]{
			if unicode.IsLetter(char){
				finalString+=strings.ToLower(string(char))
			}
		}
		words[ind]=finalString
	}

	frequency:= make(map [string]int)
	for _,word :=range words{
		frequency[word]+=1
	}

	return frequency 
}

func IsPalindrome(s string)bool{
	var finalString string
	for _,val := range s{
		if unicode.IsLetter(val) || unicode.IsDigit(val){
			finalString+=strings.ToLower(string(val))
		}
	}
	left,right:=0,len(finalString)-1
	for left<right{
		if finalString[left]!=finalString[right]{
			return false
		}
		right-=1
		left+=1
	}
	return true
}

