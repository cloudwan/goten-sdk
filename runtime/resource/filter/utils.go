package filter

func CompareSatisfies(refop, op CompareOperator, cmp int) (res bool) {
	if refop == op && cmp == 0 {
		return true
	}

	switch refop {
	case Neq:
		return (op == Eq && cmp != 0) || (op == Gt && cmp >= 0) || (op == Gte && cmp > 0) || (op == Lt && cmp <= 0) || (op == Lte && cmp < 0)
	case Lte:
		return (op == Eq && cmp <= 0) || (op == Lt && cmp <= 0) || (op == Lte && cmp <= 0)
	case Lt:
		return (op == Eq && cmp < 0) || (op == Lt && cmp <= 0) || (op == Lte && cmp < 0)
	case Gte:
		return (op == Eq && cmp >= 0) || (op == Gt && cmp >= 0) || (op == Gte && cmp >= 0)
	case Gt:
		return (op == Eq && cmp > 0) || (op == Gt && cmp >= 0) || (op == Gte && cmp > 0)
	default:
		return false
	}
}
