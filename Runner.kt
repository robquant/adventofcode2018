import kotlin.system.measureNanoTime
import java.io.File

data class Puzzle(val day: Int, val part: Int? = null)

object Puzzles {

    private val puzzles = mapOf(
        Puzzle(1) to p01,
        Puzzle(2) to p02,
        Puzzle(3) to p03,
        Puzzle(5) to p05
    )

    fun run(day: Int, part: Int? = null) {
        val input = File("%02d".format(day) + "/input.txt").readLines()
        puzzles.filter { it.key == Puzzle(day, part) }.values.single()(input)
    }

}

fun main(args: Array<String>) {

    val (repeat, day, part) = parseArgs(args)
    val times = (1..repeat).map {
        measureNanos {
            Puzzles.run(day, part)
        }
    }

    println("\nTimes: ${times.map { (it / 1e6).toInt() }.joinToString()} ms")
}

inline fun measureNanos(block: () -> Unit) = measureNanoTime(block)

fun parseArgs(args: Array<String>): Triple<Int, Int, Int?> {
    val repeat = args[0].toInt()
    val day = args[1].toInt()
    val part = if (args.size == 3) args[2].toInt() else null
    return Triple(repeat, day, part)
}