import org.example.subarraysDivByK
import org.junit.jupiter.api.Test

import org.junit.jupiter.api.Assertions.*

class MainKtTest {

    @Test
    fun subarraysDivByK_emptyArray() {
        assertEquals(0, subarraysDivByK(IntArray(0), 1))
    }

    @Test
    fun subarraysDivByK_oneElement1() {
        assertEquals(1, subarraysDivByK(intArrayOf(3), 3))
    }

    @Test
    fun subarraysDivByK_oneElement0() {
        assertEquals(0, subarraysDivByK(intArrayOf(3), 4))
    }

    @Test
    fun subarraysDivByK_test1() {
        assertEquals(7, subarraysDivByK(intArrayOf(4,5,0,-2,-3,1), 5))
    }

    @Test
    fun subarraysDivByK_test2() {
        assertEquals(2, subarraysDivByK(intArrayOf(5,3,5), 5))
    }
}