package LRU_cache


type node struct {
	Value interface{}
	Next *node
	Prev *node
}
type LRU struct {
	m          map[int]*node
	First      *node
	Last       *node
	MaxCap     int
	CurrentCap int
}

func AddElement(lru LRU, key int , element interface{}){
	if lru.MaxCap < lru.CurrentCap +1 {
		nodeToDelete := lru.Last
		lru.Last = nodeToDelete.Prev
		// en go pasa el garbage collector y borra ese nodo.
		// pero en otro lenguaje se tiene que liberar esa memoria
		lru.CurrentCap = lru.CurrentCap -1
	}
	// creo el nodo
	n := node{
		Value: element,
		Next: lru.First,
		Prev: nil,
	}
	// lo guardo en primer lugar y en el mapa
	lru.First = &n
	lru.m[key]= &n
	lru.CurrentCap++
}

func FindElement (lru LRU, key int) interface{}{
	n,ok := lru.m[key]
	if !ok {
		return nil
	}
	// si lo saco del medio,
	//referencio los punteros de los nodos prev y next
	prevNode := n.Prev
	NextNode := n.Next
	prevNode.Next = NextNode
	NextNode.Prev = prevNode

	// guardo el elemento como usado al principio
	n.Next = lru.First
	lru.First = n

	// devuelvo el elemento
	return n.Value
}