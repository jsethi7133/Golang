type Node struct{
    Data int
    KeyPtr *list.Element
}

type LRUCache struct {
    Queue *list.List
    Items map[int]*Node
    Capacity int
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
        Queue: list.New(),
        Items: make(map[int]*Node),
        Capacity: capacity,
    }
}


func (this *LRUCache) Get(key int) int {
    if val,ok:=this.Items[key];ok{
        this.Queue.MoveToFront(val.KeyPtr)
        return val.Data
    }
    return -1
}

func (this *LRUCache) Put(key int, value int)  {
    if val,ok:=this.Items[key];!ok{
        if this.Capacity==len(this.Items){
            back:=this.Queue.Back()
            this.Queue.Remove(back)
            delete(this.Items, back.Value.(int))
        }
        this.Items[key]=&Node{Data: value, KeyPtr: this.Queue.PushFront(key)}
    }else{
        val.Data=value
        this.Items[key]=val
        this.Queue.MoveToFront(val.KeyPtr)
    }
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
