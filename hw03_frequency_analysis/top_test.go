package hw03frequencyanalysis

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

// Change to true if needed.
var taskWithAsteriskIsCompleted = false

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

var smallLenText = `текст с переносами строки
	табуляцией, с запятыми и со знаками!`

func TestTop10(t *testing.T) {
	type args struct {
		row string
	}
	tests := []struct {
		name        string
		args        args
		want        []string
		wantLen     int
		excludeTest bool
	}{
		{
			name: "no words in empty string",
			args: args{
				row: "",
			},
			want:        nil,
			wantLen:     0,
			excludeTest: false,
		},
		{
			name: "positive test string with small len",
			args: args{
				row: smallLenText,
			},
			want: []string{
				"с",
				"запятыми",
				"знаками!",
				"и",
				"переносами",
				"со",
				"строки",
				"табуляцией,",
				"текст",
			},
			wantLen:     9,
			excludeTest: false,
		},
		{
			name: "positive test string with big len v1",
			args: args{
				row: text,
			},
			want: []string{
				"он",
				"а",
				"и",
				"ты",
				"что",
				"-",
				"Кристофер",
				"если",
				"не",
				"то",
			},
			wantLen:     10,
			excludeTest: taskWithAsteriskIsCompleted,
		},
		{
			name: "positive test string with big len v2",
			args: args{
				row: text,
			},
			want: []string{
				"а",
				"он",
				"и",
				"ты",
				"что",
				"в",
				"его",
				"если",
				"кристофер",
				"не",
			},
			wantLen:     10,
			excludeTest: !taskWithAsteriskIsCompleted,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.excludeTest {
				t.SkipNow()
			}
			if got := Top10(tt.args.row); !reflect.DeepEqual(got, tt.want) {
				require.Len(t, len(got), tt.wantLen)
				require.Equal(t, tt.want, got)
			}
		})
	}
}

func Test_getUniqMap(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want map[string]counter
	}{
		{
			name: "empty input",
			args: args{
				words: []string{},
			},
			want: nil,
		},
		{
			name: "positive",
			args: args{
				words: []string{"слово", "слово", "слово", "два", "два"},
			},
			want: map[string]counter{
				"слово": {
					count: 3,
					value: "слово",
				},
				"два": {
					count: 2,
					value: "два",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getUniqMap(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				require.Equal(t, tt.want, got)
			}
		})
	}
}

func Test_transformToSortSlice(t *testing.T) {
	type args struct {
		uniq map[string]counter
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "empty input",
			args: args{
				uniq: nil,
			},
			want: nil,
		},
		{
			name: "positive",
			args: args{
				uniq: map[string]counter{
					"слово": {
						count: 3,
						value: "слово",
					},
					"два": {
						count: 2,
						value: "два",
					},
					"три": {
						count: 8,
						value: "три",
					},
				},
			},
			want: []string{"три", "слово", "два"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transformToSortSlice(tt.args.uniq); !reflect.DeepEqual(got, tt.want) {
				require.Equal(t, tt.want, got)
			}
		})
	}
}
