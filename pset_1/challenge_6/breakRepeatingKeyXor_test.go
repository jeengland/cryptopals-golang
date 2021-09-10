package challenge_6

import (
	"cryptopals/pset_1/challenge_4"
	"encoding/base64"
	"fmt"
	"strings"
	"testing"
)

var inputFile = "./challenge_6_input.txt"

func TestBreakRepeatingKeyXor(t *testing.T) {
	input, err := base64.StdEncoding.DecodeString(strings.Join(challenge_4.FileToStringSlice(inputFile), ""))
	if err != nil {
		fmt.Println(err)
	}
	ans := BreakRepeatingKeyXor(input)
	if ans != decodedMessage {
		t.Errorf("rcvd %s;\n want %s\n", ans, decodedMessage)
	}
}

func TestInterpolate(t *testing.T) {
	input := []string{"HpBta!", "ayihy!", "p rd!"}
	output := "Happy Birthday!!!"
	ans := Interpolate(input)
	if ans != output {
		t.Errorf("rcvd %s;\n want %s\n", ans, output)
	}
}

func TestSplitByKeySize(t *testing.T) {
	input := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	keysizes := []int{5, 3, 2}
	slice1 := [][]byte{[]byte{1, 6}, []byte{2, 7}, []byte{3, 8}, []byte{4, 9}, []byte{5, 10}}
	slice2 := [][]byte{[]byte{1, 4, 7, 10}, []byte{2, 5, 8}, []byte{3, 6, 9}}
	slice3 := [][]byte{[]byte{1, 3, 5, 7, 9}, []byte{2, 4, 6, 8, 10}}
	output := [][][]byte{slice1, slice2, slice3}
	ans := SplitByKeySize(input, keysizes)
	for i := range ans {
		for j := range ans[i] {
			for k := range ans[i][j] {
				if ans[i][j][k] != output[i][j][k] {
					t.Errorf("rcvd %v; want %d\n", ans[i][j][k], output[i][j][k])
				}
			}
		}
	}
}

func TestFindBestKeySizes(t *testing.T) {
	input, err := base64.StdEncoding.DecodeString(strings.Join(challenge_4.FileToStringSlice(inputFile), ""))
	if err != nil {
		fmt.Println(err)
	}
	output := []int{5, 29, 3}
	ans := FindBestKeySizes(input)
	if ans[0] != output[0] || ans[1] != output[1] || ans[2] != output[2] {
		t.Errorf("rcvd %v;\n want %d\v", ans, output)
	}
}

func TestGetEditDistance(t *testing.T) {
	testString1 := []byte("this is a test")
	testString2 := []byte("wokka wokka!!!")
	output := 37
	ans := GetEditDistance(testString1, testString2)
	if ans != output {
		t.Errorf("rcvd %d;\n want %d\n", ans, output)
	}
}

func TestBytesToBinary(t *testing.T) {
	input := []byte("C")
	output := "01000011"
	ans := BytesToBinary(input)
	if ans != output {
		t.Errorf("rcvd %s;\n want %s\n", ans, output)
	}
}

var decodedMessage = `I'm back and I'm ringin' the bell 
A rockin' on the mike while the fly girls yell 
In ecstasy in the back of me 
Well that's my DJ Deshay cuttin' all them Z's 
Hittin' hard and the girlies goin' crazy 
Vanilla's on the mike, man I'm not lazy. 

I'm lettin' my drug kick in 
It controls my mouth and I begin 
To just let it flow, let my concepts go 
My posse's to the side yellin', Go Vanilla Go! 

Smooth 'cause that's the way I will be 
And if you don't give a damn, then 
Why you starin' at me 
So get off 'cause I control the stage 
There's no dissin' allowed 
I'm in my own phase 
The girlies sa y they love me and that is ok 
And I can dance better than any kid n' play 

Stage 2 -- Yea the one ya' wanna listen to 
It's off my head so let the beat play through 
So I can funk it up and make it sound good 
1-2-3 Yo -- Knock on some wood 
For good luck, I like my rhymes atrocious 
Supercalafragilisticexpialidocious 
I'm an effect and that you can bet 
I can take a fly girl and make her wet. 

I'm like Samson -- Samson to Delilah 
There's no denyin', You can try to hang 
But you'll keep tryin' to get my style 
Over and over, practice makes perfect 
But not if you're a loafer. 

You'll get nowhere, no place, no time, no girls 
Soon -- Oh my God, homebody, you probably eat 
Spaghetti with a spoon! Come on and say it! 

VIP. Vanilla Ice yep, yep, I'm comin' hard like a rhino 
Intoxicating so you stagger like a wino 
So punks stop trying and girl stop cryin' 
Vanilla Ice is sellin' and you people are buyin' 
'Cause why the freaks are jockin' like Crazy Glue 
Movin' and groovin' trying to sing along 
All through the ghetto groovin' this here song 
Now you're amazed by the VIP posse. 

Steppin' so hard like a German Nazi 
Startled by the bases hittin' ground 
There's no trippin' on mine, I'm just gettin' down 
Sparkamatic, I'm hangin' tight like a fanatic 
You trapped me once and I thought that 
You might have it 
So step down and lend me your ear 
'89 in my time! You, '90 is my year. 

You're weakenin' fast, YO! and I can tell it 
Your body's gettin' hot, so, so I can smell it 
So don't be mad and don't be sad 
'Cause the lyrics belong to ICE, You can call me Dad 
You're pitchin' a fit, so step back and endure 
Let the witch doctor, Ice, do the dance to cure 
So come up close and don't be square 
You wanna battle me -- Anytime, anywhere 

You thought that I was weak, Boy, you're dead wrong 
So come on, everybody and sing this song 

Say -- Play that funky music Say, go white boy, go white boy go 
play that funky music Go white boy, go white boy, go 
Lay down and boogie and play that funky music till you die. 

Play that funky music Come on, Come on, let me hear 
Play that funky music white boy you say it, say it 
Play that funky music A little louder now 
Play that funky music, white boy Come on, Come on, Come on 
Play that funky music 
`