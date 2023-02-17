func oddEvenJumps(arr []int) int {   
    canOddJump := make([]bool, len(arr))
    canEvenJump := make([]bool, len(arr))
    
    canOddJump[len(arr) - 1] = true
    canEvenJump[len(arr) - 1] = true
    
    total := 0
    tm := treemap.NewWithIntComparator()
    
    for i := len(arr) - 1; i >= 0; i-- {
        // Ceiling finds the smallest key that is larger than or equal to the given key
        if k, v := tm.Ceiling(arr[i]); k != nil {
            canOddJump[i] = canEvenJump[v.(int)]
        }
        
        // Floor finds the largest key that is smaller than or equal to the given key
        if k, v := tm.Floor(arr[i]); k != nil {
            canEvenJump[i] = canOddJump[v.(int)]
        }
        
        if canOddJump[i] {
            total++
        }
        tm.Put(arr[i], i)
    }
    return total
}