# Рулетка

Программа оформленна в виде unix way. Для сохранения исходной задачи потребуется порядка 350G.
Количество комбинаций считается по формуле сочетаний (комбинаторика).

```bash
$ go run roulette.go 36 18 | head
Combinations: 9075135300
000000000000000000111111111111111111
000000000000000001011111111111111111
000000000000000001101111111111111111
000000000000000001110111111111111111
```

И первая версия, сделанная для проверки идеи:

```bash
$ go run calc.go 
000000000000000000111111111111111111
111111111111111111000000000000000000
9075135300
```
