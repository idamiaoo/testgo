package mike

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if left > right {
		return left + 1
	}
	return right + 1
}

func MaxDistance(root *Node) int {
	if root == nil {
		return 0
	}

	leftMaxDistance := MaxDistance(root.Left)
	leftMaxDepth := maxDepth(root.Left)
	rightMaxDistance := MaxDistance(root.Right)
	rightMaxDepth := maxDepth(root.Right)
	distance := leftMaxDepth + rightMaxDepth
	if distance > leftMaxDistance && distance > rightMaxDistance {
		return distance
	} else {
		if leftMaxDistance > rightMaxDistance {
			return leftMaxDistance
		}
		return rightMaxDistance
	}
}
