package react

/*
        Implements a basic Reactive system with multiple types of cells:
            
            Input       cells with settable values

            Compute     cells whose values are computed in terms of other cells,
        allows for registering callbacks which are fired when the cell's value
        changes. These callbacks are known as change notification callbacks.
        
        When an input value is changed, values propogate to reach a new stable
        system state.
*/

type cb *func()

type cell struct {
        value
}

func (c *cell) Value() {
        return c.value
}

type grid struct {
        cells []*cells
}

func (s *grid) update {
        
}

func New() Reactor {
        return &grid
}
