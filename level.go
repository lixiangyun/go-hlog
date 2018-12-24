package hlog

import (
	"errors"
	"sync"
)

type level struct {
	name string
	id   int
	flag int
}

type levels struct {
	lv_map   map[string]level
	lv_array []level
	sync.RWMutex
}

var g_hlogLevels levels

func init() {
	g_hlogLevels.lv_map = make(map[string]level, 255)
	g_hlogLevels.lv_array = make([]level, 0, 255)
}

func LevelAdd(cfg LevelConfig) error {
	g_hlogLevels.Lock()

	defer g_hlogLevels.Unlock()

	for _, v := range g_hlogLevels.lv_map {
		if v.id == cfg.Level || v.name == cfg.Name {
			return errors.New("level " + cfg.Name + " duplicate!")
		}
	}

	flag, err := syslogAtoi(cfg.Syslog)
	if err != nil {
		return err
	}

	var lv level

	lv.name = cfg.Name
	lv.id = cfg.Level
	lv.flag = flag

	g_hlogLevels.lv_map[cfg.Name] = lv
	g_hlogLevels.lv_array = append(g_hlogLevels.lv_array, lv)

	return nil
}
