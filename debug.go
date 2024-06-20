package dlx

import (
	"fmt"
	"sort"
	"strings"
)

// log выводит в консоль отладочную информацию. Не используем специализированные пакеты для логирования,
// т.к. нужно печатать многострочный вывод, с которым они не справляются.
func (d *DLX) log(msg string) {
	if d.debug {
		fmt.Println(msg)
		fmt.Println("potential solution: ", d.potentialSolution)
		fmt.Println(d.dump())
	}
}

func (d *DLX) dumpOptions(opts ...int) string {
	b := strings.Builder{}
	var items []string
	for _, i := range opts {
		for j := range d.opts[i].items {
			items = append(items, d.items[j].name)
		}
		sort.Strings(items)
		b.WriteString(strings.Join(items, " "))
		b.WriteByte('\n')
		items = items[:0]
	}

	return b.String()[0 : b.Len()-1]
}

func (d *DLX) dumpItems() string {
	b := strings.Builder{}
	b.WriteString("items: ")
	for i := d.items[0].next; i != 0; i = d.items[i].next {
		b.WriteString(fmt.Sprintf("{%s: p: %d, n: %d, cnt: %d} ", d.items[i].name, d.items[i].prev,
			d.items[i].next, d.items[i].cnt))
	}
	return b.String()
}

// dump возвращает строковое представление текущего состояния матрицы (без удалённых итемов и опций)
// Элементы будут отсортированы по возрастанию.
func (d *DLX) dump() string {
	b := strings.Builder{}

	b.WriteString(d.dumpItems())
	b.WriteByte('\n')

	var items []string
	for i := d.items[0].next; i != 0; i = d.items[i].next {
		items = append(items, d.items[i].name)
	}
	sort.Strings(items)
	b.Write([]byte(strings.Join(items, " ")))
	b.WriteByte('\n')

	for i := d.opts[0].next; i != 0; i = d.opts[i].next {
		b.Write([]byte(d.dumpOptions(i)))
		b.WriteByte('\n')
	}

	return b.String()
}
