# Go Array Index Out of Bounds

This repository demonstrates a common error in Go: accessing an array index that is out of bounds.  The `bug.go` file contains the erroneous code, while `bugSolution.go` provides a corrected version.

**Problem:** Go arrays have a fixed size. Attempting to access an element using an index greater than or equal to the array's length will trigger a runtime panic.

**Solution:** Ensure that all array accesses are within the valid index range (0 to len(array)-1).  Consider using slices if the size of your data is not known in advance or if you require dynamic resizing.