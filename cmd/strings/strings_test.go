package strings_test

import (
	"golibs/cmd/strings"
	"reflect"
	"testing"
)

func TestGetSimilarValue(t *testing.T) {
	type args struct {
		source string
		target string
		option strings.Options
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Equals",
			args: args{
				source: "carcasa",
				target: "carcasa",
				option: strings.DefaultOptions,
			},
			want: 1,
		},
		{
			name: "Test_1",
			args: args{
				source: "carcasa",
				target: "carroza",
				option: strings.DefaultOptions,
			},
			want: 0.5714285714285714,
		},
		{
			name: "Test_2",
			args: args{
				source: "carcasa",
				target: "carcas",
				option: strings.DefaultOptions,
			},
			want: 0.8571428571428572,
		},
		{
			name: "Test_3",
			args: args{
				source: "carcasa",
				target: "karcaza",
				option: strings.Options{
					InsCost: 1.25,
					DelCost: 1,
					SubCost: 1.5,
				},
			},
			want: 0.5714285714285714,
		},
		{
			name: "Test_3",
			args: args{
				source: "carcasa",
				target: "karcaza",
				option: strings.Options{
					InsCost: 1.25,
					DelCost: 1,
					SubCost: 0.5,
				},
			},
			want: 0.8571428571428572,
		},
		{
			name: "Test_4",
			args: args{
				source: "Reynier Gonzalez",
				target: "reyNier Gonzalez",
				option: strings.Options{
					InsCost: 1.25,
					DelCost: 1,
					SubCost: 1.5,
				},
			},
			want: 1,
		},
		{
			name: "Test_5",
			args: args{
				source: "Reynier Gonzalez",
				target: "reyNier González",
				option: strings.DefaultOptions,
			},
			want: 1,
		},
		{
			name: "Test_6",
			args: args{
				source: "Gonzalez Reynier",
				target: "reyNier González",
				option: strings.Options{
					InsCost: 1.25,
					DelCost: 1,
					SubCost: 1.5,
				},
			},
			want: -0.050000000000000044,
		},
		{
			name: "Test_7",
			args: args{
				source: "Asheville",
				target: "Arizona",
				option: strings.DefaultOptions,
			},
			want: 0.11111111111111116,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strings.GetLevenshteinSimilarity(tt.args.source, tt.args.target, tt.args.option); got != tt.want {
				t.Errorf("GetLevenshteinSimilarity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJaroWinklerDistance(t *testing.T) {
	type args struct {
		source string
		target string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Equals",
			args: args{
				source: "carcasa",
				target: "carcasa",
			},
			want: 1,
		},
		{
			name: "Test_1",
			args: args{
				source: "carcasa",
				target: "carroza",
			},
			want: 0.8,
		},
		{
			name: "Test_2",
			args: args{
				source: "carcasa",
				target: "carcas",
			},
			want: 0.9714285714285714,
		},
		{
			name: "Test_3",
			args: args{
				source: "carcasa",
				target: "karcaza",
			},
			want: 0.8095238095238096,
		},
		{
			name: "Test_4",
			args: args{
				source: "Reynier Gonzalez",
				target: "reyNier Gonzalez",
			},
			want: 1,
		},
		{
			name: "Test_5",
			args: args{
				source: "Reynier Gonzalez",
				target: "reyNier González",
			},
			want: 1,
		},
		{
			name: "Test_6",
			args: args{
				source: "Gonzalez Reynier",
				target: "reyNier González",
			},
			want: 0.5253968253968254,
		},
		{
			name: "Test_7",
			args: args{
				source: "Asheville",
				target: "Arizona",
			},
			want: 0.5026455026455027,
		},
		{
			name: "Test_8",
			args: args{
				source: "González Reynier",
				target: "Reynier González Cruz",
			},
			want: 0.7742690058479532,
		},
		{
			name: "Test_9",
			args: args{
				source: "González Cruz Reynier",
				target: "Reynier González Cruz",
			},
			want: 0.7489035087719298,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strings.GetJaroWinklerSimilarity(tt.args.source, tt.args.target); got != tt.want {
				t.Errorf("GetJaroWinklerSimilarity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareNames(t *testing.T) {
	type args struct {
		fullName       string
		secondFullName string
	}
	tests := []struct {
		name string
		args args
		want strings.Match
	}{
		{
			name: "Test_1",
			args: args{
				fullName:       "Reynier Gonzalez Cruz",
				secondFullName: "Reynier González Cruz",
			},
			want: strings.Match{
				Percentage: strings.Distribution{
					Levenshtein: 1,
					JaroWinkler: 1,
					Media:       1,
				},
			},
		},
		{
			name: "Test_2",
			args: args{
				fullName:       "Reynier Gonzalez Cruz",
				secondFullName: "González Cruz Reynier",
			},
			want: strings.Match{
				Percentage: strings.Distribution{
					Levenshtein: 0.26315789473684215,
					JaroWinkler: 0.7489035087719298,
					Media:       0.5060307017543859,
				},
			},
		},
		{
			name: "Test_3",
			args: args{
				fullName:       "Reynier Gonzalez Cruz",
				secondFullName: "González Reynier    ",
			},
			want: strings.Match{
				Percentage: strings.Distribution{
					Levenshtein: 0.26315789473684215,
					JaroWinkler: 0.7742690058479532,
					Media:       0.5187134502923977,
				},
			},
		},
		{
			name: "Test_4",
			args: args{
				fullName:       "Arelys RIVERO CASTRO",
				secondFullName: "Reynier González Cruz",
			},
			want: strings.Match{
				Percentage: strings.Distribution{
					Levenshtein: 0.26315789473684215,
					JaroWinkler: 0.5939571150097466,
					Media:       0.42855750487329436,
				},
			},
		},
		{
			name: "Test_5",
			args: args{
				fullName:       "Arelys RIVERO CASTRO",
				secondFullName: "Arelys RIVERO",
			},
			want: strings.Match{
				Percentage: strings.Distribution{
					Levenshtein: 0.6666666666666667,
					JaroWinkler: 0.9333333333333333,
					Media:       0.8,
				},
			},
		},
		{
			name: "Test_6",
			args: args{
				fullName:       "MARTHA",
				secondFullName: "MARHTA",
			},
			want: strings.Match{
				Percentage: strings.Distribution{
					Levenshtein: 0.6666666666666667,
					JaroWinkler: 0.9611111111111111,
					Media:       0.8138888888888889,
				},
			},
		},
		{
			name: "Test_7",
			args: args{
				fullName:       "MARTHA Mesa",
				secondFullName: "MARHTA Mesa",
			},
			want: strings.Match{
				Percentage: strings.Distribution{
					Levenshtein: 0.8,
					JaroWinkler: 0.9766666666666667,
					Media:       0.8883333333333334,
				},
			},
		},
		{
			name: "Test_8",
			args: args{
				fullName:       "MARTHA Mesa Silva",
				secondFullName: "MARTA Mesa",
			},
			want: strings.Match{
				Percentage: strings.Distribution{
					Levenshtein: 0.6,
					JaroWinkler: 0.92,
					Media:       0.76,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strings.GetSimilarity(tt.args.fullName, tt.args.secondFullName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSimilarity() = %v, want %v", got, tt.want)
			}
		})
	}
}
