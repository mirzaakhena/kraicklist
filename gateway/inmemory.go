package gateway

import (
	"bufio"
	"challenge.haraj.com.sa/kraicklist/domain/repository"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"challenge.haraj.com.sa/kraicklist/domain/entity"
	"challenge.haraj.com.sa/kraicklist/domain/service"
	"challenge.haraj.com.sa/kraicklist/infrastructure/log"
)

type inmemoryGateway struct {
	RawDatas    map[int64]*entity.RawData
	IndexedData map[string][]int64
}

// NewInmemoryGateway ...
func NewInmemoryGateway() *inmemoryGateway {
	return &inmemoryGateway{}
}

func (r *inmemoryGateway) BeginTransaction(ctx context.Context) (context.Context, error) {
	log.InfoRequest(ctx, "called")

	// initialize data
	r.RawDatas = map[int64]*entity.RawData{}
	r.IndexedData = map[string][]int64{}

	return ctx, nil
}

func (r *inmemoryGateway) CommitTransaction(ctx context.Context) error {

	log.InfoRequest(ctx, "called")

	// nothing todo in in "inmemory" gateway implementation

	return nil
}

func (r *inmemoryGateway) RollbackTransaction(ctx context.Context) error {
	log.InfoRequest(ctx, "called")

	// force to clear it
	r.RawDatas = nil

	return nil
}

func (r *inmemoryGateway) ReadJSONData(ctx context.Context, req service.ReadJSONDataServiceRequest) error {
	log.InfoRequest(ctx, "called")

	// open file
	file, err := os.Open(req.Filename)
	if err != nil {
		return fmt.Errorf("unable to open source file due: %v", err)
	}
	defer file.Close()

	// read as gzip
	reader, err := gzip.NewReader(file)
	if err != nil {
		return fmt.Errorf("unable to initialize gzip reader due: %v", err)
	}

	// read the reader using scanner to construct records
	cs := bufio.NewScanner(reader)
	for cs.Scan() {
		var r entity.RawData
		if err := json.Unmarshal(cs.Bytes(), &r); err != nil {
			continue
		}
		err := req.ReadDataPerline(r)
		if err != nil {
			return err
		}
	}

	log.Info(ctx, "load data finished")

	return nil
}

func (r *inmemoryGateway) SaveData(ctx context.Context, obj *entity.RawData) error {

	// we store raw_data into slice
	r.RawDatas[obj.ID] = obj

	// we also need indexing the data

	// regex for remove all non alphanumeric
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")

	for _, word := range strings.Split(obj.Title, " ") {

		word = strings.ToLower(word)
		word = reg.ReplaceAllString(word, "")

		if _, exist := r.IndexedData[word]; !exist {
			r.IndexedData[word] = []int64{obj.ID}
		} else {
			r.IndexedData[word] = append(r.IndexedData[word], obj.ID)
		}
	}

	for _, word := range strings.Split(obj.Title, " ") {

		word = strings.ToLower(word)
		word = reg.ReplaceAllString(word, "")

		if _, exist := r.IndexedData[word]; !exist {
			r.IndexedData[word] = []int64{obj.ID}
		} else {
			r.IndexedData[word] = append(r.IndexedData[word], obj.ID)
		}
	}

	return nil
}

func (r *inmemoryGateway) FindDataByKeyword(ctx context.Context, keyword string) (*repository.FindDataByKeywordRepoResult, error) {
	log.InfoRequest(ctx, "called")

	keyword = strings.ToLower(keyword)

	result := []*entity.RawData{}
	containedKeywords := []string{}

	for keyPerWord, values := range r.IndexedData {

		if strings.Contains(keyPerWord, keyword) {

			containedKeywords = append(containedKeywords, keyPerWord)

			for keyID, data := range r.RawDatas {

				for _, value := range values {
					if keyID == value {
						result = append(result, data)
					}

				}

			}

		}

	}

	return &repository.FindDataByKeywordRepoResult{
		Result:           result,
		ContainedKeyword: containedKeywords,
	}, nil
}
