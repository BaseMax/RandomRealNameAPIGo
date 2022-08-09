import java.io.File

val femaleName = readFromResource("./src/main/resources/female-first-names.txt")
val maleName = readFromResource("./src/main/resources/male-first-names.txt")
val familyName = readFromResource("./src/main/resources/last-names.txt")

val middleChars = listOf(".", "_")

fun main() {
    println(generateRandomNames(100, Gender.MALE))
}

fun generateRandomNames(limit: Int, gender: Gender): List<String> {
    val nameList = when (gender) {
        Gender.MALE -> maleName
        Gender.FEMALE -> femaleName
        Gender.BOTH -> maleName + femaleName
    }

    var result: ArrayList<String> = arrayListOf()

    for (i in 0..limit)
        result.add(generateName(nameList))

    return result
}

fun generateName(names: List<String>): String {
    var name: String = ""
    name += names.random()
    name += addInMiddle()
    name += familyName.random()
    name += addNumber()
    return name
}

fun addNumber(): String {
    if (!getRandomBit())
        return ""

    return (0..9999).random().toString()
}

fun addInMiddle(): String {
    if (!getRandomBit())
        return ""

    return middleChars.random()
}

fun getRandomBit(): Boolean = Math.random() < 0.5

fun readFromResource(path: String): List<String> {
    val file = File(path).inputStream()
    val lineList = mutableListOf<String>()

    file.bufferedReader().forEachLine { lineList.add(it) }
    return lineList
}

