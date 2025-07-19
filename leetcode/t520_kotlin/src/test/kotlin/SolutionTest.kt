import org.example.Solution
import org.junit.jupiter.api.Assertions
import org.junit.jupiter.api.Test

import org.junit.jupiter.api.Assertions.*

class SolutionTest {

    @Test
    fun test1() {
        val s = Solution()
        //k = 2, w = 0, profits = [1,2,3], capital = [0,1,1]
        val result = s.findMaximizedCapital(2, 0, intArrayOf(1,2,3), intArrayOf(0,1,1))
        Assertions.assertEquals(4, result)
    }

    @Test
    fun test2() {
        val s = Solution()
        //k = 3, w = 0, profits = [1,2,3], capital = [0,1,2]
        val result = s.findMaximizedCapital(3, 0, intArrayOf(1,2,3), intArrayOf(0,1,2))
        Assertions.assertEquals(6, result)
    }

    // k = 1 w = 0 profits = [1,2,3] capital = [1,1,2]
    @Test
    fun test3() {
        val s = Solution()
        //k = 3, w = 0, profits = [1,2,3], capital = [0,1,2]
        val result = s.findMaximizedCapital(1, 0, intArrayOf(1,2,3), intArrayOf(1,1,2))
        Assertions.assertEquals(6, result)
    }
}