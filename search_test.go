package xivapi

import (
	"context"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"net/http"
	"reflect"
	"testing"
)

func TestSearchService_Items(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		testFormValues(t, r, values{
			"indexes": "Item",
			"string":  "aiming",
		})

		fmt.Fprint(w, `{"Pagination":{"Page":1,"PageNext":2,"PagePrev":null,"PageTotal":2,"Results":2,"ResultsPerPage":2,"ResultsTotal":4},"Results":[{"ID": 2937,"Icon": "/i/040000/040916.png","Name": "Darklight Eyepatch of Aiming","Url": "/Item/2937","UrlType": "Item","_": "item","_Score": 1},{"ID": 2949,"Icon": "/i/040000/040118.png","Name": "Allagan Visor of Aiming","Url": "/Item/2949","UrlType": "Item","_": "item","_Score": 1}],"SpeedMS":20}`)
	})

	ctx := context.Background()
	result, _, err := client.Search.Items(ctx, "aiming")
	if err != nil {
		t.Errorf("Search.Items returned error: %v", err)
	}
	var want = &ItemsSearchResult{
		PaginatedResult: &PaginatedResult{Pagination: &Pagination{
			Page:         Int(1),
			PageNext:     Int(2),
			PagePrev:     nil,
			Results:      Int(2),
			ResultsTotal: Int(4),
		}},
		Items: []*Item{{
			ID:      2937,
			Icon:    "/i/040000/040916.png",
			Name:    "Darklight Eyepatch of Aiming",
			URL:     "/Item/2937",
			URLType: "Item",
			Type:    "item",
			Score:   1,
		}, {
			ID:      2949,
			Icon:    "/i/040000/040118.png",
			Name:    "Allagan Visor of Aiming",
			URL:     "/Item/2949",
			URLType: "Item",
			Type:    "item",
			Score:   1,
		}},
		Speed: Int(20),
	}

	if !cmp.Equal(result, want) {
		t.Errorf("Search.Items returned %+v, want %+v", result, want)
	}
}

func TestSearchService_search(t *testing.T) {
	type args struct {
		ctx        context.Context
		searchType string
		query      string
		result     interface{}
	}
	tests := []struct {
		name    string
		s       SearchService
		args    args
		want    *Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.search(tt.args.ctx, tt.args.searchType, tt.args.query, tt.args.result)
			if (err != nil) != tt.wantErr {
				t.Errorf("search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("search() got = %v, want %v", got, tt.want)
			}
		})
	}
}
