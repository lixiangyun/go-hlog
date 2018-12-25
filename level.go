package hlog

import (
	"errors"
	"strings"
	"sync"
)

type level struct {
	name_lower string
	name_upper string
	id         int
	flag       int
}

type levels struct {
	lv_map   map[string]level
	lv_array map[int]level
	sync.RWMutex
}

var g_hlogLevels levels

func init() {
	g_hlogLevels.lv_map = make(map[string]level, 255)
	g_hlogLevels.lv_array = make(map[int]level, 255)
}

func LevelAdd(cfg LevelConfig) error {
	g_hlogLevels.Lock()
	defer g_hlogLevels.Unlock()

	name_lower := strings.ToLower(cfg.Name)

	for _, v := range g_hlogLevels.lv_map {
		if v.id == cfg.Level || v.name_lower == name_lower {
			return errors.New("level " + cfg.Name + " duplicate!")
		}
	}

	flag, err := syslogAtoi(cfg.Syslog)
	if err != nil {
		return err
	}

	var lv level

	lv.name_lower = strings.ToLower(cfg.Name)
	lv.name_upper = strings.ToUpper(cfg.Name)
	lv.id = cfg.Level
	lv.flag = flag

	g_hlogLevels.lv_map[name_lower] = lv
	g_hlogLevels.lv_array[cfg.Level] = lv

	return nil
}

func LevelGetById(id int) level {
	g_hlogLevels.RLock()
	defer g_hlogLevels.RUnlock()
	lv, _ := g_hlogLevels.lv_array[id]
	return lv
}

func LevelGetByName(name string) level {
	g_hlogLevels.RLock()
	defer g_hlogLevels.RUnlock()
	lv, _ := g_hlogLevels.lv_map[strings.ToLower(name)]
	return lv
}
