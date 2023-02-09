package main

import (
	"bytes"
	"database/sql"
	"encoding/csv"
	"fmt"
	punksranking "github.com/currencytycoon/punkranking"
	"html"
	"io/ioutil"
	"math"
	"strconv"
	"testing"
)

func TestScore(*testing.T) {
	values := []float64{24.59, 2.8600, 0.44}
	var y float64
	p := 2.0
	for i := range values {
		y += math.Pow(math.Log(values[i]), p)
	}
	b := math.Exp(math.Pow(y/float64(len(values)), 1.0/p))
	fmt.Println(b)
}

func TestScore2(*testing.T) {
	values := []float64{24.59, 2.8600, 0.44}
	var y float64
	p := 2.0
	for i := range values {
		y += math.Pow(math.Log(values[i]), p)
	}
	b := math.Exp(math.Pow(y/float64(len(values)), 1.0/p))
	fmt.Println(b)

	values[2] = 0.00044
	for i := range values {
		y += math.Pow(math.Log(values[i]), p)
	}
	b = math.Exp(math.Pow(y/float64(len(values)), 1.0/p))
	fmt.Println(b)

	values[2] = 1
	for i := range values {
		y += math.Pow(math.Log(values[i]), p)
	}
	b = math.Exp(math.Pow(y/float64(len(values)), 1.0/p))
	fmt.Println(b)
}

func TestScore3(*testing.T) {
	values := []float64{24.59, 2.8600, 0.44}
	var y float64
	//	p := 2.0
	for i := range values {
		y += math.Log(values[i])
	}
	b := y // math.Exp(math.Pow(y/float64(len(values)), 1.0/p))
	fmt.Println(b)

	values[2] = 0.00044
	for i := range values {
		y += math.Log(values[i])
	}
	b = y // math.Exp(math.Pow(y/float64(len(values)), 1.0/p))
	fmt.Println(b)

	values[2] = 1
	for i := range values {
		y += math.Log(values[i])
	}
	b = y // math.Exp(math.Pow(y/float64(len(values)), 1.0/p))
	fmt.Println(b)
}

func TestUnescape(t *testing.T) {

	all, err := ioutil.ReadFile("../../ranks.md")
	if err != nil {
		t.Error(err)
		return
	}
	str := html.UnescapeString(string(all))

	err = ioutil.WriteFile("../../ranks-fixed.md", []byte(str), 0644)
	if err != nil {
		t.Error(err)
		return
	}

}

