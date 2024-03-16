package modules

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/abenz1267/walker/config"
)

type Workable interface {
	Entries(term string) []Entry
	Prefix() string
	Name() string
	Setup(cfg *config.Config) Workable
}

type MatchingType int

const (
	Fuzzy MatchingType = iota
	AlwaysTop
	AlwaysBottom
)

type Entry struct {
	Label             string       `json:"label,omitempty"`
	Sub               string       `json:"sub,omitempty"`
	Exec              string       `json:"exec,omitempty"`
	Terminal          bool         `json:"terminal,omitempty"`
	Icon              string       `json:"icon,omitempty"`
	IconIsImage       bool         `json:"icon_is_image,omitempty"`
	HideText          bool         `json:"hide_text,omitempty"`
	Categories        []string     `json:"categories,omitempty"`
	Notifyable        bool         `json:"notifyable,omitempty"`
	Class             string       `json:"class,omitempty"`
	History           bool         `json:"history,omitempty"`
	HistoryIdentifier string       `json:"history_identifier,omitempty"`
	Matching          MatchingType `json:"matching,omitempty"`
	ScoreFinal        float64      `json:"score_final,omitempty"`
	MinScoreToInclude float64      `json:"min_score_to_include,omitempty"`
	ScoreFuzzy        int          `json:"score_fuzzy,omitempty"`
	Identifier        string       `json:"-"`
	Used              int          `json:"-"`
	DaysSinceUsed     int          `json:"-"`
	LastUsed          time.Time    `json:"-"`
}

func readCache(name string, data any) bool {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		log.Println(err)
		return false
	}

	cacheDir = filepath.Join(cacheDir, "walker")

	path := filepath.Join(cacheDir, fmt.Sprintf("%s.json", name))

	if _, err := os.Stat(path); err == nil {
		file, err := os.Open(path)
		if err != nil {
			log.Println(err)
		}
		defer file.Close()

		b, err := io.ReadAll(file)
		if err != nil {
			log.Fatalln(err)
		}

		err = json.Unmarshal(b, &data)
		if err != nil {
			log.Fatalln(err)
		}

		return true
	}

	return false
}

func writeCache(name string, data any) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		log.Println(err)
		return
	}

	cacheDir = filepath.Join(cacheDir, "walker")

	b, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}

	err = os.MkdirAll(cacheDir, 0755)
	if err != nil {
		log.Println(err)
		return
	}

	err = os.WriteFile(filepath.Join(cacheDir, fmt.Sprintf("%s.json", name)), b, 0644)
	if err != nil {
		log.Println(err)
	}
}

func find(modules []config.Module, name string) *config.Module {
	for _, v := range modules {
		if v.Name == name {
			return &v
		}
	}

	return nil
}
