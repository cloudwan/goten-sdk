package resource

// PagerQuery wraps paging parameters of a query.
type PagerQuery interface {
	GetCursor() Cursor
	GetOrderBy() OrderBy
	GetLimit() int
	GetPeekForward() bool
	PageDirection() PageDirection
	GetResourceDescriptor() Descriptor
}

// HasForwardCursor returns information view is long enough to justify presence of forward cursor
func HasForwardCursor(pager PagerQuery, view ResourceList) bool {
	if pager == nil || pager.GetLimit() == 0 {
		return false
	}
	minViewLen := pager.GetLimit()
	if pager.GetPeekForward() {
		minViewLen += 1
	}
	return view.Length() >= minViewLen
}

// ApplyPagerToViewUsingOffset creates next and prev page cursors by checking offset. Forward cursor will be empty, if
// insufficient results are present. Backward cursor will be reversed version of previous cursor (may be empty).
// Cursors are computed based on offset from the beginning of array. In other words, CursorValueType will be
// of OffsetCursorValueType. View will be adjusted to the pager too.
func ApplyPagerToViewUsingOffset(pager PagerQuery, view ResourceList) (pagedView ResourceList, nextPageCursor, prevPageCursor Cursor) {
	var previousCursorValue *OffsetCursorValue
	if pager.GetCursor() != nil && !pager.GetCursor().IsEmpty() {
		previousCursorValue = pager.GetCursor().GetValue().(*OffsetCursorValue)
	}

	if (pager.PageDirection().IsRight() && HasForwardCursor(pager, view)) ||
		(pager.PageDirection().IsLeft() && previousCursorValue.GetOffset() > 0) {
		nextPageCursorValue := previousCursorValue.MakeForwardValue(
			pager.PageDirection(), int32(pager.GetLimit()))
		nextPageCursor = pager.GetResourceDescriptor().NewResourceCursor()
		nextPageCursor.SetCursorValue(nextPageCursorValue)
		nextPageCursor.SetInclusion(CursorExclusive)
		nextPageCursor.SetPageDirection(pager.PageDirection())
	}

	if previousCursorValue != nil {
		prevPageCursorValue := previousCursorValue.MakeBackwardValue(
			pager.PageDirection(), int32(pager.GetLimit()))
		prevPageCursor = pager.GetResourceDescriptor().NewResourceCursor()
		prevPageCursor.SetCursorValue(prevPageCursorValue)
		prevPageCursor.SetInclusion(pager.GetCursor().GetInclusion().Reverse())
		prevPageCursor.SetPageDirection(pager.GetCursor().GetPageDirection().Reverse())
	}

	if view.Length() > pager.GetLimit() {
		pagedView = view.Slice(0, pager.GetLimit())
	} else if view.Length() <= pager.GetLimit() {
		pagedView = view
	}
	if pager.PageDirection().IsLeft() {
		nextPageCursor, prevPageCursor = prevPageCursor, nextPageCursor
	}
	return
}

// ApplyPagerToViewUsingSnapshot creates next and prev page cursors for slice of results. Forward cursor will be empty,
// if insufficient results are present. Backward cursor will be reversed version of previous cursor (empty if not set).
// Cursors are built based on resource snapshot taken from the view - taken from the end if pager points
// right, otherwise from beginning. In other words, CursorValueType will be of SnapshotCursorValueType.
// View will be adjusted to the pager too.
func ApplyPagerToViewUsingSnapshot(pager PagerQuery, view ResourceList) (pagedView ResourceList, nextPageCursor, prevPageCursor Cursor) {
	if HasForwardCursor(pager, view) {
		var idx int
		if pager.PageDirection().IsRight() {
			idx = pager.GetLimit() - 1
		} else {
			idx = view.Length() - pager.GetLimit()
		}
		nextPageCursorValue := NewSnapshotCursorValue(
			pager.GetOrderBy().GetRawFieldMask().ProjectRaw(view.At(idx)).(Resource))
		nextPageCursor = pager.GetResourceDescriptor().NewResourceCursor()
		nextPageCursor.SetCursorValue(nextPageCursorValue)
		nextPageCursor.SetInclusion(CursorExclusive)
		nextPageCursor.SetPageDirection(pager.PageDirection())
	}

	if pager.GetCursor() != nil && !pager.GetCursor().IsEmpty() {
		prevPageCursorValue := pager.GetCursor().GetValue().(*SnapshotCursorValue)
		prevPageCursor = pager.GetResourceDescriptor().NewResourceCursor()
		prevPageCursor.SetCursorValue(prevPageCursorValue)
		prevPageCursor.SetInclusion(pager.GetCursor().GetInclusion().Reverse())
		prevPageCursor.SetPageDirection(pager.GetCursor().GetPageDirection().Reverse())
	}

	if pager.PageDirection().IsRight() && pager.GetLimit() < view.Length() {
		pagedView = view.Slice(0, pager.GetLimit())
	} else if view.Length() <= pager.GetLimit() {
		pagedView = view
	} else {
		pagedView = view.Slice(1, 0)
	}
	if pager.PageDirection().IsLeft() {
		nextPageCursor, prevPageCursor = prevPageCursor, nextPageCursor
	}
	return
}

