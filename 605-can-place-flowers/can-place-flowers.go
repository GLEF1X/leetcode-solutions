func canPlaceFlowers(flowerbed []int, n int) bool {
    prev, next := -1, -1
    var howManyPlanted int
    for i := 0; i < len(flowerbed); i++ {
        if howManyPlanted == n {
            return true
        }

        if flowerbed[i] == 1 {
            continue
        }

        if i + 1 < len(flowerbed) {
            next = flowerbed[i + 1]
        } else {
            next = -1
        }
        if i - 1 >= 0 {
            prev = flowerbed[i - 1]
        } else {
            prev = -1
        }

        // [0, 0, 0] -> allowed at 0 and 2
        // [0, 0, 1] -> allowed at 0
        // [0, 1, 0] -> not allowed at all
        // [1, 0, 0] -> allowed at 2
        // [1, 1, 0] -> not allowed at all
        // [1, 0, 1] -> not allowed at all
        // [1, 1, 1] -> not allowed at all
        // [0, 1, 1] -> not allowed at all

        if (next == -1 || next == 0) && (prev == -1 || prev == 0) {
            flowerbed[i] = 1
            howManyPlanted++
        }
    } 

    return howManyPlanted == n
}