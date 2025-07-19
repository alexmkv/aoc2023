import org.example.Solution
import org.junit.jupiter.api.Assertions
import org.junit.jupiter.api.Assertions.*
import org.junit.jupiter.api.Test

class SolutionTest {
    @Test
    fun test1() {
        // nums = [1,3], n = 6
        //  Output: 1
        val solution = Solution()
        Assertions.assertEquals(1, solution.minPatches(intArrayOf(1,3), 6))
    }

    @Test
    fun test2() {
        val solution = Solution()
        Assertions.assertEquals(2, solution.minPatches(intArrayOf(1,5,10), 20))
    }

    @Test
    fun test3() {
        val solution = Solution()
        Assertions.assertEquals(0, solution.minPatches(intArrayOf(1,2,2), 5))
    }

    @Test
    fun test4() {
        val solution = Solution()
        Assertions.assertEquals(28, solution.minPatches(intArrayOf(1,2,31,33), 2147483647))
    }

}