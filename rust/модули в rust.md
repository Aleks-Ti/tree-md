# модули в Rust

## примеры

Примитивный пример:

```rust
/// Модуль с функциями-помощниками
mod helpers {

    /// Публичная функция для "приветствия"
    pub fn greet() {
        println!("Hello, world!");
    }
}

fn main() {
    // вызов функции из модуля
    helpers::greet();
}
```

Лучше всё это выносить в отдельные модули, например `helpers.rs`, положить туда код

```rust
/// Публичная функция для "приветствия"
pub fn greet() {
    println!("Hello, world!");
}
```

ТОгда в main мы вызваем функцию помощник так:

```rust
mod helpers;

fn main() {
    helpers::greet();
}
```

Или сделать более структурировано:

- Создать папку `src/utils`, а в ней уже `src/utils/helpers.rs` и `src/utils/mod.rs`

Получится такая структура:

```bash
project_name/
├── src
│   ├── utils
│   │   ├── helpers.rs
│   │   └── mod.rs
│   └── main.rs
├── Cargo.lock
└── Cargo.toml
```

в mod.rs объявим:

```rust
pub mod helpers;
```

а helpers останется как и был:

```rust
/// Публичная функция для "приветствия"
pub fn greet() {
    println!("Hello, world!");
}
```

Тогда в main мы используем код вот так:

```rust
mod helpers;

fn main() {
    helpers::greet();
}
```

## mod.rs

Как работает `mod.rs`

Компилятор, встречая декларацию mod utils, ищет нужный модуль, в том числе в utils/mod.rs. Файл `mod.rs`, в свою очередь, должен:
Явно объявить все подмодули, которые находятся в этом уровне иерархии, используя ключевое слово `mod`.
Доступные извне модули дополнить модификатором pub.
