package dlx

import (
	"fmt"
)

// cover пытается найти решение следующим образом:
// Берём итем с наименьшим количеством опции (это позволяет сильно сократить время поиска решения)
// Для каждой опции итема:
// 1. Удалить каждый итем, который покрывает данная опция
// 2. Удалить все опции, содержащие удалённые итемы
// 3. Рекурсивно запустить алгоритм на уменьшенной матрице
// 4. Откатить изменения и попробовать следующую опцию
// Если список итемов пуст, решение найдено.
// Если есть итем(ы) без опций - это тупик, нужно откатываться.
// Алгоритм, по сути, перебирает варианты, но работает быстро за счёт того, что удаление и восстановление опций и
// итемов происходит// быстро (нужно только переписать указатели соседней).
func (d *DLX) cover(item int) error {
	var removedOpts []int
	var removedItems []int

	// найдём неудалённую опцию, покрывающую заданный итем
OPTIONS:
	for i := d.opts[0].next; i != 0; i = d.opts[i].next {
		if d.maxSolutions != 0 && len(d.solutions) >= d.maxSolutions {
			return nil
		}

		if !d.opts[i].lItems.isPresent(item) {
			continue
		}

		// запишем опцию, как потенциальную часть решения
		d.potentialSolution = append(d.potentialSolution, i)
		d.log("starting with option: " + d.dumpOptions(i))

		// удалим все итемы, покрытые опцией
		for _, j := range d.opts[i].items {
			if err := d.removeItem(j); err != nil {
				return fmt.Errorf("cover: %v", err)
			}
			removedItems = append(removedItems, j)

			// удалим все опции с удалённым итемом
			for o := d.opts[0].next; o != 0; o = d.opts[o].next {
				if !d.opts[o].lItems.isPresent(j) {
					continue
				}

				if err := d.removeOption(o); err != nil {
					return fmt.Errorf("cover: %v", err)
				}
				removedOpts = append(removedOpts, o)
			}
		}

		// если список итемов пуст или остались только вторичные итемы, решение найдено, проверим другие опции
		if d.items[0].next == 0 || d.items[0].next > d.primaryBoundary {

			d.log("solution found")
			d.addSolution()

			d.restoreOptions(removedOpts)
			d.restoreItems(removedItems)
			removedItems = removedItems[:0]
			removedOpts = removedOpts[:0]

			// убираем из потенциального решения только текущую опцию
			d.potentialSolution = d.potentialSolution[:len(d.potentialSolution)-1]

			d.log("uncovered after finding a solution")
			continue OPTIONS
		}

		// если есть неудалённые итемы без опций, значит решения нет
		for it := d.items[0].next; it != 0 && it <= d.primaryBoundary; it = d.items[it].next {
			if d.items[it].cnt == 0 {

				d.log(fmt.Sprintf("empty item %v found, dead end", d.items[it]))

				d.restoreOptions(removedOpts)
				d.restoreItems(removedItems)
				removedItems = removedItems[:0]
				removedOpts = removedOpts[:0]

				// убираем из потенциального решения только текущую опцию
				d.potentialSolution = d.potentialSolution[:len(d.potentialSolution)-1]
				d.log("uncovered after dead end")

				continue OPTIONS
			}
		}

		d.log("recursing")

		if err := d.cover(d.findBestItem()); err != nil {
			return err
		}
		d.potentialSolution = d.potentialSolution[:len(d.potentialSolution)-1]

		d.log("done recursing")
		d.restoreOptions(removedOpts)
		d.restoreItems(removedItems)
		removedItems = removedItems[:0]
		removedOpts = removedOpts[:0]

		d.log("uncovered after recursing")
	}

	return nil
}

func (d *DLX) restoreItems(ri []int) {
	for i := len(ri) - 1; i >= 0; i-- {
		d.items[d.items[ri[i]].next].prev, d.items[d.items[ri[i]].prev].next = ri[i], ri[i]
	}
}

func (d *DLX) restoreOptions(ro []int) {
	for i := len(ro) - 1; i >= 0; i-- {
		d.opts[d.opts[ro[i]].next].prev, d.opts[d.opts[ro[i]].prev].next = ro[i], ro[i]
		for _, j := range d.opts[ro[i]].items {
			d.items[j].cnt++
		}
	}
}

// removeItem удаляет итем из связанного списка опций (но не из массива), только перезаписывая указатели
// соседних итемов, но не меня указатели самого итема для его последующего восстановления.
func (d *DLX) removeItem(i int) error {
	if i < 1 || i > len(d.items)-1 {
		return CoverError{
			msg: fmt.Sprintf("bad item index %d", i),
		}
	}

	d.items[d.items[i].prev].next, d.items[d.items[i].next].prev = d.items[i].next, d.items[i].prev
	return nil
}

// removeOption удаляет опцию из связанного списка опций (но не из массива), только перезаписывая указатели
// соседних опций, но не меня указатели самой опции для её последующего восстановления.
// Также уменьшает cnt всех итемов, покрываемых данной опцией на 1.
func (d *DLX) removeOption(i int) error {
	if i < 1 || i > len(d.opts)-1 {
		return CoverError{
			msg: fmt.Sprintf("bad option index %d", i),
		}
	}

	d.opts[d.opts[i].prev].next, d.opts[d.opts[i].next].prev = d.opts[i].next, d.opts[i].prev

	for _, j := range d.opts[i].items {
		d.items[j].cnt--
	}
	return nil
}

// findBestItem возвращает индекс первого неудалённого ПЕРВИЧНОГО итема с наименьшим количеством опций.
// Паникует, если в d.items нет ни одного итема, кроме корневого или у итема 0 опций. Вторичные итемы
// игнорируются.
func (d *DLX) findBestItem() int {
	best := d.items[0].next
	if best == 0 {
		panic("find best item: trying to find the best item in an empty items list")
	}

	for i := d.items[0].next; i != 0 && i <= d.primaryBoundary; i = d.items[i].next {
		if d.items[i].cnt == 0 {
			panic(fmt.Sprintf("find best item: item %v has 0 options", d.items[i]))
		}
		if d.items[i].cnt < d.items[best].cnt {
			best = i
		}
	}

	return best
}

func (d *DLX) addSolution() {
	s := make([]int, len(d.potentialSolution))
	copy(s, d.potentialSolution)
	d.solutions = append(d.solutions, s)
}
