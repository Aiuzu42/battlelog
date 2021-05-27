package stratagems

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	factions  = []Faction{}
	factionSt = []FactionStratagems{}

	phases = [...]string{"any", "before", "deployment", "start", "command", "movement", "enemy movement", "psychic", "enemy psychic", "shooting", "enemy shooting", "charge", "enemy charge", "fight", "enemy fight", "morale", "enemy morale", "end"}
)

const (
	suffix = ".txt"
)

type Stratagem struct {
	Name        string   `json:"name"`
	Sub         string   `json:"sub"`
	Type        string   `json:"type"`
	Fluff       string   `json:"fluff"`
	Description string   `json:"description"`
	CPCost      string   `json:"cp_cost"`
	Tags        []string `json:"tags"`
}

type FactionStratagems struct {
	Name       string      `json:"name"`
	Stratagems []Stratagem `json:"stratagems"`
}

type Faction struct {
	Name string   `json:"name"`
	Subs []string `json:"subs"`
}

func Load() {
	files, err := ioutil.ReadDir("../../")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), suffix) {
			f, fs, err := parseFile(f.Name())
			if err != nil {
				log.Printf("Error parsing file [%s]: %s", f.Name, err.Error())
			} else {
				factions = append(factions, f)
				factionSt = append(factionSt, fs)
			}
		}
	}
}

func parseFile(p string) (Faction, FactionStratagems, error) {
	lines, err := readLines(p)
	if err != nil {
		return Faction{}, FactionStratagems{}, err
	}
	ns := strings.TrimSuffix(p, suffix)
	fs := FactionStratagems{Name: ns}
	f := Faction{Name: ns}
	for i := 0; i < len(lines); i = i + +6 {
		s := Stratagem{}
		s.Name = lines[i+1]
		s.Description = lines[i+4]
		s.Fluff = lines[i+3]
		s.Sub, s.Type = parseSubAndType(lines[i+2])
		if s.Sub != f.Name {
			f.Subs = addAsSet(f.Subs, s.Sub)
		}
		s.CPCost = lines[i]
		s.Tags = strings.Split(lines[i+5], ",")
		fs.Stratagems = append(fs.Stratagems, s)
	}
	return f, fs, nil
}

func readLines(p string) ([]string, error) {
	file, err := os.Open("../../" + p)
	if err != nil {
		return nil, fmt.Errorf("unable to open file %s", p)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if len(lines) < 6 || len(lines)%6 != 0 {
		return nil, fmt.Errorf("incorrect number of lines [%d] on file %s", len(lines), p)
	}
	return lines, nil
}

func addAsSet(set []string, e string) []string {
	for _, s := range set {
		if s == e {
			return set
		}
	}
	return append(set, e)
}

func in(s string, a []string) bool {
	for _, e := range a {
		if e == s {
			return true
		}
	}
	return false
}

func sliceIn(s []string, a []string) bool {
	for _, se := range s {
		for _, ae := range a {
			if se == ae {
				return true
			}
		}
	}
	return false
}

func filterBySubfaction(b string, st string, s []string) bool {
	if b == st {
		return true
	}
	return in(st, s)
}

func filterByPhase(sp []string, rp []string) bool {
	if len(rp) == 0 {
		return true
	}
	if len(sp) > 0 && sp[0] == "any" {
		return true
	}
	return sliceIn(sp, rp)
}

func parseSubAndType(s string) (string, string) {
	sp := strings.Split(s, "%")
	if len(sp) > 1 {
		return sp[0], sp[1]
	} else {
		return sp[0], ""
	}
}

func GetAllStratagemFactions() []Faction {
	return factions
}

//GetFactionStratagemsSub obtains the stratagems for the faction with name n
//s is a list of which sub-factions stratagems should also be returned, an empty list indicates to return only base
//stratagems
//p is a list of in which phases are the stratagems returned active, an empty slice indicates to return all
func GetFactionStratagems(n string, s []string, p []string) FactionStratagems {
	if n == "" || s == nil || p == nil {
		return FactionStratagems{}
	}
	for i := range factionSt {
		if factionSt[i].Name == n {
			fs := FactionStratagems{Name: factionSt[i].Name, Stratagems: []Stratagem{}}
			for _, st := range factionSt[i].Stratagems {
				if filterBySubfaction(n, st.Sub, s) && filterByPhase(st.Tags, p) {
					fs.Stratagems = append(fs.Stratagems, st)
				}
			}
			return fs
		}
	}
	return FactionStratagems{}
}

func ValidatePhase(p string) bool {
	return in(p, phases[:])
}

func GetPhases() []string {
	return phases[0:]
}
