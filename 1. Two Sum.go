 func twoSum(nums []int, target int) []int {
    hashmap := make(map[int]int)
    for i:=0; i<len(nums); i++ {
        wanted := target - nums[i]
        if v, found := hashmap[wanted]; found {
            return []int{i,v}
        } else {
            hashmap[nums[i]] = i
        }
    }
    return []int{0,0} // to make compiler happy but doesnt do anything because of the constraint
}

// really wanted to redo the easier ones in go to learn it more

// still 0ms beats 100% but ofc faster then python
