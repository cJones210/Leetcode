func convert(s string, numRows int) string {
    if numRows == 1 {return s}
    matrix := [][]string{}
    for i:=0; i < numRows; i++ {
        row := []string{}
        matrix = append(matrix, row)
    }

    var goingDown bool
    rowPointer := 0
    for i:=0; i < len(s); i++ {
        matrix[rowPointer] = append(matrix[rowPointer], string(s[i]))
        if rowPointer == numRows - 1 {
            goingDown = false
        } else if rowPointer == 0 {
            goingDown = true
        }
        if goingDown {
            rowPointer++
        } else {
            rowPointer--
        }
    }
    var result string
    for _, row := range matrix {
        for _, val := range row {
            result += val
        }
    }
    return result
}
//1 # # 7 # # 3
//2 # 6 8 # 2 4
//3 5 # 9 1 # 5
//4 # # 0 # # 6
//its always a full numRows amount of chars used and then the
//diagonal is always numRows-2 long

//this problem took me a bit of time because i needed to learn go slices, but it is pretty unique 
//i solved this a few times but the code was a slow for loop mess
//if i had to guess the fastest way to do this would be to not actually make an array, but just make a function based off the table and ideas i had above but that is a little hard

// while my solution is slowish it is still o(n) so i am happy with it
// 10ms beats 37% 
