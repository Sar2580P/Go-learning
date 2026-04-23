package game

import (
	"main/entity"

)

type World struct {
	entities []entity.Entity // need to define getter-setter for out of package access
}

// constructor
func NewWorld() *World{
	return &World{
		entities: []entity.Entity{},
	}
}

func (w *World) AddEntity(e entity.Entity){
	w.entities=append(w.entities, e)
}

func (w World) Entities() []entity.Entity{
	return w.entities
}

func (w World) GetEntities(tag string) []entity.Entity{
	var res []entity.Entity

	for _, ent :=range w.entities{
		if tag==ent.Tag(){
			res=append(res, ent)
		}
	}
	return res
}


func (w World) GetFirstEntity(tag string) (entity.Entity, bool){
	for _, ent :=range w.entities{
		if tag==ent.Tag(){
			return ent, true
		}
	}
	return nil, false
}