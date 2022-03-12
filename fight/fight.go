package fight

import "letthemfight/repositories"

type Fight struct {
	Groups map[string]repositories.Group
}

func (f *Fight) AddGroup(name string, group *repositories.Group) {
	f.Groups[name] = *group
}

func NewFight() *Fight {
	return &Fight{Groups: make(map[string]repositories.Group)}
}
