package design_hit_counter_362

import java.util.ArrayDeque

class Solution {
    private val hits: ArrayDeque<Int> = ArrayDeque()

    fun hit(timestamp: Int) {
        hits.addLast(timestamp)
    }

    fun getHits(timestamp: Int): Int {
        cleanOldHits(timestamp)
        return hits.size
    }

    private fun cleanOldHits(currentTime: Int) {
        val threshold = currentTime - 5 * 60
        while (hits.isNotEmpty() && hits.first() <= threshold) {
            hits.removeFirst()
        }
    }
}

fun test() {
    val obj = Solution()
    obj.hit(1)
    obj.hit(2)
    obj.hit(3)
    println(obj.getHits(4))
    obj.hit(300)
    println(obj.getHits(300))
    println(obj.getHits(301))
}
