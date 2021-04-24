## Практические задания

### 1. 99 bottles

Дано: [простой скрипт](/99bottles/main.go) с [тестами](/99bottles/main_test.go).

Требуется найти как можно больше стилистических проблем в коде, которые можно исправить с помощью ruleguard правил.
Вам нужно написать эти правила.

Требования: после изменений, предлагаемых диагностиками, тесты должны по-прежнему проходить.

### 2. Разрешение конфликтов для правил

Нужно реализовать диагностику, которая будет предлагать предлагать заменять `x = x + 2` на `x += 2`, но для случая
`x = x + 1` рекомендация будет выглядеть как `x++`.

Успешно выполненное задание означает два предупреждения на этом коде:

```go
package target

func example(x int) {
  x = x + 4 // want `\Qcould rewite as x += 4`
  x = x + 1 // want `\Qcould rewrite as x++`
}
```

### 3. unslice чекер из go-critic

Портируем [unslice](https://go-critic.github.io/overview.html#unslice) диагостику из анализатора [gocritic](https://github.com/go-critic/go-critic).

![image](https://user-images.githubusercontent.com/6286655/115932102-ef30cd00-a494-11eb-8ed6-96e70cf9ff85.png)

Стоит учесть все граничные случаи (их не так много, но они есть).

Вам потребуются [паттерны типов](https://go-ruleguard.github.io/by-example/type-patterns).
