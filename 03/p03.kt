import java.io.File

data class Claim(val id: Int, val start: Pair<Int, Int>, val extent: Pair<Int, Int> )

fun claimIt(field: Array<IntArray>, claim: Claim) {
    val start_x = claim.start.first
    val start_y = claim.start.second
    val extent_x = claim.extent.first
    val extent_y = claim.extent.second
    for (row in start_y until start_y + extent_y){
        for (col in start_x until start_x + extent_x){
            field[row][col] = field[row][col] + 1
        }
    }
}

fun isIntact(field: Array<IntArray>, claim: Claim) : Boolean {
    val start_x = claim.start.first
    val start_y = claim.start.second
    val extent_x = claim.extent.first
    val extent_y = claim.extent.second
    for (row in start_y until start_y + extent_y){
        for (col in start_x until start_x + extent_x){
            if (field[row][col] > 1){
                return false
            }
        }
    }
    return true
}

fun parseClaim(line: String): Claim {
    val fields = line.split(" ")
    val start = fields[2].split(",")
    val extent = fields[3].split("x")
    val start_pair = Pair(start[0].toInt(), start[1].removeSuffix(":").toInt())
    val extend_pair = Pair(extent[0].toInt(), extent[1].toInt())
    return Claim(fields[0].removePrefix("#").toInt(), start_pair, extend_pair)
}

val p03 =  fun(input: List<String>) {
    var claims = input.map { parseClaim(it) }

    var field = Array(2000, {IntArray(2000)})
    for (claim in claims) {
        claimIt(field, claim)
    }
    var count = 0
    for (line in field) {
        count += line.count { it -> it > 1 }
    }
    println(count)
    for (claim in claims) {
        if (isIntact(field, claim)) {
            println(claim.id)
        }
    }
}