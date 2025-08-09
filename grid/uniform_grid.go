package grid

import "math"

type cellKey struct{
	X, Y int
}

// axis-aligned bounding box
type AABB struct {
	MinX, MinY, MaxX, MaxY float64
}

// holds a collider's info for storing in buckets
type record struct {
	EntityID int
	Bounds AABB
	Static bool
}

// divides space into sizes of CellSize, separate buckets for load time vs update
type UniformGrid struct {
	CellSize float64
	staticBuckets map[cellKey][]record
	dynamicBuckets map[cellKey][]record
}

func NewUniformGrid(cellSize float64) *UniformGrid {
	return &UniformGrid{
		CellSize: cellSize,
		staticBuckets: make(map[cellKey][]record),
		dynamicBuckets: make(map[cellKey][]record),
	}
}

// converts a coord into the coords of the corresponding bucket/cell
func (grid *UniformGrid) cellOf(x, y float64) cellKey {
	return cellKey{
		X: int(math.Floor(x/grid.CellSize)),
		Y: int(math.Floor(y/grid.CellSize)),
	}
}

func (grid *UniformGrid) InsertStatic(id int, b AABB) {
	min := grid.cellOf(b.MinX, b.MinY)
	max := grid.cellOf(b.MaxX, b.MaxY)

	// loop goes through every single bucket AABB touches
	for cy := min.Y; cy <= max.Y; cy++ {
		for cx := min.X; cx <= max.X; cx++ {
			key := cellKey{cx, cy}
			grid.staticBuckets[key] = append(
				grid.staticBuckets[key],
				record{EntityID: id, Bounds: b, Static: true},
			)
		}
	}
}

func (grid *UniformGrid) InsertDynamic(id int, b AABB) {
	min := grid.cellOf(b.MinX, b.MinY)
	max := grid.cellOf(b.MaxX, b.MaxY)

	// loop goes through every single bucket AABB touches
	for cy := min.Y; cy <= max.Y; cy++ {
		for cx := min.X; cx <= max.X; cx++ {
			key := cellKey{cx, cy}
			grid.dynamicBuckets[key] = append(
				grid.dynamicBuckets[key],
				record{EntityID: id, Bounds: b, Static: false},
			)
		}
	}
}

func (grid *UniformGrid) ClearDynamic() {
	grid.dynamicBuckets = make(map[cellKey][]record)
}

// returns all records (both static and dynamic) whose AABB overlaps b
func (grid *UniformGrid) Query(b AABB) []record {
	min := grid.cellOf(b.MinX, b.MinY)
	max := grid.cellOf(b.MaxX, b.MaxY)
	seen := make(map[int]bool)
	var out []record

	collect := func(bucket map[cellKey][]record) {
		for cy := min.Y; cy <= max.Y; cy++ {
			for cx := min.X; cx <= max.X; cx++ {
				for _, rec := range bucket[cellKey{cx, cy}] {
					if !seen[rec.EntityID] && overlaps(rec.Bounds, b) {
						seen[rec.EntityID] = true
						out = append(out, rec)
					}
				}
			}
		}
	}

	collect(grid.staticBuckets)
	collect(grid.dynamicBuckets)
	return out
}

// check for overlap between two AABBs
func overlaps(a, b AABB) bool {
	return a.MinX <= b.MaxX && a.MaxX >= b.MinX &&
		a.MinY <= b.MaxY && a.MaxY >= b.MinY
}
