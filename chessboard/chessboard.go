package chessboard
// import "fmt"

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
    count := 0
    for square, r := range cb {
        // fmt.Printf("square: " + square)
        // fmt.Printf("r: %b", r)
        if square == rank{
            for _, value := range r {
				if value{
            		count++
                }
            }
        }
    }
	return count
	panic("Please implement CountInRank()")
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
    count := 0
    if file > 0 && file < 9{
        for _, r := range cb {
        	for i, value := range r {
				if i+1 == file && value{
            		count++
                }
            }
        }
    }
    return count
	panic("Please implement CountInFile()")
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
    count := 0
    for _, r := range cb {
    	for _, value := range r {
            if value || !value{
            	count++
            }
        }
    }
	return count
	panic("Please implement CountAll()")
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
    count := 0
    for _, r := range cb {
    	for _, value := range r {
            if value{
            	count++
            }
        }
    }
	return count
	panic("Please implement CountOccupied()")
}
