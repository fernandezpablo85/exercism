import scala.collection.immutable.TreeMap

class School {

  type Grade = Int
  type Name = String
  type Db = TreeMap[Grade, List[Name]]

  private val emptyNames = List.empty[Name]
  private val empty = TreeMap[Grade, List[Name]]()

  var db: Db = empty

  def add(name: Name, grade: Grade): Unit = {
    val updatedNames = db.getOrElse(grade, emptyNames) :+ name
    db = db + (grade -> updatedNames)
  }

  def grade(grade: Grade): List[Name] = db.getOrElse(grade, emptyNames)

  def sorted: Db = db.foldLeft(empty)(sortNames)

  def sortNames(map: Db, entry: (Grade, List[Name])): Db = {
    val (k, names) = entry
    map + (k -> names.sorted)
  }
}