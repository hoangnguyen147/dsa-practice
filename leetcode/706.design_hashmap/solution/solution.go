package solution

type MyHashMap struct {
    data map[int]int
}


func Constructor() MyHashMap {
    return MyHashMap{
        data: make(map[int]int),
    }
}


func (this *MyHashMap) Put(key int, value int)  {
    this.data[key] = value
}


func (this *MyHashMap) Get(key int) int {
    v, ok := this.data[key]
    if !ok {
        return -1
    }
    return v
}


func (this *MyHashMap) Remove(key int)  {
    delete(this.data, key)
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
