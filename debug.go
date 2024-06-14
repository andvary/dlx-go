package mydlx

import (
	"fmt"
	"sort"
	"strings"
)

// log выводит в консоль отладочную информацию. Не используем специализированные пакеты для логирования,
// т.к. нам нужно печатать многострочный вывод, с которым они не справляются.
func (d *DLX) log(f func() string) {
	if d.debug {
		fmt.Print(f())
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

// dump возвращает строковое представление текущего состояния матрицы (без удалённых итемов и опций)
// Элементы будут отсортированы по возрастанию.
func (d *DLX) dump() string {
	b := strings.Builder{}
	var items []string
	for i := d.items[0].next; i != 0; i = d.items[i].next {
		items = append(items, d.items[i].name)
	}
	sort.Strings(items)
	b.Write([]byte(strings.Join(items, " ")))

	for i := d.opts[0].next; i != 0; i = d.opts[i].next {
		b.Write([]byte(d.dumpOptions(i)))
	}

	return b.String()
}
