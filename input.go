package dlx

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

// readInput считывает входные данные в следующем виде:
// Первая строка: список всех итемов (items у Кнута), разделённых пробелами
// Вторая и последующая строки: опции(options у Кнута), каждая из которых покрывает не менее одного итема из
// перечисленных в первой строке.
// Итемы представлены структурами item, которые хранятся в массиве. Дополнительно у каждого итема есть указатель
// на предыдущий и следующий итем (вместо обычного указателя используется индекс в массиве).
// Нулевой элемент массива с итемами всегда пустой, он служит для удобства обхода списка итемов.
// Аналогично организованы опции. Здесь я немного отступаю от того, как сделано у кнута и храню спиок итемов, покрываемых
// Каждой опцие в мапе, а не в ещё одном линкованном списке, что сильно упрощает код.
func (d *DLX) readInput(r io.Reader) error {
	var firstLine = true
	d.opts = make([]*opt, 0)
	d.opts = append(d.opts, &opt{})

	s := bufio.NewScanner(r)
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
