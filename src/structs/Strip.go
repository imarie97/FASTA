package structs

// Strip in a matrix generated by S1 and S2 of some SeqPair.
// Determined by left and right bound diagonals (offsets).
type Strip struct {
    Left, Right Diagonal
}
