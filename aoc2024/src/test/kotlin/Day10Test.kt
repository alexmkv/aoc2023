import org.junit.jupiter.api.Assertions.*
import org.junit.jupiter.api.Test

class Day10Test {
    @Test
    fun calc01_0() {
        assertEquals(36, Day10("data/data10_1_0.txt").calc())
    }
    @Test
    fun calc01_0_1() {
        assertEquals(1, Day10("data/data10_1_0_1.txt").calc())
    }
    @Test
    fun calc01_1() {
        assertEquals(737, Day10("data/data10_1_1.txt").calc())
    }

    @Test
    fun calc02_0() {
        assertEquals(3, Day10("data/data10_2_0.txt", 2).calc())
    }

    @Test
    fun calc02_1() {
        assertEquals(3, Day10("data/data10_2_1.txt", 2).calc())
    }

    @Test
    fun calc02_3() {
        assertEquals(6373055193464, Day10("data/data10_1_1.txt", 2).calc())
    }
}