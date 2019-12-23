package queue

type Queue []int

func (q *Queue) push(v int){
	*q = append(*q,v)
}

func (q *Queue) pop () int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) isEmpty() bool {

}
