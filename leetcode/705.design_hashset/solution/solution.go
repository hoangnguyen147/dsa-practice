package solution

type MyHashSet struct {
    data map[int]struct{}
}


func Constructor() MyHashSet {
    return MyHashSet{
        data: make(map[int]struct{}),
    }
}


func (this *MyHashSet) Add(key int)  {
    if _, exist := this.data[key]; !exist {
        this.data[key] = struct{}{}
    }
}


func (this *MyHashSet) Remove(key int)  {
    delete(this.data, key)
}


func (this *MyHashSet) Contains(key int) bool {
    _, exist := this.data[key]
    return exist
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
