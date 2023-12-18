package utilities

// Queue Structure de la file
type Queue struct {
	items []interface{}
}

// Put ajoute un élément à la file
func (q *Queue) Put(item interface{}) {
	q.items = append(q.items, item)
}

// Get retire et renvoie l'élément en tête de la file
func (q *Queue) Get() interface{} {
	if len(q.items) == 0 {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// IsEmpty vérifie si la file est vide
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size renvoie la taille de la file
func (q *Queue) Size() int {
	return len(q.items)
}
