import org.junit.jupiter.api.Assertions.*
import org.junit.jupiter.api.Test

class Day09Test {
    @Test
    fun calc01_0() {
        assertEquals(1928, Day09("data/data09_1_0.txt").calc())
    }
    @Test
    fun calc01_1() {
        assertEquals(6346871685398, Day09("data/data09_1_1.txt").calc())
    }

    @Test
    fun calc02_0() {
        assertEquals(2858, Day09("data/data09_1_0.txt", 2).calc())
    }

    @Test
    fun calc02_1() {
        assertEquals(6373055193464, Day09("data/data09_1_1.txt", 2).calc())
    }
}