// ChangesWatchPager is a helper stateful object containing cached view to inform of page token changes.
// It should be used in conjunction with stateful watch methods only if cursor values are based on resource snapshot.
// In offset-based queries, cursors never change so there is no equivalent structure for it.
type ChangesWatchPager struct {
	pager PagerQuery
	view  ResourceList

	// flag relevant only for handling backward cursor, which doesn't change
	// we want to return it on first run regardless of dynamic forward cursor
	pastInitialProcessing bool
}

func NewChangesWatchPager(pager PagerQuery, descriptor Descriptor) *ChangesWatchPager {
	return &ChangesWatchPager{pager: pager, view: descriptor.NewResourceList(0, 0)}
}

// ProcessChanges adjust view of page and returns new cursors (next and prev page) when they need to be updated
func (wp *ChangesWatchPager) ProcessChanges(changes ResourceChangeList) (Cursor, Cursor, bool) {
	var changed, hadForwardTokenBeforeChanges bool
	var fwdIdx int

	// handle forward cursor - dynamically updated
	if wp.pager == nil || wp.view.Length() < wp.pager.GetLimit() || wp.pager.GetLimit() == 0 {
		fwdIdx = -1
		hadForwardTokenBeforeChanges = false
	} else {
		if wp.pager.PageDirection().IsRight() {
			fwdIdx = wp.pager.GetLimit() - 1
		} else {
			fwdIdx = wp.view.Length() - wp.pager.GetLimit()
		}
	}

	for i := 0; i < changes.Length(); i++ {
		change := changes.At(i)

		if change.IsAdd() {
			// insert at index
			viewIdx := int(change.GetCurrentViewIndex())
			wp.view = wp.view.Slice(0, viewIdx).Append(change.GetResource()).AppendList(wp.view.Slice(viewIdx, 0))
		} else if change.IsDelete() {
			// remove at idx
			viewIdx := int(change.GetPreviousViewIndex())
			wp.view = wp.view.Slice(0, viewIdx).AppendList(wp.view.Slice(viewIdx+1, 0))
		} else if change.IsModify() {
			// modify - if index diffs, remove and add
			prevViewIdx := int(change.GetPreviousViewIndex())
			newViewIdx := int(change.GetCurrentViewIndex())
			if prevViewIdx != newViewIdx {
				wp.view = wp.view.Slice(0, prevViewIdx).AppendList(wp.view.Slice(prevViewIdx+1, 0))
				wp.view = wp.view.Slice(0, newViewIdx).Append(change.GetResource()).AppendList(wp.view.Slice(newViewIdx, 0))
			}
			// modify may affect present forward cursor with property update or position shift
			if hadForwardTokenBeforeChanges && prevViewIdx == fwdIdx || newViewIdx == fwdIdx {
				changed = true
			}
		}
	}
	if hadForwardTokenBeforeChanges != HasForwardCursor(wp.pager, wp.view) {
		changed = true
	}
	if !changed && wp.pastInitialProcessing {
		return nil, nil, changed
	}
	_, nextPageCursor, prevPageCursor := ApplyPagerToViewUsingSnapshot(wp.pager, wp.view)
	wp.pastInitialProcessing = true
	return nextPageCursor, prevPageCursor, changed
}
