package tunnel

import (
    "container/list"
    "time"
)

type recyclerItem struct {
    when time.Time
    buf []byte
}

type recycler struct {
    q *list.List
    takeChan, giveChan chan []byte
}

func NewRecycler(size uint32) *recycler {
    r := &recycler{
        q: new(list.List),
        takeChan: make(chan []byte),
        giveChan: make(chan []byte),
    }
    go r.cycle(size)
    return r
}

func (r *recycler) cycle(size uint32) {
    for {
        if r.q.Len() == 0 {
            // put to front so we always use the most recent buf
            r.q.PushFront(recyclerItem{when: time.Now(), buf: make([]byte, size)})
        }
        i := r.q.Front()
        timeout := time.NewTimer(time.Minute)
        select {
        case b:= <-r.giveChan:
            timeout.Stop()
            r.q.PushFront(recyclerItem{when: time.Now(), buf: b})
        case r.takeChan <- i.Value.(recyclerItem).buf:
            timeout.Stop()
            r.q.Remove(i)
        case <-timeout.C:
            i := r.q.Front()
            for i != nil {
                n := i.Next()
                if time.Since(i.Value.(recyclerItem).when) > time.Minute {
                    r.q.Remove(i)
                    i.Value = nil
                }
                i = n
            }
        }
    }
}

func (r *recycler) take() []byte {
    return <-r.takeChan
}

func (r *recycler) give(b []byte) {
    r.giveChan <- b
}
