class DNA (chain: String) {

  type Count = Int
  type Nucleotide = Char
  type Nucleotides = Map[Nucleotide, Count]

  private val valid = List('A','T','C','G')
  private val defaults = valid.map(_ -> 0)
  private val empty = Map[Nucleotide, Count](defaults: _*)

  private val nucleotides = chain.foldLeft(empty) { (map, n) =>
    require(valid.exists(n == _))
    map + (n -> (map(n) + 1))
  }

  def count(n: Nucleotide): Count = {
    require(valid.exists(n == _))
    nucleotides(n)
  }
  def nucleotideCounts: Nucleotides = nucleotides
}