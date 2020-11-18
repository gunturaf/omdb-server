package presenters

import (
	"reflect"
	"testing"

	"github.com/gunturaf/omdb-server/entity"
)

func TestSearchResultToProto(t *testing.T) {
	type args struct {
		inp *entity.OMDBSearchResult
	}
	tests := []struct {
		name string
		args args
		want *entity.SearchReply
	}{
		{
			name: "1",
			args: args{
				inp: &entity.OMDBSearchResult{
					Search: []entity.OMDBResultCompact{
						{
							IMDBID: "DavidBowie",
						},
					},
					Response:     "True",
					TotalResults: "3",
				},
			},
			want: &entity.SearchReply{
				Search: []*entity.SearchEntry{
					{
						ImdbID: "DavidBowie",
					},
				},
				Response:     "True",
				TotalResults: "3",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchResultToProto(tt.args.inp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchResultToProto() = %v, want %v", got, tt.want)
			}
		})
	}
}
