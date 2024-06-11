package mydlx

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func (d *DLX) readInput(p string) error {
	f, err := os.OpenFile(p, os.O_RDONLY, 0111)
	if err != nil {
		return err
	}
	defer f.Close()

	var firstLine = true
	d.opts = make([]*opt, 0)
	d.opts = append(d.opts, &opt{})

	s := bufio.NewScanner(f)
	for s.Scan() {
		bb := bytes.Fields(bytes.TrimSpace(s.Bytes()))
		// пропускаем пустые строки
		if len(bb) == 0 {
			continue
		}
		// заполняем список итемов
		if firstLine {
			if err := d.addItems(bb); err != nil {
				return err
			}
			firstLine = false
			continue
		}

		o := &opt{items: make(map[int]struct{})}

		// добавляем итемы в опцию
		for i, itemName := range bb {
			if len(bb[i]) > 8 {
				return InputError{
					msg: fmt.Sprintf("bad item name %s (must be no more than 8 characters long)", bb[i]),
				}
			}

			n := d.getItem(string(itemName))
			if n < 0 {
				return InputError{
					msg: fmt.Sprintf("option refers to a no-existent item %q", itemName),
				}
			}
			o.items[n] = struct{}{}
			// увеличиваем количество опций, покрывающих этот итем
			d.items[n].cnt++
		}

		if len(o.items) == 0 {
			return InputError{
				msg: fmt.Sprintf("empty option"),
			}
		}
		// добавляем опцию и прилинковываем её к предыдущей
		d.opts = append(d.opts, o)
		o.prev = len(d.opts) - 2
		d.opts[len(d.opts)-2].next = len(d.opts) - 1
	}

	// линкуем первую и последнюю опции
	d.opts[0].prev, d.opts[len(d.opts)-1].next = len(d.opts)-1, 0

	// убеждаемся, что у всех итемов есть хотя бы одна опция, иначе решение не существует
	for i := 1; i < len(d.items); i++ {
		if d.items[i].cnt == 0 {
			return InputError{
				msg: fmt.Sprintf("item has 0 options: %v", d.items[i]),
			}
		}
	}

	return nil
}

func (d *DLX) addItems(bb [][]byte) error {
	d.items = make([]*item, len(bb)+1)
	// добавляем корневой элемент
	d.items[0] = &item{
		name: "",
		prev: len(d.items) - 1,
		next: 1,
	}

	for i := range bb {
		if len(bb[i]) > 8 {
			return InputError{
				msg: fmt.Sprintf("bad item name %s (must be no more than 8 characters long)", bb[i]),
			}
		}

		d.items[i+1] = &item{
			name: string(bb[i]),
			prev: i,
			next: i + 2,
		}
	}
	// линкуем последний итем с первым
	d.items[len(d.items)-1].next = 0
	return nil
}

func (d *DLX) getItem(name string) int {
	for i := range d.items {
		if d.items[i].name == name {
			return i
		}
	}
	return -1
}