var params = `0,4,m,0,598
1,3,m,0,1861
2,2,m,0,1857
3,1,m,0,1723
4,4,f,0,420
5,3,f,0,1145
6,2,f,0,1174
7,1,f,0,1101
8,5,m,0,88
9,6,m,0,24
10,7,m,0,9
11,Rosy Cheeks,m,1,128
12,Luxurious Beard,m,4,286
13,Clown Hair Green,m,3,148
14,Mohawk Dark,m,3,429
15,Cowboy Hat,m,11,142
16,Mustache,m,4,288
17,Clown Nose,m,7,212
18,Cigarette,m,9,961
19,Nerd Glasses,m,6,572
20,Regular Shades,m,6,527
21,Knitted Cap,m,11,419
22,Shadow Beard,m,4,526
23,Frown,m,8,261
24,Cap Forward,m,11,254
25,Goat,m,4,295
26,Mole,m,2,644
27,Purple Hair,m,3,165
28,Small Shades,m,6,378
29,Shaved Head,m,3,300
30,Classic Shades,m,6,502
31,Vape,m,9,272
32,Silver Chain,m,12,156
33,Smile,m,8,238
34,Big Shades,m,6,535
35,Mohawk Thin,m,3,441
36,Beanie,m,11,44
37,Cap,m,11,351
38,Clown Eyes Green,m,5,382
39,Normal Beard Black,m,4,289
40,Medical Mask,m,9,175
41,Normal Beard,m,4,292
42,VR,m,6,332
43,Eye Patch,m,5,461
44,Wild Hair,m,3,447
45,Top Hat,m,11,115
46,Bandana,m,11,481
47,Handlebars,m,4,263
48,Frumpy Hair,m,3,442
49,Crazy Hair,m,3,414
50,Police Cap,m,11,203
51,Buck Teeth,m,8,78
52,Do-rag,m,11,300
53,Front Beard,m,4,273
54,Spots,m,2,124
55,Big Beard,m,4,146
56,Vampire Hair,m,3,147
57,Peak Spike,m,3,303
58,Chinstrap,m,4,282
59,Fedora,m,11,186
60,Earring,m,10,2459
61,Horned Rim Glasses,m,6,535
62,Headband,m,11,406
63,Pipe,m,9,317
64,Messy Hair,m,3,460
65,Front Beard Dark,m,4,260
66,Hoodie,m,11,259
67,Gold Chain,m,12,169
68,Muttonchops,m,4,303
69,Stringy Hair,m,3,463
70,Eye Mask,m,6,293
71,3D Glasses,m,6,286
72,Clown Eyes Blue,m,5,384
73,Mohawk,m,3,441
74,Pilot Helmet,f,11,54
75,Tassle Hat,f,11,178
76,Hot Lipstick,f,8,696
77,Blue Eye Shadow,f,5,266
78,Straight Hair Dark,f,3,148
79,Choker,f,12,48
80,Crazy Hair,f,3,414
81,Regular Shades,f,6,527
82,Wild Blonde,f,3,144
83,3D Glasses,f,6,286
84,Mole,f,2,644
85,Wild White Hair,f,3,136
86,Spots,f,2,124
87,Frumpy Hair,f,3,442
88,Nerd Glasses,f,6,572
89,Tiara,f,11,55
90,Orange Side,f,3,68
91,Red Mohawk,f,3,147
92,Messy Hair,f,3,460
93,Clown Eyes Blue,f,5,384
94,Pipe,f,9,317
95,Wild Hair,f,3,447
96,Purple Eye Shadow,f,5,262
97,Stringy Hair,f,3,463
98,Dark Hair,f,3,157
99,Eye Patch,f,6,461
100,Blonde Short,f,3,129
101,Classic Shades,f,6,502
102,Eye Mask,f,6,293
103,Clown Hair Green,f,3,148
104,Cap,f,11,351
105,Medical Mask,f,9,175
106,Bandana,f,11,481
107,Purple Lipstick,f,8,655
108,Clown Nose,f,7,212
109,Headband,f,11,406
110,Pigtails,f,3,94
111,Straight Hair Blonde,f,3,144
112,Knitted Cap,f,11,419
113,Clown Eyes Green,f,5,382
114,Cigarette,f,9,961
115,Welding Goggles,f,6,86
116,Mohawk Thin,f,3,441
117,Gold Chain,f,12,169
118,VR,f,6,332
119,Vape,f,9,272
120,Pink With Hat,f,3,95
121,Blonde Bob,f,3,147
122,Mohawk,f,3,441
123,Big Shades,f,6,535
124,Earring,f,10,2459
125,Green Eye Shadow,f,5,271
126,Straight Hair,f,3,151
127,Rosy Cheeks,f,1,128
128,Half Shaved,f,3,147
129,Mohawk Dark,f,3,429
130,Black Lipstick,f,8,617
131,Horned Rim Glasses,f,6,535
132,Silver Chain,f,12,156
`

func TestPopulate(t *testing.T) {
	var db *sql.DB
	var err error
	if db, err = punksranking.SetupDB("./config.json"); err != nil {
		t.Error(err)
		return
	}
	defer punksranking.Close()

	csvReader := csv.NewReader(bytes.NewReader([]byte(params)))
	records, _ := csvReader.ReadAll()

	if stmt, err := db.Prepare(
		`UPDATE attributes SET layer_id = ? WHERE name =? ;`,
	); err != nil {
		t.Error(err)
		return
	} else {

		for i := range records {
			lid, _ := strconv.Atoi(records[i][3])

			_, err := stmt.Exec(lid, records[i][1])
			if err != nil {
				t.Error(err)
			}
		}

	}

}
